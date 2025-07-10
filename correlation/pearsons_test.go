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
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"testing"
)

func TestPearsonSinglePassVsTwoPass(t *testing.T) {
	tests := []struct {
		name string
		x    []float64
		y    []float64
	}{
		{
			name: "perfect positive correlation",
			x:    []float64{1, 2, 3, 4, 5},
			y:    []float64{2, 4, 6, 8, 10},
		},
		{
			name: "perfect negative correlation",
			x:    []float64{1, 2, 3, 4, 5},
			y:    []float64{10, 8, 6, 4, 2},
		},
		{
			name: "moderate positive correlation",
			x:    []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			y:    []float64{1.5, 2.2, 2.8, 4.1, 4.9, 6.2, 7.1, 7.8, 9.2, 10.1},
		},
		{
			name: "negative numbers",
			x:    []float64{-5, -3, -1, 1, 3, 5},
			y:    []float64{-10, -6, -2, 2, 6, 10},
		},
		{
			name: "floating point precision",
			x:    []float64{1.0000001, 2.0000001, 3.0000001},
			y:    []float64{2.0000002, 4.0000002, 6.0000002},
		},
		{
			name: "real world example",
			x:    []float64{43, 21, 25, 42, 57, 59},
			y:    []float64{99, 65, 79, 75, 87, 81},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultSinglePass, errSinglePass := pearsonsSinglePass(tt.x, tt.y)
			resultTwoPass, errTwoPass := pearsonsTwoPass(tt.x, tt.y)

			// Check that both methods have the same error status
			if (errSinglePass == nil) != (errTwoPass == nil) {
				t.Errorf("Error status mismatch: singlePass err=%v, twoPass err=%v", errSinglePass, errTwoPass)

				return
			}

			// If both have errors, they should be the same type
			if errSinglePass != nil && errTwoPass != nil {
				if errSinglePass.Error() != errTwoPass.Error() {
					t.Errorf("Different error messages: singlePass=%v, twoPass=%v", errSinglePass, errTwoPass)
				}

				return
			}

			// Compare results with small tolerance for floating point differences
			tolerance := 1e-10
			if math.Abs(resultSinglePass-resultTwoPass) > tolerance {
				t.Errorf("Results differ: singlePass=%v, twoPass=%v, diff=%v",
					resultSinglePass, resultTwoPass, math.Abs(resultSinglePass-resultTwoPass))
			}

			t.Logf("singlePass: %.15f, twoPass: %.15f", resultSinglePass, resultTwoPass)
		})
	}
}

func TestPearsonSinglePassVsTwoPassRandomData(t *testing.T) {
	rng := rand.New(rand.NewSource(getSeed()))

	// Test with various sizes of random data
	sizes := []int{2, 10, 100, 1000, 10000}

	for _, size := range sizes {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			x := make([]float64, size)
			y := make([]float64, size)

			for i := 0; i < size; i++ {
				x[i] = rng.Float64()*200 - 100 // Random values between -100 and 100
				y[i] = rng.Float64()*200 - 100
			}

			resultSinglePass, errSinglePass := pearsonsSinglePass(x, y)
			resultTwoPass, errTwoPass := pearsonsTwoPass(x, y)

			// Check that both methods have the same error status
			if (errSinglePass == nil) != (errTwoPass == nil) {
				t.Errorf("Error status mismatch: singlePass err=%v, twoPass err=%v", errSinglePass, errTwoPass)

				return
			}

			if errSinglePass != nil {
				return // Both have errors, that's fine for random data
			}

			// Compare results with tolerance appropriate for the data size
			tolerance := 1e-10
			if math.Abs(resultSinglePass-resultTwoPass) > tolerance {
				t.Errorf("Results differ for size %d: singlePass=%v, twoPass=%v, diff=%v",
					size, resultSinglePass, resultTwoPass, math.Abs(resultSinglePass-resultTwoPass))
			}
		})
	}
}

