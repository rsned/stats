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

// GoodmanKruskals calculates Goodman and Kruskal's gamma correlation coefficient.
// Gamma is a rank-based measure of association that ranges from -1 to +1.
//
// Unlike Kendall's Tau, Gamma ignores tied pairs entirely in the calculation,
// making it more suitable when there are many tied values.
//
// It returns a value between -1 and 1, where:
//   - 1 indicates a perfect positive monotonic relationship
//   - 0 indicates no monotonic relationship
//   - -1 indicates a perfect negative monotonic relationship
//
// An error is returned if the slices have different lengths or are empty.
func GoodmanKruskals[T Numeric](x, y []T) (float64, error) {
	return 0, errors.New("not implemented")
}

// GoodmanKruskalsBig calculates Goodman and Kruskal's gamma correlation coefficient
// using big number types (*big.Float or *big.Int).
//
// Gamma is a rank-based measure of association that ranges from -1 to +1.
// Unlike Kendall's Tau, Gamma ignores tied pairs entirely in the calculation.
func GoodmanKruskalsBig[T BigNumeric](x, y []T) (float64, error) {
	return 0, errors.New("not implemented")
}

// GoodmanKruskalsMixed calculates Goodman and Kruskal's gamma correlation coefficient
// with mixed type inputs.
//
// Gamma is a rank-based measure of association that ranges from -1 to +1.
// Unlike Kendall's Tau, Gamma ignores tied pairs entirely in the calculation.
// It converts the inputs using mixedToBig and then calls GoodmanKruskalsBig.
func GoodmanKruskalsMixed[T MixedNumeric](x, y []T) (float64, error) {
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

	return GoodmanKruskalsBig(xVals, yVals)
}
