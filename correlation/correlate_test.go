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
	"flag"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strings"
	"testing"

	"github.com/rsned/stats/datasets"
)

const defaultSeed = int64(42)

// seed is the seed for random number generation in tests.
var seed = flag.Int64("seed", defaultSeed, "Random seed for correlation tests")

// getSeed returns the configured seed value for random number generation.
func getSeed() int64 {
	if seed == nil {
		return defaultSeed
	}

	return *seed
}

func TestCorrelate(t *testing.T) {
	tests := []struct {
		name     string
		x        []float64
		y        []float64
		expected float64
		wantErr  bool
	}{
		{
			name:     "perfect positive correlation",
			x:        []float64{1, 2, 3, 4, 5},
			y:        []float64{2, 4, 6, 8, 10},
			expected: 1.0,
			wantErr:  false,
		},
		{
			name:     "perfect negative correlation",
			x:        []float64{1, 2, 3, 4, 5},
			y:        []float64{10, 8, 6, 4, 2},
			expected: -1.0,
			wantErr:  false,
		},
		{
			name:     "no correlation",
			x:        []float64{1, 2, 3, 4, 5},
			y:        []float64{5, 5, 5, 5, 5},
			expected: 0,
			wantErr:  true, // zero variance in y
		},
		{
			name:     "moderate positive correlation",
			x:        []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			y:        []float64{1.5, 2.2, 2.8, 4.1, 4.9, 6.2, 7.1, 7.8, 9.2, 10.1},
			expected: 0.998,
			wantErr:  false,
		},
		{
			name:     "empty slices",
			x:        []float64{},
			y:        []float64{},
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "different lengths",
			x:        []float64{1, 2, 3},
			y:        []float64{1, 2},
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "single element",
			x:        []float64{1},
			y:        []float64{1},
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "zero variance in x",
			x:        []float64{3, 3, 3, 3},
			y:        []float64{1, 2, 3, 4},
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "real world example",
			x:        []float64{43, 21, 25, 42, 57, 59},
			y:        []float64{99, 65, 79, 75, 87, 81},
			expected: 0.529,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Correlate(tt.x, tt.y, Pearson)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Correlate() expected error but got none")
				}

				return
			}

			if err != nil {
				t.Errorf("Correlate() unexpected error: %v", err)

				return
			}

			if math.Abs(result-tt.expected) > 0.001 {
				t.Errorf("Correlate() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestCorrelateBig(t *testing.T) {
	tests := []struct {
		name     string
		x        []*big.Float
		y        []*big.Float
		expected float64
		wantErr  bool
	}{
		{
			name:     "perfect positive correlation",
			x:        []*big.Float{big.NewFloat(1), big.NewFloat(2), big.NewFloat(3), big.NewFloat(4), big.NewFloat(5)},
			y:        []*big.Float{big.NewFloat(2), big.NewFloat(4), big.NewFloat(6), big.NewFloat(8), big.NewFloat(10)},
			expected: 1.0,
			wantErr:  false,
		},
		{
			name: "large numbers exceeding float64 limits",
			x: func() []*big.Float {
				x := make([]*big.Float, 5)
				x[0], _ = new(big.Float).SetString("1e400") // Exceeds float64 max (~1.8e308)
				x[1], _ = new(big.Float).SetString("2e400")
				x[2], _ = new(big.Float).SetString("3e400")
				x[3], _ = new(big.Float).SetString("4e400")
				x[4], _ = new(big.Float).SetString("5e400")

				return x
			}(),
			y: func() []*big.Float {
				y := make([]*big.Float, 5)
				y[0], _ = new(big.Float).SetString("2e400") // Perfect linear correlation
				y[1], _ = new(big.Float).SetString("4e400")
				y[2], _ = new(big.Float).SetString("6e400")
				y[3], _ = new(big.Float).SetString("8e400")
				y[4], _ = new(big.Float).SetString("10e400")

				return y
			}(),
			expected: 1.0,
			wantErr:  false,
		},
		{
			name: "tiny numbers below float64 precision",
			x: func() []*big.Float {
				x := make([]*big.Float, 5)
				x[0], _ = new(big.Float).SetString("1e-400") // Below float64 min (~2.2e-308)
				x[1], _ = new(big.Float).SetString("2e-400")
				x[2], _ = new(big.Float).SetString("3e-400")
				x[3], _ = new(big.Float).SetString("4e-400")
				x[4], _ = new(big.Float).SetString("5e-400")

				return x
			}(),
			y: func() []*big.Float {
				y := make([]*big.Float, 5)
				y[0], _ = new(big.Float).SetString("2e-400") // Perfect linear correlation
				y[1], _ = new(big.Float).SetString("4e-400")
				y[2], _ = new(big.Float).SetString("6e-400")
				y[3], _ = new(big.Float).SetString("8e-400")
				y[4], _ = new(big.Float).SetString("10e-400")

				return y
			}(),
			expected: 1.0,
			wantErr:  false,
		},
		{
			name: "large numbers with moderate positive correlation",
			x: func() []*big.Float {
				x := make([]*big.Float, 6)
				x[0], _ = new(big.Float).SetString("1e350")
				x[1], _ = new(big.Float).SetString("2e350")
				x[2], _ = new(big.Float).SetString("3e350")
				x[3], _ = new(big.Float).SetString("4e350")
				x[4], _ = new(big.Float).SetString("5e350")
				x[5], _ = new(big.Float).SetString("6e350")

				return x
			}(),
			y: func() []*big.Float {
				y := make([]*big.Float, 6)
				y[0], _ = new(big.Float).SetString("1.5e350") // Not perfect correlation
				y[1], _ = new(big.Float).SetString("3.8e350")
				y[2], _ = new(big.Float).SetString("6.2e350")
				y[3], _ = new(big.Float).SetString("8.1e350")
				y[4], _ = new(big.Float).SetString("9.9e350")
				y[5], _ = new(big.Float).SetString("12.3e350")

				return y
			}(),
			expected: 0.999, // Strong but not perfect correlation
			wantErr:  false,
		},
		{
			name: "tiny numbers with negative correlation",
			x: func() []*big.Float {
				x := make([]*big.Float, 5)
				x[0], _ = new(big.Float).SetString("1e-350")
				x[1], _ = new(big.Float).SetString("2e-350")
				x[2], _ = new(big.Float).SetString("3e-350")
				x[3], _ = new(big.Float).SetString("4e-350")
				x[4], _ = new(big.Float).SetString("5e-350")

				return x
			}(),
			y: func() []*big.Float {
				y := make([]*big.Float, 5)
				y[0], _ = new(big.Float).SetString("10e-350") // Perfect negative correlation
				y[1], _ = new(big.Float).SetString("8e-350")
				y[2], _ = new(big.Float).SetString("6e-350")
				y[3], _ = new(big.Float).SetString("4e-350")
				y[4], _ = new(big.Float).SetString("2e-350")

				return y
			}(),
			expected: -1.0,
			wantErr:  false,
		},
		{
			name: "mixed scale numbers with weak correlation",
			x: func() []*big.Float {
				x := make([]*big.Float, 7)
				x[0], _ = new(big.Float).SetString("1e400")
				x[1], _ = new(big.Float).SetString("3e400")
				x[2], _ = new(big.Float).SetString("2e400")
				x[3], _ = new(big.Float).SetString("6e400")
				x[4], _ = new(big.Float).SetString("4e400")
				x[5], _ = new(big.Float).SetString("7e400")
				x[6], _ = new(big.Float).SetString("5e400")

				return x
			}(),
			y: func() []*big.Float {
				y := make([]*big.Float, 7)
				y[0], _ = new(big.Float).SetString("2.1e-400") // Very weak correlation pattern
				y[1], _ = new(big.Float).SetString("3.8e-400")
				y[2], _ = new(big.Float).SetString("2.9e-400")
				y[3], _ = new(big.Float).SetString("5.2e-400")
				y[4], _ = new(big.Float).SetString("4.1e-400")
				y[5], _ = new(big.Float).SetString("5.9e-400")
				y[6], _ = new(big.Float).SetString("4.8e-400")

				return y
			}(),
			expected: 0.993, // Strong positive correlation
			wantErr:  false,
		},
		{
			name: "ultra-large numbers with scatter pattern",
			x: func() []*big.Float {
				x := make([]*big.Float, 8)
				x[0], _ = new(big.Float).SetString("1e500")
				x[1], _ = new(big.Float).SetString("2e500")
				x[2], _ = new(big.Float).SetString("3e500")
				x[3], _ = new(big.Float).SetString("4e500")
				x[4], _ = new(big.Float).SetString("5e500")
				x[5], _ = new(big.Float).SetString("6e500")
				x[6], _ = new(big.Float).SetString("7e500")
				x[7], _ = new(big.Float).SetString("8e500")

				return x
			}(),
			y: func() []*big.Float {
				y := make([]*big.Float, 8)
				y[0], _ = new(big.Float).SetString("8e500") // Scattered pattern
				y[1], _ = new(big.Float).SetString("3e500")
				y[2], _ = new(big.Float).SetString("7e500")
				y[3], _ = new(big.Float).SetString("2e500")
				y[4], _ = new(big.Float).SetString("6e500")
				y[5], _ = new(big.Float).SetString("1e500")
				y[6], _ = new(big.Float).SetString("5e500")
				y[7], _ = new(big.Float).SetString("4e500")

				return y
			}(),
			expected: -0.38, // Moderate negative correlation
			wantErr:  false,
		},
		{
			name: "large numbers with negligible correlation",
			x: func() []*big.Float {
				x := make([]*big.Float, 6)
				x[0], _ = new(big.Float).SetString("1e450")
				x[1], _ = new(big.Float).SetString("2e450")
				x[2], _ = new(big.Float).SetString("3e450")
				x[3], _ = new(big.Float).SetString("4e450")
				x[4], _ = new(big.Float).SetString("5e450")
				x[5], _ = new(big.Float).SetString("6e450")

				return x
			}(),
			y: func() []*big.Float {
				y := make([]*big.Float, 6)
				y[0], _ = new(big.Float).SetString("5e-450") // Designed for zero correlation
				y[1], _ = new(big.Float).SetString("2e-450")
				y[2], _ = new(big.Float).SetString("4e-450")
				y[3], _ = new(big.Float).SetString("1e-450")
				y[4], _ = new(big.Float).SetString("3e-450")
				y[5], _ = new(big.Float).SetString("6e-450")

				return y
			}(),
			expected: 0.143, // Near-zero correlation
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CorrelateBig(tt.x, tt.y, Pearson)

			if tt.wantErr {
				if err == nil {
					t.Errorf("CorrelateBig() expected error but got none")
				}

				return
			}

			if err != nil {
				t.Errorf("CorrelateBig() unexpected error: %v", err)

				return
			}

			if math.Abs(result-tt.expected) > 0.001 {
				t.Errorf("CorrelateBig() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestCorrelateMixed(t *testing.T) {
	// Test with int slices
	t.Run("int types", func(t *testing.T) {
		x := []int{1, 2, 3, 4, 5}
		y := []int{2, 4, 6, 8, 10}
		result, err := CorrelateMixed(x, y, Pearson)
		if err != nil {
			t.Errorf("CorrelateMixed() with int failed: %v", err)
		}
		if math.Abs(result-1.0) > 0.001 {
			t.Errorf("CorrelateMixed() with int = %v, expected 1.0", result)
		}
	})

	// Test with int32 slices
	t.Run("int32 types", func(t *testing.T) {
		x := []int32{1, 2, 3, 4, 5}
		y := []int32{10, 8, 6, 4, 2}
		result, err := CorrelateMixed(x, y, Pearson)
		if err != nil {
			t.Errorf("CorrelateMixed() with int32 failed: %v", err)
		}
		if math.Abs(result-(-1.0)) > 0.001 {
			t.Errorf("CorrelateMixed() with int32 = %v, expected -1.0", result)
		}
	})

	// Test with uint64 slices
	t.Run("uint64 types", func(t *testing.T) {
		x := []uint64{10, 20, 30, 40, 50}
		y := []uint64{5, 10, 15, 20, 25}
		result, err := CorrelateMixed(x, y, Pearson)
		if err != nil {
			t.Errorf("CorrelateMixed() with uint64 failed: %v", err)
		}
		if math.Abs(result-1.0) > 0.001 {
			t.Errorf("CorrelateMixed() with uint64 = %v, expected 1.0", result)
		}
	})

	// Test with float32 slices
	t.Run("float32 types", func(t *testing.T) {
		x := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
		y := []float32{2.2, 4.4, 6.6, 8.8, 11.0}
		result, err := CorrelateMixed(x, y, Pearson)
		if err != nil {
			t.Errorf("CorrelateMixed() with float32 failed: %v", err)
		}
		if math.Abs(result-1.0) > 0.001 {
			t.Errorf("CorrelateMixed() with float32 = %v, expected 1.0", result)
		}
	})

	// Test with float64 slices
	t.Run("float64 types", func(t *testing.T) {
		x := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
		y := []float64{2.0, 4.0, 6.0, 8.0, 10.0}
		result, err := CorrelateMixed(x, y, Pearson)
		if err != nil {
			t.Errorf("CorrelateMixed() with float64 failed: %v", err)
		}
		if math.Abs(result-1.0) > 0.001 {
			t.Errorf("CorrelateMixed() with float64 = %v, expected 1.0", result)
		}
	})

	// Test with *big.Float slices
	t.Run("big.Float types", func(t *testing.T) {
		x := []*big.Float{big.NewFloat(1.0), big.NewFloat(2.0), big.NewFloat(3.0)}
		y := []*big.Float{big.NewFloat(2.0), big.NewFloat(4.0), big.NewFloat(6.0)}
		result, err := CorrelateMixed(x, y, Pearson)
		if err != nil {
			t.Errorf("CorrelateMixed() with big.Float failed: %v", err)
		}
		if math.Abs(result-1.0) > 0.001 {
			t.Errorf("CorrelateMixed() with big.Float = %v, expected 1.0", result)
		}
	})

	// Test with *big.Int slices
	t.Run("big.Int types", func(t *testing.T) {
		x := []*big.Int{big.NewInt(10), big.NewInt(20), big.NewInt(30)}
		y := []*big.Int{big.NewInt(5), big.NewInt(10), big.NewInt(15)}
		result, err := CorrelateMixed(x, y, Pearson)
		if err != nil {
			t.Errorf("CorrelateMixed() with big.Int failed: %v", err)
		}
		if math.Abs(result-1.0) > 0.001 {
			t.Errorf("CorrelateMixed() with big.Int = %v, expected 1.0", result)
		}
	})

	t.Run("mixed numeric constraint validation", func(t *testing.T) {
		// Test that the function accepts all types in the MixedNumeric interface
		// Using different ranges to verify the constraint system works

		// Test with very large int64 values
		xLargeInt := []int64{1000000000000000, 2000000000000000, 3000000000000000, 4000000000000000}
		ySmallFloat := []float64{2e-10, 4e-10, 6e-10, 8e-10}
		result, err := CorrelateMixed(xLargeInt, ySmallFloat, Pearson)
		if err != nil {
			t.Errorf("CorrelateMixed() with large int64/small float64 failed: %v", err)
		}
		if math.Abs(result-1.0) > 0.001 {
			t.Errorf("CorrelateMixed() with large int64/small float64 = %v, expected 1.0", result)
		}

		// Test with very small float64 values and large ints.
		xSmallFloat := []float64{1e-10, 2e-10, 3e-10, 4e-10}
		yLargeInt := []int64{2000000000000000, 4000000000000000, 6000000000000000, 8000000000000000}
		result2, err := CorrelateMixed(xSmallFloat, yLargeInt, Pearson)
		if err != nil {
			t.Errorf("CorrelateMixed() with small float64/large int64 failed: %v", err)
		}
		if math.Abs(result2-1.0) > 0.001 {
			t.Errorf("CorrelateMixed() with small float64/large int64 = %v, expected 1.0", result2)
		}
	})

	// Note: The current Go generic type system requires both x and y to be the same type []T.
	// Testing with different types for x and y (e.g., []int64 and []*big.Int) is not supported
	// by the CorrelateMixed function signature. To test truly mixed types, you would need
	// separate functions with different type parameters for x and y.

	// Test with mixed types that have different precision characteristics
	t.Run("mixed types with different characteristics", func(t *testing.T) {
		// Test mixing int64 with *big.Int (now possible with updated signature)
		x := []int64{1000000, 2000000, 3000000, 4000000}
		y := []*big.Int{big.NewInt(2000000), big.NewInt(4000000), big.NewInt(6000000), big.NewInt(8000000)}

		result, err := CorrelateMixed(x, y, Pearson)
		if err != nil {
			t.Errorf("CorrelateMixed() with int64/*big.Int failed: %v", err)
		} else if math.Abs(result-1.0) > 0.001 {
			t.Errorf("CorrelateMixed() with int64/*big.Int = %v, expected 1.0", result)
		}

		// Test mixing float64 with *big.Float
		xFloat := []float64{1.5, 2.5, 3.5, 4.5}
		yBigFloat := []*big.Float{big.NewFloat(3.0), big.NewFloat(5.0), big.NewFloat(7.0), big.NewFloat(9.0)}

		result2, err := CorrelateMixed(xFloat, yBigFloat, Pearson)
		if err != nil {
			t.Errorf("CorrelateMixed() with float64/*big.Float failed: %v", err)
		} else if math.Abs(result2-1.0) > 0.001 {
			t.Errorf("CorrelateMixed() with float64/*big.Float = %v, expected 1.0", result2)
		}
	})

	// Test edge cases that might cause issues in the conversion process
	t.Run("edge cases and boundary values", func(t *testing.T) {
		// Test with maximum values that might cause overflow issues
		xMax := []int64{9223372036854775807}   // Max int64
		yMax := []uint64{18446744073709551615} // Max uint64

		_, err := CorrelateMixed(xMax, yMax, Pearson)
		if err == nil {
			t.Errorf("CorrelateMixed() with single max values should fail due to insufficient data points but didn't")
		} else if strings.Contains(err.Error(), "data points") || strings.Contains(err.Error(), "variance") || strings.Contains(err.Error(), "length") {
			// This should fail because we need at least 2 data points for correlation
			t.Logf("CorrelateMixed() failed as expected with max values: %v", err)
		}

		// Test with boundary float values
		xFloat := []float32{1.175494e-38, 3.402823e+38}         // Min and max float32
		yFloat := []float32{2.350988e-38, 3.402823e+38 * 0.999} // Proportional values (stay within limits)

		result, err := CorrelateMixed(xFloat, yFloat, Pearson)
		if err != nil {
			t.Logf("CorrelateMixed() with extreme float32 values failed: %v", err)
		} else if math.Abs(result-1.0) > 0.001 {
			t.Errorf("CorrelateMixed() with extreme float32 = %v, expected 1.0", result)
		}
	})

	// Test that would potentially cause mixedToBig issues with nil pointers
	t.Run("potential mixedToBig edge cases", func(t *testing.T) {
		// Test with nil big.Float pointers (this causes mixedToBig to panic)
		x := []int{1, 2, 3}
		y := make([]*big.Float, 3)
		y[0] = big.NewFloat(1.0)
		y[1] = nil // This nil pointer will cause a panic in mixedToBig
		y[2] = big.NewFloat(3.0)

		// Capture panic to verify mixedToBig fails with nil pointers
		defer func() {
			if r := recover(); r != nil {
				t.Logf("CorrelateMixed() with nil big.Float panicked as expected: %v", r)
			}
		}()

		_, err := CorrelateMixed(x, y, Pearson)
		if err == nil {
			t.Errorf("CorrelateMixed() with nil big.Float should fail but didn't")
		} else {
			t.Logf("CorrelateMixed() with nil big.Float failed as expected: %v", err)
		}
	})

	// Test error cases
	t.Run("empty slices", func(t *testing.T) {
		x := []int{}
		y := []int{}
		_, err := CorrelateMixed(x, y, Pearson)
		if err == nil {
			t.Errorf("CorrelateMixed() with empty slices expected error but got none")
		}
	})

	t.Run("different lengths", func(t *testing.T) {
		x := []int{1, 2, 3}
		y := []int{1, 2}
		_, err := CorrelateMixed(x, y, Pearson)
		if err == nil {
			t.Errorf("CorrelateMixed() with different lengths expected error but got none")
		}
	})
}

func TestCorrelateEdgeCases(t *testing.T) {
	// Test with negative numbers
	x := []float64{-5, -3, -1, 1, 3, 5}
	y := []float64{-10, -6, -2, 2, 6, 10}
	result, err := Correlate(x, y, Pearson)
	if err != nil {
		t.Errorf("Correlate() with negative numbers failed: %v", err)
	}
	if math.Abs(result-1.0) > 0.001 {
		t.Errorf("Correlate() with negative numbers = %v, expected 1.0", result)
	}

	// Test with floating point precision
	x2 := []float64{1.0000001, 2.0000001, 3.0000001}
	y2 := []float64{2.0000002, 4.0000002, 6.0000002}
	result2, err := Correlate(x2, y2, Pearson)
	if err != nil {
		t.Errorf("Correlate() with floating point precision failed: %v", err)
	}
	if math.Abs(result2-1.0) > 0.001 {
		t.Errorf("Correlate() with floating point precision = %v, expected 1.0", result2)
	}
}

func TestCorrelateGenericTypes(t *testing.T) {
	// Test with int
	xInt := []int{1, 2, 3, 4, 5}
	yInt := []int{2, 4, 6, 8, 10}
	result, err := Correlate(xInt, yInt, Pearson)
	if err != nil {
		t.Errorf("Correlate() with int failed: %v", err)
	}
	if math.Abs(result-1.0) > 0.001 {
		t.Errorf("Correlate() with int = %v, expected 1.0", result)
	}

	// Test with int32
	xInt32 := []int32{1, 2, 3, 4, 5}
	yInt32 := []int32{10, 8, 6, 4, 2}
	result32, err := Correlate(xInt32, yInt32, Pearson)
	if err != nil {
		t.Errorf("Correlate() with int32 failed: %v", err)
	}
	if math.Abs(result32-(-1.0)) > 0.001 {
		t.Errorf("Correlate() with int32 = %v, expected -1.0", result32)
	}

	// Test with uint
	xUint := []uint{10, 20, 30, 40, 50}
	yUint := []uint{5, 10, 15, 20, 25}
	resultUint, err := Correlate(xUint, yUint, Pearson)
	if err != nil {
		t.Errorf("Correlate() with uint failed: %v", err)
	}
	if math.Abs(resultUint-1.0) > 0.001 {
		t.Errorf("Correlate() with uint = %v, expected 1.0", resultUint)
	}

	// Test with float32
	xFloat32 := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	yFloat32 := []float32{2.2, 4.4, 6.6, 8.8, 11.0}
	resultFloat32, err := Correlate(xFloat32, yFloat32, Pearson)
	if err != nil {
		t.Errorf("Correlate() with float32 failed: %v", err)
	}
	if math.Abs(resultFloat32-1.0) > 0.001 {
		t.Errorf("Correlate() with float32 = %v, expected 1.0", resultFloat32)
	}

	// Test mixed precision (int64 with large numbers)
	xInt64 := []int64{1000000, 2000000, 3000000, 4000000, 5000000}
	yInt64 := []int64{500000, 1000000, 1500000, 2000000, 2500000}
	resultInt64, err := Correlate(xInt64, yInt64, Pearson)
	if err != nil {
		t.Errorf("Correlate() with int64 failed: %v", err)
	}
	if math.Abs(resultInt64-1.0) > 0.001 {
		t.Errorf("Correlate() with int64 = %v, expected 1.0", resultInt64)
	}
}

func TestCorrelateAnscombesQuartet(t *testing.T) {
	expectedCorr := 0.816
	tolerance := 0.005 // Allow for small floating point differences

	for _, dataset := range datasets.AnscombeQuartet.Data {
		t.Run(dataset.Name, func(t *testing.T) {
			result, err := Correlate(dataset.X, dataset.Y, Pearson)
			if err != nil {
				t.Errorf("Correlate() failed for %s: %v", dataset.Name, err)

				return
			}

			if math.Abs(result-expectedCorr) > tolerance {
				t.Errorf("Correlate() for %s = %v, expected ≈ %v (within %v)",
					dataset.Name, result, expectedCorr, tolerance)
			}

			t.Logf("%s correlation: %.6f", dataset.Name, result)
		})
	}

	// Verify that all correlations are nearly identical despite different distributions
	var correlations []float64
	for _, dataset := range datasets.AnscombeQuartet.Data {
		corr, err := Correlate(dataset.X, dataset.Y, Pearson)
		if err != nil {
			t.Fatalf("Failed to calculate correlation for %s: %v", dataset.Name, err)
		}
		correlations = append(correlations, corr)
	}

	// Check that all correlations are within a small range of each other
	for i := 1; i < len(correlations); i++ {
		diff := math.Abs(correlations[i] - correlations[0])
		if diff > 0.002 { // Very tight tolerance showing they're nearly identical
			t.Errorf("Correlation differences too large between datasets: %.6f vs %.6f (diff: %.6f)",
				correlations[0], correlations[i], diff)
		}
	}
}

func TestCorrelateWithDifferentTypes(t *testing.T) {
	tests := []struct {
		name     string
		x        []float64
		y        []float64
		corrType Type
		expected float64
	}{
		{
			name:     "Pearson",
			x:        []float64{1, 2, 3, 4, 5},
			y:        []float64{2, 4, 6, 8, 10},
			corrType: Pearson,
			expected: 1.0,
		},
		/*
			{
				name:     "Spearman",
				x:        []float64{1, 2, 3, 4, 5},
				y:        []float64{2, 4, 6, 8, 10},
				corrType: Spearman,
				expected: 1.0,
			},
			{
				name:     "KendallTau",
				x:        []float64{1, 2, 3, 4, 5},
				y:        []float64{2, 4, 6, 8, 10},
				corrType: KendallTau,
				expected: 1.0,
			},
			{
				name:     "GoodmanKruskal",
				x:        []float64{1, 2, 3, 4, 5},
				y:        []float64{2, 4, 6, 8, 10},
				corrType: GoodmanKruskal,
				expected: 1.0,
			},
		*/
	}

	for _, test := range tests {
		result, err := Correlate(test.x, test.y, test.corrType)
		if err != nil {
			t.Errorf("%s correlation failed: %v", test.name, err)
		}
		if math.Abs(result-test.expected) > 0.001 {
			t.Errorf("%s correlation = %v, expected %v", test.name, result, test.expected)
		}
	}

	// Test invalid correlation type
	if _, err := Correlate([]float64{1, 2, 3}, []float64{4, 5, 6}, Type(999)); err == nil {
		t.Errorf("Expected error for invalid correlation type but got none")
	}
}

func TestSeedFunctionality(t *testing.T) {
	// This test demonstrates that the seed flag affects random number generation
	// It's mainly for documentation purposes to show the feature works

	seed := getSeed()
	t.Logf("Using seed: %d", seed)

	// Generate some random data using the configured seed
	rng := rand.New(rand.NewSource(seed))
	randomValues := make([]float64, 10)
	for i := range randomValues {
		randomValues[i] = rng.Float64()
	}

	// The test always passes, but logs the seed and values for verification
	t.Logf("Random values with seed %d: %v", seed, randomValues)
}

func TestMixedToBigIntegers(t *testing.T) {
	tests := []struct {
		name      string
		input     any // because we can't use []MixedNumeric
		expected  []float64
		tolerance float64
	}{
		{
			name:      "int types",
			input:     []int{1, 2, 3, 4, 5},
			expected:  []float64{1, 2, 3, 4, 5},
			tolerance: 1e-10,
		},
		{
			name:      "int8 types",
			input:     []int8{10, 20, 30},
			expected:  []float64{10, 20, 30},
			tolerance: 1e-10,
		},
		{
			name:      "int16 types",
			input:     []int16{100, 200, 300},
			expected:  []float64{100, 200, 300},
			tolerance: 1e-10,
		},
		{
			name:      "int32 types",
			input:     []int32{1000, 2000, 3000},
			expected:  []float64{1000, 2000, 3000},
			tolerance: 1e-10,
		},
		{
			name:      "int64 types",
			input:     []int64{1000000000000, 2000000000000, 3000000000000},
			expected:  []float64{1e12, 2e12, 3e12},
			tolerance: 1e6, // Allow some precision loss for large numbers
		},
		{
			name:      "uint types",
			input:     []uint{10, 20, 30},
			expected:  []float64{10, 20, 30},
			tolerance: 1e-10,
		},
		{
			name:      "uint8 types",
			input:     []uint8{50, 100, 150},
			expected:  []float64{50, 100, 150},
			tolerance: 1e-10,
		},
		{
			name:      "uint16 types",
			input:     []uint16{500, 1000, 1500},
			expected:  []float64{500, 1000, 1500},
			tolerance: 1e-10,
		},
		{
			name:      "uint32 types",
			input:     []uint32{5000, 10000, 15000},
			expected:  []float64{5000, 10000, 15000},
			tolerance: 1e-10,
		},
		{
			name:      "uint64 types",
			input:     []uint64{50000000000, 100000000000, 150000000000},
			expected:  []float64{5e10, 1e11, 1.5e11},
			tolerance: 1e4, // Allow some precision loss for large numbers
		},
		{
			name:      "big.Int types",
			input:     []*big.Int{big.NewInt(100), big.NewInt(200), big.NewInt(300)},
			expected:  []float64{100, 200, 300},
			tolerance: 1e-10,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result []*big.Float
			var err error

			// Have to enumerate all types for testing because it can't infer the type from any by default.
			switch v := test.input.(type) {
			case []int:
				result, err = mixedToBig(v)
			case []int8:
				result, err = mixedToBig(v)
			case []int16:
				result, err = mixedToBig(v)
			case []int32:
				result, err = mixedToBig(v)
			case []int64:
				result, err = mixedToBig(v)
			case []uint:
				result, err = mixedToBig(v)
			case []uint8:
				result, err = mixedToBig(v)
			case []uint16:
				result, err = mixedToBig(v)
			case []uint32:
				result, err = mixedToBig(v)
			case []uint64:
				result, err = mixedToBig(v)
			case []*big.Int:
				result, err = mixedToBig(v)
			default:
				t.Fatalf("Unsupported input type: %T", v)
			}

			if err != nil {
				t.Errorf("mixedToBig() with %s failed: %v", test.name, err)

				return
			}

			if len(result) != len(test.expected) {
				t.Errorf("mixedToBig() returned wrong length: got %d, expected %d", len(result), len(test.expected))

				return
			}

			// Check values
			for i, val := range result {
				actual, _ := val.Float64()
				if math.Abs(actual-test.expected[i]) > test.tolerance {
					t.Errorf("mixedToBig() %s[%d] = %v, expected %v", test.name, i, actual, test.expected[i])
				}
			}
		})
	}
}

func TestMixedToBigFloatingPoint(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		expected  []float64
		tolerance float64
	}{
		{
			name:      "float32 types",
			input:     []float32{1.5, 2.5, 3.5},
			expected:  []float64{1.5, 2.5, 3.5},
			tolerance: 1e-6,
		},
		{
			name:      "float64 types",
			input:     []float64{1.123456789, 2.987654321, 3.141592653},
			expected:  []float64{1.123456789, 2.987654321, 3.141592653},
			tolerance: 1e-10,
		},
		{
			name:      "big.Float types",
			input:     []*big.Float{big.NewFloat(1.23), big.NewFloat(4.56), big.NewFloat(7.89)},
			expected:  []float64{1.23, 4.56, 7.89},
			tolerance: 1e-10,
		},
		{
			name:      "float32 precision edge cases",
			input:     []float32{1.175494e-38, 3.402823e+38},
			expected:  []float64{1.175494e-38, 3.402823e+38},
			tolerance: 1e+35, // Very large tolerance for extreme float32 max values due to precision conversion
		},
		{
			name:      "float64 precision edge cases",
			input:     []float64{2.2250738585072014e-308, 1.7976931348623157e+308},
			expected:  []float64{2.2250738585072014e-308, 1.7976931348623157e+308},
			tolerance: 1e-300, // Very small tolerance for extreme values
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result []*big.Float
			var err error

			switch v := test.input.(type) {
			case []float32:
				result, err = mixedToBig(v)
			case []float64:
				result, err = mixedToBig(v)
			case []*big.Float:
				result, err = mixedToBig(v)
			default:
				t.Fatalf("Unsupported input type: %T", v)
			}

			if err != nil {
				t.Errorf("mixedToBig() with %s failed: %v", test.name, err)

				return
			}

			if len(result) != len(test.expected) {
				t.Errorf("mixedToBig() returned wrong length: got %d, expected %d", len(result), len(test.expected))

				return
			}

			// Check values
			for i, val := range result {
				actual, _ := val.Float64()
				if math.Abs(actual-test.expected[i]) > test.tolerance {
					t.Errorf("mixedToBig() %s[%d] = %v, expected %v", test.name, i, actual, test.expected[i])
				}
			}
		})
	}
}

func TestMixedToBigEdgeCases(t *testing.T) {
	// Test empty slice
	t.Run("empty slice", func(t *testing.T) {
		emptyData := []int{}
		result, err := mixedToBig(emptyData)
		if err != nil {
			t.Errorf("mixedToBig() with empty slice failed: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("mixedToBig() returned wrong length for empty slice: got %d, expected 0", len(result))
		}
	})

	// Test nil pointer in *big.Float slice
	t.Run("nil big.Float pointer", func(t *testing.T) {
		bigFloatData := make([]*big.Float, 2)
		bigFloatData[0] = big.NewFloat(1.0)
		bigFloatData[1] = nil // This will cause a panic

		defer func() {
			if r := recover(); r != nil {
				t.Logf("mixedToBig() with nil *big.Float panicked as expected: %v", r)
			} else {
				t.Errorf("mixedToBig() with nil *big.Float should have panicked but didn't")
			}
		}()

		_, _ = mixedToBig(bigFloatData)
	})

	// Test nil pointer in *big.Int slice
	t.Run("nil big.Int pointer", func(t *testing.T) {
		bigIntData := make([]*big.Int, 2)
		bigIntData[0] = big.NewInt(1)
		bigIntData[1] = nil // This will cause a panic

		defer func() {
			if r := recover(); r != nil {
				t.Logf("mixedToBig() with nil *big.Int panicked as expected: %v", r)
			} else {
				t.Errorf("mixedToBig() with nil *big.Int should have panicked but didn't")
			}
		}()

		_, _ = mixedToBig(bigIntData)
	})
}

func TestBigNumericToBigFloat(t *testing.T) {
	t.Run("big.Float input", func(t *testing.T) {
		input := big.NewFloat(123.456)
		result := bigNumericToBigFloat(input)

		// Should be a copy, not the same pointer
		if result == input {
			t.Error("bigNumericToBigFloat() should return a copy, not the same pointer")
		}

		// Should have the same value
		if result.Cmp(input) != 0 {
			t.Errorf("bigNumericToBigFloat() = %v, want %v", result, input)
		}

		// Verify it's actually a copy by modifying the original
		input.SetFloat64(999.999)
		expected := big.NewFloat(123.456)
		if result.Cmp(expected) != 0 {
			t.Error("bigNumericToBigFloat() should return an independent copy")
		}
	})

	t.Run("big.Int input", func(t *testing.T) {
		input := big.NewInt(987654321)
		result := bigNumericToBigFloat(input)

		expected := big.NewFloat(987654321)
		if result.Cmp(expected) != 0 {
			t.Errorf("bigNumericToBigFloat() = %v, want %v", result, expected)
		}

		// Verify original big.Int is unchanged
		originalValue := big.NewInt(987654321)
		if input.Cmp(originalValue) != 0 {
			t.Error("bigNumericToBigFloat() should not modify the input big.Int")
		}
	})

	t.Run("negative big.Float", func(t *testing.T) {
		input := big.NewFloat(-42.5)
		result := bigNumericToBigFloat(input)

		if result.Cmp(input) != 0 {
			t.Errorf("bigNumericToBigFloat() = %v, want %v", result, input)
		}
	})

	t.Run("negative big.Int", func(t *testing.T) {
		input := big.NewInt(-123)
		result := bigNumericToBigFloat(input)

		expected := big.NewFloat(-123)
		if result.Cmp(expected) != 0 {
			t.Errorf("bigNumericToBigFloat() = %v, want %v", result, expected)
		}
	})

	t.Run("zero values", func(t *testing.T) {
		// Test zero big.Float
		zeroFloat := big.NewFloat(0)
		result := bigNumericToBigFloat(zeroFloat)
		if result.Sign() != 0 {
			t.Errorf("bigNumericToBigFloat() with zero big.Float = %v, want 0", result)
		}

		// Test zero big.Int
		zeroInt := big.NewInt(0)
		result = bigNumericToBigFloat(zeroInt)
		if result.Sign() != 0 {
			t.Errorf("bigNumericToBigFloat() with zero big.Int = %v, want 0", result)
		}
	})

	// And also test cases that should panic.
	t.Run("nil big.Float pointer - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("bigNumericToBigFloat() with nil *big.Float panicked as expected: %v", r)
			} else {
				t.Error("bigNumericToBigFloat() with nil *big.Float should have panicked but didn't")
			}
		}()

		var nilFloat *big.Float
		bigNumericToBigFloat(nilFloat)
	})

	t.Run("nil big.Int pointer - should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("bigNumericToBigFloat() with nil *big.Int panicked as expected: %v", r)
			} else {
				t.Error("bigNumericToBigFloat() with nil *big.Int should have panicked but didn't")
			}
		}()

		var nilInt *big.Int
		bigNumericToBigFloat(nilInt)
	})
}

