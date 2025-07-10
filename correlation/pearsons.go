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
	"math"
	"math/big"
)

// Pearsons calculates Pearson's product-moment correlation coefficient
// between two datasets x and y of any numeric type.
//
// It returns a value between -1 and 1, where:
//   - 1 indicates a perfect positive linear relationship
//   - 0 indicates no linear relationship
//   - -1 indicates a perfect negative linear relationship
//
// An error is returned if the slices have different lengths or are empty.
func Pearsons[T Numeric](x, y []T) (float64, error) {
	return pearsonsSinglePass(x, y)
}

// correlatePearsonTwoPass calculates Pearson's correlation using the classic two-pass algorithm.
func pearsonsTwoPass[T Numeric](x, y []T) (float64, error) {
	if len(x) == 0 || len(y) == 0 {
		return 0, errors.New("input slices cannot be empty")
	}

	if len(x) != len(y) {
		return 0, errors.New("input slices must have the same length")
	}

	n := len(x)
	if n == 1 {
		return 0, errors.New("correlation requires at least 2 data points")
	}

	// Calculate means
	var sumX, sumY float64
	for i := 0; i < n; i++ {
		sumX += float64(x[i])
		sumY += float64(y[i])
	}
	meanX := sumX / float64(n)
	meanY := sumY / float64(n)

	// Calculate numerator and denominators
	var numerator, sumXX, sumYY float64
	for i := 0; i < n; i++ {
		dx := float64(x[i]) - meanX
		dy := float64(y[i]) - meanY
		numerator += dx * dy
		sumXX += dx * dx
		sumYY += dy * dy
	}

	// Check for zero variance
	if sumXX == 0 || sumYY == 0 {
		return 0, errors.New("correlation undefined: one or both variables have zero variance")
	}

	denominator := math.Sqrt(sumXX * sumYY)

	return numerator / denominator, nil
}

// correlatePearsonSinglePass calculates Pearson's correlation using a single-pass algorithm.
// This is more numerically efficient as it computes the correlation in one iteration
// through the data without requiring separate passes for means and correlation.
//
// The incoming values in X must be ordered.
func pearsonsSinglePass[T Numeric](x, y []T) (float64, error) {
	if len(x) == 0 || len(y) == 0 {
		return 0, errors.New("input slices cannot be empty")
	}

	if len(x) != len(y) {
		return 0, errors.New("input slices must have the same length")
	}

	n := len(x)
	if n == 1 {
		return 0, errors.New("correlation requires at least 2 data points")
	}

	// Single-pass algorithm using Welford's online algorithm approach
	var sumX, sumY, sumXY, sumXX, sumYY float64

	for i := range n {
		fx := float64(x[i])
		fy := float64(y[i])

		sumX += fx
		sumY += fy
		sumXY += fx * fy
		sumXX += fx * fx
		sumYY += fy * fy
	}

	// We need to check if any of these blew past math.MaxFloat64
	if math.IsInf(sumX, 0) || math.IsInf(sumY, 0) || math.IsInf(sumXY, 0) || math.IsInf(sumXX, 0) || math.IsInf(sumYY, 0) {
		bigX, err := mixedToBig(x)
		if err != nil {
			return 0, errors.New("Pearson's calculation needs to convert to big, but conversion failed: " + err.Error())
		}
		bigY, err := mixedToBig(y)
		if err != nil {
			return 0, errors.New("Pearson's calculation needs to convert to big, but conversion failed: " + err.Error())
		}

		return PearsonsBig(bigX, bigY)
	}

	nf := float64(n)

	// Calculate numerator: sum(xy) - n*mean(x)*mean(y)
	numerator := sumXY - (sumX*sumY)/nf

	// Calculate denominators: sqrt((sum(x²) - n*mean(x)²) * (sum(y²) - n*mean(y)²))
	varX := sumXX - (sumX*sumX)/nf
	varY := sumYY - (sumY*sumY)/nf

	// Check for zero variance
	if varX <= 0 || varY <= 0 {
		return 0, errors.New("correlation undefined: one or both variables have zero variance")
	}

	denominator := math.Sqrt(varX * varY)

	return numerator / denominator, nil
}

