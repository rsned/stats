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

import (
	"math"
	"testing"
)

func TestDatasets(t *testing.T) {
	// Test AnscombeI dataset
	t.Run("AnscombeI dataset", func(t *testing.T) {
		if len(AnscombeI.X) != len(AnscombeI.Y) {
			t.Errorf("AnscombeI X and Y slices have different lengths: %d vs %d", len(AnscombeI.X), len(AnscombeI.Y))
		}

		if len(AnscombeI.X) != 11 {
			t.Errorf("AnscombeI should have 11 data points, got %d", len(AnscombeI.X))
		}

		if AnscombeI.Name == "" {
			t.Error("AnscombeI should have a name")
		}

		if AnscombeI.Description == "" {
			t.Error("AnscombeI should have a description")
		}

		// Verify some statistical properties of Anscombe's Quartet
		// Mean of X should be approximately 9
		sumX := 0.0
		for _, x := range AnscombeI.X {
			sumX += x
		}
		meanX := sumX / float64(len(AnscombeI.X))
		if math.Abs(meanX-9.0) > 0.01 {
			t.Errorf("AnscombeI mean of X = %v, expected ≈ 9.0", meanX)
		}

		// Mean of Y should be approximately 7.5
		sumY := 0.0
		for _, y := range AnscombeI.Y {
			sumY += y
		}
		meanY := sumY / float64(len(AnscombeI.Y))
		if math.Abs(meanY-7.5) > 0.1 {
			t.Errorf("AnscombeI mean of Y = %v, expected ≈ 7.5", meanY)
		}
	})

	// Test DatasaurusDino dataset
	t.Run("DatasaurusDino dataset", func(t *testing.T) {
		if len(DatasaurusDino.X) != len(DatasaurusDino.Y) {
			t.Errorf("DatasaurusDino X and Y slices have different lengths: %d vs %d", len(DatasaurusDino.X), len(DatasaurusDino.Y))
		}

		if len(DatasaurusDino.X) == 0 {
			t.Error("DatasaurusDino should have data points")
		}

		if DatasaurusDino.Name == "" {
			t.Error("DatasaurusDino should have a name")
		}

		if DatasaurusDino.Description == "" {
			t.Error("DatasaurusDino should have a description")
		}

		// Verify some statistical properties of the Datasaurus Dozen subset
		// Mean of X should be reasonable for this subset
		sumX := 0.0
		for _, x := range DatasaurusDino.X {
			sumX += x
		}
		meanX := sumX / float64(len(DatasaurusDino.X))
		if meanX < 20 || meanX > 80 {
			t.Errorf("DatasaurusDino mean of X = %v, expected to be in reasonable range [20-80]", meanX)
		}

		// Mean of Y should be reasonable for this subset
		sumY := 0.0
		for _, y := range DatasaurusDino.Y {
			sumY += y
		}
		meanY := sumY / float64(len(DatasaurusDino.Y))
		if meanY < 40 || meanY > 100 {
			t.Errorf("DatasaurusDino mean of Y = %v, expected to be in reasonable range [40-100]", meanY)
		}

		t.Logf("DatasaurusDino has %d data points", len(DatasaurusDino.X))
		t.Logf("DatasaurusDino mean X: %.2f, mean Y: %.2f", meanX, meanY)
	})
}

func TestDatasetType(t *testing.T) {
	// Test that we can create a custom Dataset
	custom := Dataset{
		X:           []float64{1, 2, 3, 4, 5},
		Y:           []float64{2, 4, 6, 8, 10},
		Name:        "Linear Test",
		Description: "A simple linear relationship for testing",
		Attribution: "",
	}

	if len(custom.X) != 5 {
		t.Errorf("Custom dataset X should have 5 elements, got %d", len(custom.X))
	}

	if len(custom.Y) != 5 {
		t.Errorf("Custom dataset Y should have 5 elements, got %d", len(custom.Y))
	}

	if custom.Name != "Linear Test" {
		t.Errorf("Custom dataset name = %q, expected 'Linear Test'", custom.Name)
	}
}