func TestPearsonVsPearsonBig(t *testing.T) {
	tests := []struct {
		name string
		x    []float64
		y    []float64
	}{
		{
			name: "perfect positive correlation",
			x:    []float64{1, 2, 3, 4, 5},
			y:    []float64{2, 4, 6, 8, 10},
		},
		{
			name: "perfect negative correlation",
			x:    []float64{1, 2, 3, 4, 5},
			y:    []float64{10, 8, 6, 4, 2},
		},
		{
			name: "moderate positive correlation",
			x:    []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			y:    []float64{1.5, 2.2, 2.8, 4.1, 4.9, 6.2, 7.1, 7.8, 9.2, 10.1},
		},
		{
			name: "negative numbers",
			x:    []float64{-5, -3, -1, 1, 3, 5},
			y:    []float64{-10, -6, -2, 2, 6, 10},
		},
		{
			name: "real world example",
			x:    []float64{43, 21, 25, 42, 57, 59},
			y:    []float64{99, 65, 79, 75, 87, 81},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Calculate using regular Pearson
			resultFloat, errFloat := Pearsons(tt.x, tt.y)

			// Convert to big.Float and calculate using PearsonBig
			xBig := make([]*big.Float, len(tt.x))
			yBig := make([]*big.Float, len(tt.y))
			for i := range tt.x {
				xBig[i] = big.NewFloat(tt.x[i])
				yBig[i] = big.NewFloat(tt.y[i])
			}
			resultBig, errBig := PearsonsBig(xBig, yBig)

			// Check that both methods have the same error status
			if (errFloat == nil) != (errBig == nil) {
				t.Errorf("Error status mismatch: float err=%v, big err=%v", errFloat, errBig)

				return
			}

			// If both have errors, they should be similar
			if errFloat != nil && errBig != nil {
				return // Both failed, that's acceptable for this test
			}

			// Compare results with small tolerance for floating point differences
			tolerance := 1e-10
			if math.Abs(resultFloat-resultBig) > tolerance {
				t.Errorf("Results differ: float=%v, big=%v, diff=%v",
					resultFloat, resultBig, math.Abs(resultFloat-resultBig))
			}

			t.Logf("float: %.15f, big: %.15f", resultFloat, resultBig)
		})
	}
}

// TestPearsonsBigInfinityHandling tests that PearsonsBig properly handles
// infinite values without causing panics
func TestPearsonsBigInfinityHandling(t *testing.T) {
	infFloat := big.NewFloat(0)
	infFloat.SetInf(false)

	tests := []struct {
		name        string
		x           []*big.Float
		y           []*big.Float
		expectError bool
		errorMsg    string
	}{
		{
			name: "both infinite with same sign",
			x: []*big.Float{
				big.NewFloat(1),
				big.NewFloat(2),
				infFloat, // +Inf
			},
			y: []*big.Float{
				big.NewFloat(1),
				big.NewFloat(4),
				infFloat, // +Inf
			},
			expectError: true,
			errorMsg:    "correlation undefined: infinite values with same sign detected",
		},
		{
			name: "infinite variance in X",
			x: []*big.Float{
				infFloat, // +Inf
				big.NewFloat(1),
				big.NewFloat(2),
			},
			y: []*big.Float{
				big.NewFloat(1),
				big.NewFloat(2),
				big.NewFloat(3),
			},
			expectError: true,
			errorMsg:    "correlation undefined: infinite values with same sign detected", // This is what actually happens first
		},
		{
			name: "finite values work normally",
			x: []*big.Float{
				big.NewFloat(1),
				big.NewFloat(2),
				big.NewFloat(3),
			},
			y: []*big.Float{
				big.NewFloat(2),
				big.NewFloat(4),
				big.NewFloat(6),
			},
			expectError: false,
			errorMsg:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			corr, err := PearsonsBig(tt.x, tt.y)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error %q, got none", tt.errorMsg)
				} else if err.Error() != tt.errorMsg {
					t.Errorf("expected error %q, got %q", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				// For finite values, should get reasonable correlation
				if math.IsNaN(corr) || math.IsInf(corr, 0) {
					t.Errorf("expected finite correlation, got %v", corr)
				}
			}
		})
	}
}

func BenchmarkPearsonFloat(b *testing.B) {
	const limit = 10000
	x := make([]float64, limit)
	y := make([]float64, limit)
	rng := rand.New(rand.NewSource(getSeed()))
	for i := range limit {
		x[i] = rng.Float64() * 1000
		y[i] = rng.Float64() * 100
	}

	for b.Loop() {
		_, _ = Pearsons(x, y)
	}
}

func BenchmarkPearsonSinglePass(b *testing.B) {
	const limit = 10000
	x := make([]float64, limit)
	y := make([]float64, limit)
	rng := rand.New(rand.NewSource(getSeed()))
	for i := range limit {
		x[i] = rng.Float64() * 1000
		y[i] = rng.Float64() * 100
	}

	for b.Loop() {
		_, _ = pearsonsSinglePass(x, y)
	}
}

func BenchmarkPearsonTwoPass(b *testing.B) {
	const limit = 10000
	x := make([]float64, limit)
	y := make([]float64, limit)
	rng := rand.New(rand.NewSource(getSeed()))
	for i := range limit {
		x[i] = rng.Float64() * 1000
		y[i] = rng.Float64() * 100
	}

	for b.Loop() {
		_, _ = pearsonsTwoPass(x, y)
	}
}

func BenchmarkPearsonsBig(b *testing.B) {
	const limit = 10000
	x := make([]*big.Float, limit)
	y := make([]*big.Float, limit)
	rng := rand.New(rand.NewSource(getSeed()))
	for i := range limit {
		x[i] = big.NewFloat(rng.Float64() * 1000)
		y[i] = big.NewFloat(rng.Float64() * 100)
	}

	for b.Loop() {
		_, _ = PearsonsBig(x, y)
	}
}
