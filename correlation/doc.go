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

/*
Package correlation holds routines for performing the various standard
types of correlations on data in both numeric and big.Float types.

The simplest use is to pass your two sets of data into Correlate()
and choose which type of correlation to use on the data.

Currently supported correlation types are:
  - Pearson's product-moment correlation
  - Kendall's tau rank correlation
  - Spearman's rank correlation
  - Goodman and Kruskals Gamma rank correlation

For example:

	x:= []float64{43, 21, 25, 42, 57, 59},
	y:= []float64{99, 65, 79, 75, 87, 81},

	result, err := Correlate(x, y, correlation.Pearson)  // result will be ~0.529

To use with big.Floats:

	x:= []*big.Float{
		big.NewFloat(1e30),
		big.NewFloat(2e74),
		big.NewFloat(3e150),
		big.NewFloat(4e311),
		big.NewFloat(5e500),
	}
	y:= []*big.Float{
		big.NewFloat(2e30),
		big.NewFloat(4e75),
		big.NewFloat(6e150),
		big.NewFloat(8e310),
		big.NewFloat(10e502),
	}
	result, err := CorrelateBig(x, y, correlate.Pearson)
*/
package correlation
