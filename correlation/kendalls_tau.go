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
)

// KendallsTau calculates Kendall's Tau correlation coefficient
// between two datasets x and y of any numeric type.
//
// Kendall's Tau measures the ordinal association between two measured quantities.
// It is based on the number of concordant and discordant pairs in the data.
//
// It returns a value between -1 and 1, where:
//   - 1 indicates a perfect positive monotonic relationship
//   - 0 indicates no monotonic relationship
//   - -1 indicates a perfect negative monotonic relationship
//
// An error is returned if the slices have different lengths or are empty.
func KendallsTau[T Numeric](x, y []T) (float64, error) {
	return 0, errors.New("not implemented")
}

// KendallsTauBig calculates Kendall's Tau correlation coefficient
// between two datasets x and y of big number types (*big.Float or *big.Int).
//
// Kendall's Tau measures the ordinal association between two measured quantities.
// It is based on the number of concordant and discordant pairs in the data.

// KendallsTauBig calculates Kendall's Tau correlation coefficient
// between two datasets x and y of big number types (*big.Float or *big.Int).
//
// Kendall's Tau measures the ordinal association between two measured quantities.
// It is based on the number of concordant and discordant pairs in the data.
func KendallsTauBig[T BigNumeric](x, y []T) (float64, error) {
	return 0, errors.New("not implemented")
}

// KendallsTauMixed calculates Kendall's Tau correlation coefficient
// between two datasets x and y with a set of mixed type inputs.
//
// Kendall's Tau measures the ordinal association between two measured quantities.
// It is based on the number of concordant and discordant pairs in the data.
// It converts the inputs using mixedToBig and then calls KendallsTauBig.
func KendallsTauMixed[T MixedNumeric](x, y []T) (float64, error) {
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

	return KendallsTauBig(xVals, yVals)
}
