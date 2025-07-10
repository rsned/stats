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

	x := stringsToBigFloat([]string{"1e1000", "2e1000", "3e1000", "4e1000", "5e1000"})
	y := stringsToBigFloat([]string{"3e1000", "6e1000", "7e1000", "3e1000", "9e1000"})

	result, err := CorrelateBig(x, y, correlate.Pearson)  // result will be ~0.545705
*/
package correlation