func BenchmarkCorrelate(b *testing.B) {
	x := make([]float64, 1000)
	y := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		x[i] = float64(i)
		y[i] = float64(i*2 + 1)
	}

	b.ResetTimer()
	for b.Loop() {
		_, _ = Correlate(x, y, Pearson)
	}
}

func BenchmarkCorrelateInt(b *testing.B) {
	x := make([]int, 1000)
	y := make([]int, 1000)
	for i := range 1000 {
		x[i] = i
		y[i] = i*2 + 1
	}

	b.ResetTimer()
	for b.Loop() {
		_, _ = Correlate(x, y, Pearson)
	}
}

func BenchmarkCorrelateBigFloat(b *testing.B) {
	x := make([]*big.Float, 1000)
	y := make([]*big.Float, 1000)
	for i := range 1000 {
		x[i] = big.NewFloat(float64(i))
		y[i] = big.NewFloat(float64(i*2 + 1))
	}

	b.ResetTimer()
	for b.Loop() {
		_, _ = CorrelateBig(x, y, Pearson)
	}
}

func BenchmarkCorrelateBigInt(b *testing.B) {
	x := make([]*big.Int, 1000)
	y := make([]*big.Int, 1000)
	for i := range 1000 {
		x[i] = big.NewInt(int64(i))
		y[i] = big.NewInt(int64(i*2 + 1))
	}

	b.ResetTimer()
	for b.Loop() {
		_, _ = CorrelateBig(x, y, Pearson)
	}
}

