// Copyright 2025 Robert Snedegar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package correlation

import (
	"errors"
	"math/big"
	"strconv"
)

// Type represents the type of correlation coefficient to calculate.
type Type int

const (
	// Pearson correlation coefficient (linear correlation) This is the default value.
	Pearson Type = iota
	// Spearman rank correlation coefficient (monotonic correlation)
	Spearman
	// KendallTau rank correlation coefficient (ordinal association)
	KendallTau
	// GoodmanKruskal rank correlation coefficient (gamma statistic)
	GoodmanKruskal

	// TODO(rsned): Add any other types that come up.
)

// String returns the string representation of the CorrelationType.
func (c Type) String() string {
	switch c {
	case Pearson:
		return "Pearson"
	case Spearman:
		return "Spearman"
	case KendallTau:
		return "Kendall's Tau"
	case GoodmanKruskal:
		return "Goodman and Kruskal's Gamma"
	default:
		return "Unknown"
	}
}

// Numeric represents any primitive numeric type that can be used in correlation calculations.
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// BigNumeric represents big number types that can be used in correlation calculations.
type BigNumeric interface {
	*big.Float | *big.Int
}

// MixedNumeric represents numeric types that can be used in correlation calculations.
type MixedNumeric interface {
	Numeric | BigNumeric
}

// Correlate calculates the specified correlation coefficient between two datasets x and y
// of any numeric type. It returns a value between -1 and 1, where:
// - 1 indicates a perfect positive relationship
// - 0 indicates no relationship
// - -1 indicates a perfect negative relationship
//
// Returns an error if the slices have different lengths, are empty, or if the
// correlation type is not supported.
func Correlate[T Numeric](x, y []T, correlationType Type) (float64, error) {
	switch correlationType {
	case Pearson:
		return Pearsons(x, y)
	case Spearman:
		return Spearmans(x, y)
	case KendallTau:
		return KendallsTau(x, y)
	case GoodmanKruskal:
		return GoodmanKruskals(x, y)
	default:
		return 0, errors.New("unsupported correlation type")
	}
}

// CorrelateBig calculates the specified correlation coefficient between two datasets x and y
// of big number types (*big.Float or *big.Int). It returns a float64 value between -1 and 1, where:
// - 1 indicates a perfect positive relationship
// - 0 indicates no relationship
// - -1 indicates a perfect negative relationship
//
// Returns an error if the slices have different lengths, are empty, or if the
// correlation type is not supported, or any conversion errors occur.
func CorrelateBig[T BigNumeric](x, y []T, correlationType Type) (float64, error) {
	if len(x) != len(y) {
		return 0, errors.New("slices must have the same length")
	}
	if len(x) == 0 {
		return 0, errors.New("slices cannot be empty")
	}

	switch correlationType {
	case Pearson:
		return PearsonsBig(x, y)
	case Spearman:
		return SpearmansBig(x, y)
	case KendallTau:
		return KendallsTauBig(x, y)
	case GoodmanKruskal:
		return GoodmanKruskalsBig(x, y)
	default:
		return 0, errors.New("unsupported correlation type")
	}
}

// CorrelateMixed calculates the specified correlation coefficient
// between two datasets x and y with a set of mixed type inputs.
//
// TODO(rsned): Make this smarter by checking if all the mixed inputs
// are of primitive types and staying in the float64 realm where possible.
func CorrelateMixed[T1, T2 MixedNumeric](x []T1, y []T2, correlationType Type) (float64, error) {
	if len(x) != len(y) {
		return 0, errors.New("slices must have the same length")
	}
	if len(x) == 0 {
		return 0, errors.New("slices cannot be empty")
	}

	// Convert mixed types to big.Float using the helper function
	xVals, err := mixedToBig(x)
	if err != nil {
		return 0, err
	}

	yVals, err := mixedToBig(y)
	if err != nil {
		return 0, err
	}

	return CorrelateBig(xVals, yVals, correlationType)
}

// mixedToBig takes a slice of generic type MixedNumeric and converts each
// element to *big.Float for consistent arithmetic. The method makes no attempt
// at optimizing to float64 even if all values fall within the valid range
// of float64.
//
// This helper function handles the conversion from various numeric types
// (int, float, *big.Int, *big.Float) to *big.Float.
//
// Returns an error if any element cannot be converted.
func mixedToBig[T MixedNumeric](data []T) ([]*big.Float, error) {
	n := len(data)
	result := make([]*big.Float, n)

	for i := range n {
		val := data[i]
		switch v := any(val).(type) {
		case *big.Float:
			result[i] = new(big.Float).Copy(v)
		case *big.Int:
			result[i] = new(big.Float).SetInt(v)
		case int:
			result[i] = new(big.Float).SetInt64(int64(v))
		case int8:
			result[i] = new(big.Float).SetInt64(int64(v))
		case int16:
			result[i] = new(big.Float).SetInt64(int64(v))
		case int32:
			result[i] = new(big.Float).SetInt64(int64(v))
		case int64:
			result[i] = new(big.Float).SetInt64(v)
		case uint:
			result[i] = new(big.Float).SetUint64(uint64(v))
		case uint8:
			result[i] = new(big.Float).SetUint64(uint64(v))
		case uint16:
			result[i] = new(big.Float).SetUint64(uint64(v))
		case uint32:
			result[i] = new(big.Float).SetUint64(uint64(v))
		case uint64:
			result[i] = new(big.Float).SetUint64(v)
		case float32:
			result[i] = new(big.Float).SetFloat64(float64(v))
		case float64:
			result[i] = new(big.Float).SetFloat64(v)
		default:
			return nil, errors.New("unsupported type at index " + strconv.Itoa(i))
		}
	}

	return result, nil
}

// bigNumericToBigFloat converts a single BigNumeric value to *big.Float.
// This is a helper function for converting *big.Float or *big.Int to *big.Float
// for consistent arithmetic operations.
//
// Panics if the input type is not supported (should only be used with BigNumeric types).
func bigNumericToBigFloat[T BigNumeric](val T) *big.Float {
	switch v := any(val).(type) {
	case *big.Float:
		return new(big.Float).Copy(v)
	case *big.Int:
		return new(big.Float).SetInt(v)
	default:
		panic("unsupported big numeric type")
	}
}

// TODO(rsned): Consider adding a variation of Correlate that takes slices of string
// values that represent numbers. (To allow for passing in values in scientific
// notation, or in a format that is not easily converted to a number.) In this
// method, detect when to switch to math/big values based on parsing the strings.
// TODO(rsned): Once string support is in, i18n the method to handle numbers
// formatted in different styles, e.g. ' ', or '.' as the Thousands Separator, etc.
