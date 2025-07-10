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
	"math/rand"
	"testing"
)

func BenchmarkKendallsTau100(b *testing.B) {
	x := make([]float64, 100)
	y := make([]float64, 100)
	rng := rand.New(rand.NewSource(getSeed()))
	for i := 0; i < 100; i++ {
		x[i] = rng.Float64() * 100
		y[i] = rng.Float64() * 100
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KendallsTau(x, y)
	}
}

func BenchmarkKendallsTau1000(b *testing.B) {
	x := make([]float64, 1000)
	y := make([]float64, 1000)
	rng := rand.New(rand.NewSource(getSeed()))
	for i := 0; i < 1000; i++ {
		x[i] = rng.Float64() * 100
		y[i] = rng.Float64() * 100
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KendallsTau(x, y)
	}
}