func BenchmarkCorrelateBigFloatFromFloat64s(b *testing.B) {
	// Test different precision levels
	precisions := []uint{53, 64, 128, 256, 512, 1024, 2048}

	for _, prec := range precisions {
		b.Run(fmt.Sprintf("Precision_%d", prec), func(b *testing.B) {
			// Create test data with specified precision
			x := make([]*big.Float, 1000)
			y := make([]*big.Float, 1000)

			for i := range 1000 {
				x[i] = new(big.Float).SetPrec(prec).SetFloat64(float64(i))
				y[i] = new(big.Float).SetPrec(prec).SetFloat64(float64(i*2 + 1))
			}

			b.ResetTimer()
			for b.Loop() {
				_, _ = CorrelateBig(x, y, Pearson)
			}
		})
	}
}

func BenchmarkCorrelateBigFloatPrecisionLargeNumbers(b *testing.B) {
	// Test performance with large numbers that exceed float64 limits
	precisions := []uint{53, 64, 128, 256, 512, 1024}

	for _, prec := range precisions {
		b.Run(fmt.Sprintf("LargeNumbers_Precision_%d", prec), func(b *testing.B) {
			// Create test data with large numbers
			x := make([]*big.Float, 500)
			y := make([]*big.Float, 500)

			for i := range 500 {
				// Use numbers that exceed float64 limits
				xStr := fmt.Sprintf("%de400", i+1)
				yStr := fmt.Sprintf("%de400", (i+1)*2)

				x[i] = new(big.Float).SetPrec(prec)
				y[i] = new(big.Float).SetPrec(prec)

				x[i].SetString(xStr)
				y[i].SetString(yStr)
			}

			b.ResetTimer()
			for b.Loop() {
				_, _ = CorrelateBig(x, y, Pearson)
			}
		})
	}
}