// PearsonsBig calculates Pearson's product-moment correlation coefficient
// between two datasets x and y of big number types (*big.Float or *big.Int).
// Uses a single-pass algorithm for efficiency.
//
// It returns a value between -1 and 1, where:
//   - 1 indicates a perfect positive linear relationship
//   - 0 indicates no linear relationship
//   - -1 indicates a perfect negative linear relationship
//
// An error is returned if the slices have different lengths or are empty.
func PearsonsBig[T BigNumeric](x, y []T) (float64, error) {
	if len(x) == 0 || len(y) == 0 {
		return 0, errors.New("input slices cannot be empty")
	}

	if len(x) != len(y) {
		return 0, errors.New("input slices must have the same length")
	}

	n := len(x)
	if n == 1 {
		return 0, errors.New("correlation requires at least 2 data points")
	}

	// Convert all inputs to *big.Float for consistent arithmetic
	xVals := make([]*big.Float, n)
	yVals := make([]*big.Float, n)

	for i := 0; i < n; i++ {
		switch v := any(x[i]).(type) {
		case *big.Float:
			xVals[i] = new(big.Float).Copy(v)
		case *big.Int:
			xVals[i] = new(big.Float).SetInt(v)
		}

		switch v := any(y[i]).(type) {
		case *big.Float:
			yVals[i] = new(big.Float).Copy(v)
		case *big.Int:
			yVals[i] = new(big.Float).SetInt(v)
		}
	}

	// Single-pass algorithm using big.Float arithmetic
	sumX := new(big.Float)
	sumY := new(big.Float)
	sumXY := new(big.Float)
	sumXX := new(big.Float)
	sumYY := new(big.Float)

	temp := new(big.Float)

	for i := range n {
		fx := xVals[i]
		fy := yVals[i]

		// sumX += fx
		sumX.Add(sumX, fx)
		// sumY += fy
		sumY.Add(sumY, fy)
		// sumXY += fx * fy
		temp.Mul(fx, fy)
		sumXY.Add(sumXY, temp)
		// sumXX += fx * fx
		temp.Mul(fx, fx)
		sumXX.Add(sumXX, temp)
		// sumYY += fy * fy
		temp.Mul(fy, fy)
		sumYY.Add(sumYY, temp)
	}

	nf := new(big.Float).SetInt64(int64(n))

	// Calculate numerator: sumXY - (sumX * sumY) / n
	numerator := new(big.Float)
	temp.Mul(sumX, sumY)
	temp.Quo(temp, nf)

	// Check for infinity cases that would cause "subtraction of infinities with equal signs"
	if sumXY.IsInf() && temp.IsInf() && sumXY.Signbit() == temp.Signbit() {
		// Both are infinite with same sign - correlation is undefined
		return 0, errors.New("correlation undefined: infinite values with same sign detected")
	}
	numerator.Sub(sumXY, temp)

	// Calculate varX: sumXX - (sumX * sumX) / n
	varX := new(big.Float)
	temp.Mul(sumX, sumX)
	temp.Quo(temp, nf)

	// Check for infinity cases in variance calculation
	if sumXX.IsInf() && temp.IsInf() && sumXX.Signbit() == temp.Signbit() {
		return 0, errors.New("correlation undefined: infinite variance detected in X")
	}
	varX.Sub(sumXX, temp)

	// Calculate varY: sumYY - (sumY * sumY) / n
	varY := new(big.Float)
	temp.Mul(sumY, sumY)
	temp.Quo(temp, nf)

	// Check for infinity cases in variance calculation
	if sumYY.IsInf() && temp.IsInf() && sumYY.Signbit() == temp.Signbit() {
		return 0, errors.New("correlation undefined: infinite variance detected in Y")
	}
	varY.Sub(sumYY, temp)

	// Check for zero variance
	zero := new(big.Float)
	if varX.Cmp(zero) <= 0 || varY.Cmp(zero) <= 0 {
		return 0, errors.New("correlation undefined: one or both variables have zero variance")
	}

	// Calculate denominator: sqrt(varX * varY)
	denominator := new(big.Float)
	denominator.Mul(varX, varY)
	denominator.Sqrt(denominator)

	// Calculate correlation: numerator / denominator
	correlation := new(big.Float)
	correlation.Quo(numerator, denominator)

	// Convert to float64 for return
	result, _ := correlation.Float64()

	return result, nil
}

// PearsonsMixed calculates Pearson's product-moment correlation coefficient
// between two datasets x and y with mixed type inputs.
// It converts the inputs using mixedToBig and then calls PearsonsBig.
func PearsonsMixed[T MixedNumeric](x, y []T) (float64, error) {
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

	return PearsonsBig(xVals, yVals)
}