func TestDatasetsType(t *testing.T) {
	// Test the predefined ExampleDatasets collection
	t.Run("ExampleDatasets collection", func(t *testing.T) {
		if ExampleDatasets.Name == "" {
			t.Error("ExampleDatasets should have a name")
		}

		if ExampleDatasets.Description == "" {
			t.Error("ExampleDatasets should have a description")
		}

		if len(ExampleDatasets.Data) != 15 {
			t.Errorf("ExampleDatasets should contain 15 datasets, got %d", len(ExampleDatasets.Data))
		}

		// Verify the datasets are correctly included
		expectedDatasets := []string{
			"Anscombe I",
			"Anscombe II",
			"Anscombe III",
			"Anscombe IV",
			"Datasaurus Dozen - Dino",
			"Datasaurus Dozen - Slant Down",
			"Datasaurus Dozen - Slant Up",
			"Datasaurus Dozen - Wide Lines",
			"Datasaurus Dozen - H Lines",
			"Datasaurus Dozen - V Lines",
			"Datasaurus Dozen - X Shape",
			"Datasaurus Dozen - Star",
			"Datasaurus Dozen - High Lines",
			"Datasaurus Dozen - Dots",
			"Datasaurus Dozen - Circle",
		}

		for _, expected := range expectedDatasets {
			found := false
			for _, dataset := range ExampleDatasets.Data {
				if dataset.Name == expected {
					found = true

					break
				}
			}
			if !found {
				t.Errorf("ExampleDatasets should contain dataset '%s'", expected)
			}
		}

		t.Logf("ExampleDatasets contains %d datasets: %s", len(ExampleDatasets.Data), ExampleDatasets.Name)
	})

	// Test creating a custom Datasets collection
	t.Run("custom datasets collection", func(t *testing.T) {
		dataset1 := Dataset{
			X:           []float64{1, 2, 3},
			Y:           []float64{2, 4, 6},
			Name:        "Linear",
			Description: "Perfect linear relationship",
			Attribution: "",
		}

		dataset2 := Dataset{
			X:           []float64{1, 2, 3},
			Y:           []float64{6, 4, 2},
			Name:        "Inverse",
			Description: "Perfect inverse relationship",
			Attribution: "",
		}

		customCollection := Datasets{
			Name:        "Test Collection",
			Description: "A collection of test datasets for validation",
			Attribution: "",
			Data:        []Dataset{dataset1, dataset2},
		}

		if len(customCollection.Data) != 2 {
			t.Errorf("Custom collection should have 2 datasets, got %d", len(customCollection.Data))
		}

		if customCollection.Name != "Test Collection" {
			t.Errorf("Custom collection name = %q, expected 'Test Collection'", customCollection.Name)
		}

		// Test accessing individual datasets from collection
		if customCollection.Data[0].Name != "Linear" {
			t.Errorf("First dataset name = %q, expected 'Linear'", customCollection.Data[0].Name)
		}

		if customCollection.Data[1].Name != "Inverse" {
			t.Errorf("Second dataset name = %q, expected 'Inverse'", customCollection.Data[1].Name)
		}
	})

	// Test empty collection
	t.Run("empty datasets collection", func(t *testing.T) {
		emptyCollection := Datasets{
			Name:        "Empty Collection",
			Description: "A collection with no datasets",
			Attribution: "",
			Data:        []Dataset{},
		}

		if len(emptyCollection.Data) != 0 {
			t.Errorf("Empty collection should have 0 datasets, got %d", len(emptyCollection.Data))
		}

		if emptyCollection.Name != "Empty Collection" {
			t.Errorf("Empty collection name = %q, expected 'Empty Collection'", emptyCollection.Name)
		}
	})
}
