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

package datasets

// Dataset represents a pair of related data series for statistical analysis.
// It contains two slices of numeric values, typically used for correlation
// analysis, regression, or other bivariate statistical operations.
type Dataset struct {
	// Name provides a descriptive name for the dataset
	Name string
	// Description provides additional context about the dataset
	Description string
	// Attribution provides reference to the authoritative source for this dataset
	Attribution string
	// X contains the independent variable values
	X []float64
	// Y contains the dependent variable values
	Y []float64
}

// Datasets represents a collection of related datasets with metadata.
// This type is useful for grouping multiple datasets that share common
// characteristics or are part of the same study or analysis.
type Datasets struct {
	// Name provides a descriptive name for the collection
	Name string
	// Description provides additional context about the collection
	Description string
	// Attribution provides reference to the authoritative source for this collection
	Attribution string
	// Data contains the slice of datasets in this collection
	Data []Dataset
}
