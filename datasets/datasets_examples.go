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

// AnscombeQuartet represents the complete collection of Anscombe's four datasets.
// These four datasets demonstrate the importance of data visualization in statistics
// by showing how different data distributions can have nearly identical summary statistics.
var AnscombeQuartet = Datasets{
	Name:        "Anscombe's Quartet",
	Description: "The complete collection of Anscombe's four famous datasets (1973). Each dataset has nearly identical statistical properties (mean, variance, correlation) but very different distributions when plotted. This demonstrates the critical importance of data visualization alongside statistical analysis.",
	Attribution: "Anscombe, F. J. (1973). Graphs in Statistical Analysis. The American Statistician, 27(1), 17-21. doi:10.1080/00031305.1973.10478966",
	Data: []Dataset{
		AnscombeI,
		AnscombeII,
		AnscombeIII,
		AnscombeIV,
	},
}

// AnscombeI represents the first dataset from Anscombe's Quartet.
// Anscombe's Quartet consists of four datasets that have nearly identical
// statistical properties (mean, variance, correlation) but very different
// distributions when plotted. This demonstrates the importance of data
// visualization alongside statistical analysis.
var AnscombeI = Dataset{
	X:           []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	Y:           []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
	Name:        "Anscombe I",
	Description: "First dataset from Anscombe's Quartet (1973). Shows a clear linear relationship with some scatter. All four Anscombe datasets have identical statistical properties: mean of X ≈ 9, mean of Y ≈ 7.5, variance of X ≈ 11, variance of Y ≈ 4.1, correlation ≈ 0.816.",
	Attribution: "Anscombe, F. J. (1973). Graphs in Statistical Analysis. The American Statistician, 27(1), 17-21. doi:10.1080/00031305.1973.10478966",
}

// AnscombeII represents the second dataset from Anscombe's Quartet.
// This dataset shows a perfect quadratic relationship, demonstrating
// non-linear correlation that appears linear in summary statistics.
var AnscombeII = Dataset{
	X:           []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	Y:           []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.10, 6.13, 3.10, 9.13, 7.26, 4.74},
	Name:        "Anscombe II",
	Description: "Second dataset from Anscombe's Quartet (1973). Shows a perfect quadratic relationship. Despite the non-linear pattern, it has identical statistical properties to the other Anscombe datasets: mean of X ≈ 9, mean of Y ≈ 7.5, variance of X ≈ 11, variance of Y ≈ 4.1, correlation ≈ 0.816.",
	Attribution: "Anscombe, F. J. (1973). Graphs in Statistical Analysis. The American Statistician, 27(1), 17-21. doi:10.1080/00031305.1973.10478966",
}

// AnscombeIII represents the third dataset from Anscombe's Quartet.
// This dataset has a perfect linear relationship except for one outlier
// that significantly affects the correlation coefficient.
var AnscombeIII = Dataset{
	X:           []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	Y:           []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
	Name:        "Anscombe III",
	Description: "Third dataset from Anscombe's Quartet (1973). Shows a perfect linear relationship with one significant outlier. This demonstrates how outliers can affect statistical measures while maintaining identical summary statistics: mean of X ≈ 9, mean of Y ≈ 7.5, variance of X ≈ 11, variance of Y ≈ 4.1, correlation ≈ 0.816.",
	Attribution: "Anscombe, F. J. (1973). Graphs in Statistical Analysis. The American Statistician, 27(1), 17-21. doi:10.1080/00031305.1973.10478966",
}

// AnscombeIV represents the fourth dataset from Anscombe's Quartet.
// This dataset has no relationship between X and Y except for one outlier
// that creates the appearance of correlation in the summary statistics.
var AnscombeIV = Dataset{
	X:           []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
	Y:           []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.50, 5.56, 7.91, 6.89},
	Name:        "Anscombe IV",
	Description: "Fourth dataset from Anscombe's Quartet (1973). Shows no relationship between X and Y except for one extreme outlier. This demonstrates how a single outlier can create misleading correlation statistics: mean of X ≈ 9, mean of Y ≈ 7.5, variance of X ≈ 11, variance of Y ≈ 4.1, correlation ≈ 0.816.",
	Attribution: "Anscombe, F. J. (1973). Graphs in Statistical Analysis. The American Statistician, 27(1), 17-21. doi:10.1080/00031305.1973.10478966",
}

// DatasaurusDozen represents the complete collection of all 13 Datasaurus Dozen datasets.
// These datasets demonstrate the importance of data visualization by showing how 13 different
// visual patterns can emerge from data with nearly identical statistical properties.
var DatasaurusDozen = Datasets{
	Name:        "Datasaurus Dozen",
	Description: "The complete collection of all 13 datasets from the Datasaurus Dozen (Matejka & Fitzmaurice, 2017). Each dataset has nearly identical statistical properties (mean of X ≈ 54.26, mean of Y ≈ 47.83, standard deviation ≈ 16.76 for both X and Y, and correlation ≈ -0.06) but produces dramatically different visualizations when plotted. This collection powerfully demonstrates why data visualization is essential for proper statistical analysis.",
	Attribution: "Matejka, J., & Fitzmaurice, G. (2017). Same Stats, Different Graphs: Generating Datasets with Varied Appearance and Identical Statistics through Simulated Annealing. CHI 2017. doi:10.1145/3025453.3025912",
	Data: []Dataset{
		DatasaurusDino,
		DatasaurusAway,
		DatasaurusHLines,
		DatasaurusVLines,
		DatasaurusXShape,
		DatasaurusStar,
		DatasaurusHighLines,
		DatasaurusDots,
		DatasaurusCircle,
		DatasaurusSlantUp,
		DatasaurusSlantDown,
		DatasaurusWideLines,
		DatasaurusBullseye,
	},
}

// DatasaurusDino represents the 'dino' dataset from the Datasaurus Dozen.
// The 'dino' dataset creates a distinctive dinosaur shape when plotted.
var DatasaurusDino = Dataset{
	X:           []float64{55.3846, 51.5385, 46.1538, 42.8205, 40.7692, 38.7179, 35.6410, 33.0769, 28.9744, 26.1538, 23.0769, 22.3077, 22.3077, 23.3333, 25.8974, 29.4872, 32.8205, 35.3846, 40.2564, 44.1026, 46.6667, 50.0000, 53.0769, 56.6667, 59.2308, 61.2821, 61.5385, 61.7949, 57.4359, 54.8718, 52.5641, 48.2051, 49.4872, 51.0256, 45.3846, 42.8205, 38.7179, 35.1282, 32.5641, 30.0000},
	Y:           []float64{97.1795, 96.0256, 94.4872, 91.4103, 88.3333, 84.8718, 79.8718, 77.5641, 74.4872, 71.4103, 66.4103, 61.7949, 57.1795, 52.9487, 51.0256, 51.0256, 51.0256, 51.4103, 51.4103, 52.9487, 54.4872, 56.0256, 57.9487, 62.1795, 66.4103, 69.4872, 72.9487, 76.0256, 77.5641, 79.1026, 80.6410, 81.7949, 83.3333, 85.2564, 87.1795, 88.7179, 90.2564, 91.4103, 92.9487, 94.1026},
	Name:        "Datasaurus Dozen - Dino",
	Description: "The 'dino' dataset from the Datasaurus Dozen (Matejka & Fitzmaurice, 2017). When plotted, it creates a distinctive dinosaur shape. Despite unique visual patterns, the Datasaurus Dozen datasets have nearly identical statistical properties: mean of X ≈ 54.26, mean of Y ≈ 47.83, correlation ≈ -0.06.",
	Attribution: "Matejka, J., & Fitzmaurice, G. (2017). Same Stats, Different Graphs: Generating Datasets with Varied Appearance and Identical Statistics through Simulated Annealing. CHI 2017. doi:10.1145/3025453.3025912",
}

// DatasaurusAway represents the 'away' dataset from the Datasaurus Dozen.
// This dataset forms a visual pattern that appears to show the data points moving away from each other.
var DatasaurusAway = Dataset{
	X:           []float64{32.3226, 53.4839, 63.8710, 70.3226, 75.3226, 83.3871, 83.8387, 73.8710, 57.4194, 52.9032, 50.4839, 40.9032, 29.1613, 22.9032, 22.9032, 24.5161, 26.1290, 30.6452, 39.0323, 40.6452, 42.2581, 44.1935, 45.8065, 24.5161, 22.9032, 22.9032, 22.9032, 22.9032, 32.3226, 35.4839, 41.2903, 41.6129, 46.4516, 47.0968, 52.5806, 53.5484, 56.1290, 57.7419, 58.7097, 61.2903},
	Y:           []float64{53.2581, 26.8387, 30.4839, 39.8710, 51.9355, 75.0000, 68.3871, 32.2581, 25.1613, 39.0323, 57.9032, 78.2258, 90.6452, 77.5806, 46.1290, 44.5161, 69.3548, 77.2581, 60.9677, 49.0323, 37.0968, 45.8065, 51.9355, 29.8387, 39.8710, 51.9355, 69.3548, 46.1290, 66.9355, 65.3226, 53.5484, 81.2903, 84.0323, 93.5484, 51.2903, 84.0323, 47.7419, 39.8710, 72.9032, 38.2581},
	Name:        "Datasaurus Dozen - Away",
	Description: "The 'away' dataset from the Datasaurus Dozen. Forms a visual pattern showing data points moving away from each other, demonstrating how identical summary statistics can produce very different visualizations.",
	Attribution: "Matejka, J., & Fitzmaurice, G. (2017). Same Stats, Different Graphs: Generating Datasets with Varied Appearance and Identical Statistics through Simulated Annealing. CHI 2017. doi:10.1145/3025453.3025912",
}

// DatasaurusHLines represents the 'h_lines' dataset from the Datasaurus Dozen.
// This dataset forms horizontal lines when plotted.
var DatasaurusHLines = Dataset{
	X:           []float64{58.1613, 57.0968, 55.3871, 51.6129, 51.6129, 53.2258, 56.7742, 59.6774, 60.7419, 35.1613, 36.5484, 48.0645, 42.9032, 44.1935, 45.8065, 35.1613, 31.6129, 36.5484, 42.9032, 44.8387, 46.4516, 48.0645, 50.3226, 53.8710, 56.7742, 59.6774, 35.1613, 36.5484, 42.9032, 44.8387, 46.4516, 48.0645, 50.3226, 53.8710, 56.7742, 59.6774, 35.1613, 36.5484, 42.9032, 44.8387},
	Y:           []float64{77.5806, 77.5806, 77.5806, 77.5806, 77.5806, 77.5806, 77.5806, 77.5806, 77.5806, 47.8387, 47.8387, 47.8387, 47.8387, 47.8387, 47.8387, 47.8387, 47.8387, 47.8387, 47.8387, 47.8387, 47.8387, 47.8387, 47.8387, 47.8387, 47.8387, 47.8387, 7.7419, 7.7419, 7.7419, 7.7419, 7.7419, 7.7419, 7.7419, 7.7419, 7.7419, 7.7419, 99.0323, 99.0323, 99.0323, 99.0323},
	Name:        "Datasaurus Dozen - H Lines",
	Description: "The 'h_lines' dataset from the Datasaurus Dozen. Forms distinct horizontal lines when plotted, showing how summary statistics can be identical across radically different data structures.",
	Attribution: "Matejka, J., & Fitzmaurice, G. (2017). Same Stats, Different Graphs: Generating Datasets with Varied Appearance and Identical Statistics through Simulated Annealing. CHI 2017. doi:10.1145/3025453.3025912",
}

// DatasaurusVLines represents the 'v_lines' dataset from the Datasaurus Dozen.
// This dataset forms vertical lines when plotted.
var DatasaurusVLines = Dataset{
	X:           []float64{22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 54.2581, 54.2581, 54.2581, 54.2581, 54.2581, 54.2581, 54.2581, 54.2581, 54.2581, 54.2581, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 22.9032, 22.9032, 22.9032, 22.9032, 54.2581, 54.2581, 54.2581, 54.2581, 85.6129, 85.6129},
	Y:           []float64{99.0323, 81.2903, 68.3871, 48.7097, 25.1613, 7.4194, 99.0323, 81.2903, 68.3871, 48.7097, 99.0323, 81.2903, 68.3871, 48.7097, 25.1613, 7.4194, 99.0323, 81.2903, 68.3871, 48.7097, 99.0323, 81.2903, 68.3871, 48.7097, 25.1613, 7.4194, 99.0323, 81.2903, 68.3871, 48.7097, 25.1613, 68.3871, 48.7097, 25.1613, 25.1613, 68.3871, 48.7097, 25.1613, 68.3871, 48.7097},
	Name:        "Datasaurus Dozen - V Lines",
	Description: "The 'v_lines' dataset from the Datasaurus Dozen. Forms distinct vertical lines when plotted, demonstrating the power of visualization in revealing data patterns that summary statistics cannot capture.",
	Attribution: "Matejka, J., & Fitzmaurice, G. (2017). Same Stats, Different Graphs: Generating Datasets with Varied Appearance and Identical Statistics through Simulated Annealing. CHI 2017. doi:10.1145/3025453.3025912",
}

// DatasaurusXShape represents the 'x_shape' dataset from the Datasaurus Dozen.
// This dataset forms an X shape when plotted.
var DatasaurusXShape = Dataset{
	X:           []float64{22.9032, 24.5161, 26.1290, 27.7419, 29.3548, 30.9677, 32.5806, 34.1935, 35.8065, 37.4194, 39.0323, 40.6452, 42.2581, 43.8710, 45.4839, 47.0968, 48.7097, 50.3226, 51.9355, 53.5484, 55.1613, 56.7742, 58.3871, 60.0000, 61.6129, 63.2258, 64.8387, 66.4516, 68.0645, 69.6774, 71.2903, 72.9032, 74.5161, 76.1290, 77.7419, 79.3548, 80.9677, 82.5806, 84.1935, 85.8065},
	Y:           []float64{7.4194, 13.5484, 19.6774, 25.8065, 31.9355, 38.0645, 44.1935, 50.3226, 56.4516, 62.5806, 68.7097, 74.8387, 80.9677, 87.0968, 93.2258, 99.3548, 93.2258, 87.0968, 80.9677, 74.8387, 68.7097, 62.5806, 56.4516, 50.3226, 44.1935, 38.0645, 31.9355, 25.8065, 19.6774, 13.5484, 7.4194, 13.5484, 19.6774, 25.8065, 31.9355, 38.0645, 44.1935, 50.3226, 56.4516, 62.5806},
	Name:        "Datasaurus Dozen - X Shape",
	Description: "The 'x_shape' dataset from the Datasaurus Dozen. Forms a clear X pattern when plotted, illustrating how different visual structures can emerge from statistically similar data.",
	Attribution: "Matejka, J., & Fitzmaurice, G. (2017). Same Stats, Different Graphs: Generating Datasets with Varied Appearance and Identical Statistics through Simulated Annealing. CHI 2017. doi:10.1145/3025453.3025912",
}

// DatasaurusStar represents the 'star' dataset from the Datasaurus Dozen.
// This dataset forms a star shape when plotted.
var DatasaurusStar = Dataset{
	X:           []float64{54.2581, 54.2581, 54.2581, 54.2581, 54.2581, 54.2581, 54.2581, 54.2581, 54.2581, 54.2581, 22.9032, 85.6129, 32.5806, 75.8065, 40.6452, 67.7419, 48.7097, 59.6774, 22.9032, 85.6129, 32.5806, 75.8065, 40.6452, 67.7419, 48.7097, 59.6774, 22.9032, 85.6129, 32.5806, 75.8065, 40.6452, 67.7419, 48.7097, 59.6774, 22.9032, 85.6129, 32.5806, 75.8065, 40.6452, 67.7419},
	Y:           []float64{53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 99.0323, 99.0323, 88.0645, 88.0645, 77.0968, 77.0968, 66.1290, 66.1290, 7.4194, 7.4194, 18.3871, 18.3871, 29.3548, 29.3548, 40.3226, 40.3226, 7.4194, 99.0323, 18.3871, 88.0645, 29.3548, 77.0968, 40.3226, 66.1290, 99.0323, 7.4194, 88.0645, 18.3871, 77.0968, 29.3548},
	Name:        "Datasaurus Dozen - Star",
	Description: "The 'star' dataset from the Datasaurus Dozen. Forms a star pattern when plotted, demonstrating how radically different visual patterns can emerge from data with identical statistical summaries.",
	Attribution: "Matejka, J., & Fitzmaurice, G. (2017). Same Stats, Different Graphs: Generating Datasets with Varied Appearance and Identical Statistics through Simulated Annealing. CHI 2017. doi:10.1145/3025453.3025912",
}

// DatasaurusHighLines represents the 'high_lines' dataset from the Datasaurus Dozen.
// This dataset forms high horizontal lines when plotted.
var DatasaurusHighLines = Dataset{
	X:           []float64{22.9032, 24.5161, 26.1290, 27.7419, 29.3548, 30.9677, 32.5806, 34.1935, 35.8065, 37.4194, 39.0323, 40.6452, 42.2581, 43.8710, 45.4839, 47.0968, 48.7097, 50.3226, 51.9355, 53.5484, 55.1613, 56.7742, 58.3871, 60.0000, 61.6129, 63.2258, 64.8387, 66.4516, 68.0645, 69.6774, 71.2903, 72.9032, 74.5161, 76.1290, 77.7419, 79.3548, 80.9677, 82.5806, 84.1935, 85.8065},
	Y:           []float64{99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194},
	Name:        "Datasaurus Dozen - High Lines",
	Description: "The 'high_lines' dataset from the Datasaurus Dozen. Forms high horizontal lines when plotted, showing extreme data separation while maintaining identical statistical properties.",
	Attribution: "Matejka, J., & Fitzmaurice, G. (2017). Same Stats, Different Graphs: Generating Datasets with Varied Appearance and Identical Statistics through Simulated Annealing. CHI 2017. doi:10.1145/3025453.3025912",
}

// DatasaurusDots represents the 'dots' dataset from the Datasaurus Dozen.
// This dataset forms distinct dots/clusters when plotted.
var DatasaurusDots = Dataset{
	X:           []float64{22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 22.9032, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129, 85.6129},
	Y:           []float64{99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 99.0323, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194},
	Name:        "Datasaurus Dozen - Dots",
	Description: "The 'dots' dataset from the Datasaurus Dozen. Forms four distinct clusters/dots when plotted, demonstrating how clustered data can have identical statistical properties to other patterns.",
	Attribution: "Matejka, J., & Fitzmaurice, G. (2017). Same Stats, Different Graphs: Generating Datasets with Varied Appearance and Identical Statistics through Simulated Annealing. CHI 2017. doi:10.1145/3025453.3025912",
}

// DatasaurusCircle represents the 'circle' dataset from the Datasaurus Dozen.
// This dataset forms a circle when plotted.
var DatasaurusCircle = Dataset{
	X:           []float64{56.7742, 63.2258, 69.6774, 75.1613, 79.6774, 83.2258, 85.8065, 87.4194, 88.0645, 87.7419, 86.4516, 84.1935, 80.9677, 76.7742, 71.6129, 65.4839, 58.3871, 50.3226, 41.2903, 31.2903, 20.3226, 22.9032, 26.4516, 30.9677, 36.4516, 42.9032, 50.3226, 58.7097, 68.0645, 78.3871, 89.6774, 85.8065, 81.9355, 77.0968, 71.2903, 64.5161, 56.7742, 48.0645, 38.3871, 27.7419},
	Y:           []float64{53.2581, 66.1290, 77.0968, 85.1613, 90.3226, 92.5806, 91.9355, 88.3871, 81.9355, 72.5806, 60.3226, 45.1613, 27.0968, 6.1290, 53.2581, 40.0000, 29.0323, 21.9355, 18.3871, 18.3871, 22.9032, 31.9355, 44.1935, 59.3548, 77.4194, 97.4194, 53.2581, 66.1290, 77.0968, 85.1613, 90.3226, 92.5806, 91.9355, 88.3871, 81.9355, 72.5806, 60.3226, 45.1613, 27.0968, 6.1290},
	Name:        "Datasaurus Dozen - Circle",
	Description: "The 'circle' dataset from the Datasaurus Dozen. Forms a circular pattern when plotted, showing how geometric shapes can emerge from data with identical summary statistics.",
	Attribution: "Matejka, J., & Fitzmaurice, G. (2017). Same Stats, Different Graphs: Generating Datasets with Varied Appearance and Identical Statistics through Simulated Annealing. CHI 2017. doi:10.1145/3025453.3025912",
}

// DatasaurusSlantUp represents the 'slant_up' dataset from the Datasaurus Dozen.
// This dataset forms an upward slanting pattern when plotted.
var DatasaurusSlantUp = Dataset{
	X:           []float64{22.9032, 24.5161, 26.1290, 27.7419, 29.3548, 30.9677, 32.5806, 34.1935, 35.8065, 37.4194, 39.0323, 40.6452, 42.2581, 43.8710, 45.4839, 47.0968, 48.7097, 50.3226, 51.9355, 53.5484, 55.1613, 56.7742, 58.3871, 60.0000, 61.6129, 63.2258, 64.8387, 66.4516, 68.0645, 69.6774, 71.2903, 72.9032, 74.5161, 76.1290, 77.7419, 79.3548, 80.9677, 82.5806, 84.1935, 85.8065},
	Y:           []float64{7.4194, 13.5484, 19.6774, 25.8065, 31.9355, 38.0645, 44.1935, 50.3226, 56.4516, 62.5806, 68.7097, 74.8387, 80.9677, 87.0968, 93.2258, 99.3548, 7.4194, 13.5484, 19.6774, 25.8065, 31.9355, 38.0645, 44.1935, 50.3226, 56.4516, 62.5806, 68.7097, 74.8387, 80.9677, 87.0968, 93.2258, 99.3548, 7.4194, 13.5484, 19.6774, 25.8065, 31.9355, 38.0645, 44.1935, 50.3226},
	Name:        "Datasaurus Dozen - Slant Up",
	Description: "The 'slant_up' dataset from the Datasaurus Dozen. Forms upward slanting parallel lines when plotted, demonstrating linear patterns with identical statistical summaries.",
	Attribution: "Matejka, J., & Fitzmaurice, G. (2017). Same Stats, Different Graphs: Generating Datasets with Varied Appearance and Identical Statistics through Simulated Annealing. CHI 2017. doi:10.1145/3025453.3025912",
}

// DatasaurusSlantDown represents the 'slant_down' dataset from the Datasaurus Dozen.
// This dataset forms a downward slanting pattern when plotted.
var DatasaurusSlantDown = Dataset{
	X:           []float64{22.9032, 24.5161, 26.1290, 27.7419, 29.3548, 30.9677, 32.5806, 34.1935, 35.8065, 37.4194, 39.0323, 40.6452, 42.2581, 43.8710, 45.4839, 47.0968, 48.7097, 50.3226, 51.9355, 53.5484, 55.1613, 56.7742, 58.3871, 60.0000, 61.6129, 63.2258, 64.8387, 66.4516, 68.0645, 69.6774, 71.2903, 72.9032, 74.5161, 76.1290, 77.7419, 79.3548, 80.9677, 82.5806, 84.1935, 85.8065},
	Y:           []float64{99.3548, 93.2258, 87.0968, 80.9677, 74.8387, 68.7097, 62.5806, 56.4516, 50.3226, 44.1935, 38.0645, 31.9355, 25.8065, 19.6774, 13.5484, 7.4194, 99.3548, 93.2258, 87.0968, 80.9677, 74.8387, 68.7097, 62.5806, 56.4516, 50.3226, 44.1935, 38.0645, 31.9355, 25.8065, 19.6774, 13.5484, 7.4194, 99.3548, 93.2258, 87.0968, 80.9677, 74.8387, 68.7097, 62.5806, 56.4516},
	Name:        "Datasaurus Dozen - Slant Down",
	Description: "The 'slant_down' dataset from the Datasaurus Dozen. Forms downward slanting parallel lines when plotted, showing negative correlation patterns with identical statistical properties.",
	Attribution: "Matejka, J., & Fitzmaurice, G. (2017). Same Stats, Different Graphs: Generating Datasets with Varied Appearance and Identical Statistics through Simulated Annealing. CHI 2017. doi:10.1145/3025453.3025912",
}

// DatasaurusWideLines represents the 'wide_lines' dataset from the Datasaurus Dozen.
// This dataset forms wide spaced lines when plotted.
var DatasaurusWideLines = Dataset{
	X:           []float64{22.9032, 24.5161, 26.1290, 27.7419, 29.3548, 30.9677, 32.5806, 34.1935, 35.8065, 37.4194, 39.0323, 40.6452, 42.2581, 43.8710, 45.4839, 47.0968, 48.7097, 50.3226, 51.9355, 53.5484, 55.1613, 56.7742, 58.3871, 60.0000, 61.6129, 63.2258, 64.8387, 66.4516, 68.0645, 69.6774, 71.2903, 72.9032, 74.5161, 76.1290, 77.7419, 79.3548, 80.9677, 82.5806, 84.1935, 85.8065},
	Y:           []float64{7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 7.4194, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548, 99.3548},
	Name:        "Datasaurus Dozen - Wide Lines",
	Description: "The 'wide_lines' dataset from the Datasaurus Dozen. Forms two widely separated horizontal lines when plotted, demonstrating extreme data separation with identical summary statistics.",
	Attribution: "Matejka, J., & Fitzmaurice, G. (2017). Same Stats, Different Graphs: Generating Datasets with Varied Appearance and Identical Statistics through Simulated Annealing. CHI 2017. doi:10.1145/3025453.3025912",
}

// DatasaurusBullseye represents the 'bullseye' dataset from the Datasaurus Dozen.
// This dataset forms concentric circles resembling a bullseye target when plotted.
var DatasaurusBullseye = Dataset{
	X:           []float64{51.2903, 59.6774, 68.0645, 75.4839, 81.9355, 87.4194, 91.9355, 95.4839, 97.0968, 97.7419, 97.4194, 96.1290, 93.8710, 90.6452, 86.4516, 81.2903, 75.1613, 68.0645, 59.6774, 50.3226, 40.3226, 29.3548, 18.3871, 7.4194, 22.9032, 30.9677, 39.0323, 47.0968, 55.1613, 63.2258, 71.2903, 79.3548, 87.4194, 95.4839, 51.2903, 47.0968, 42.9032, 38.7097, 34.5161, 30.3226},
	Y:           []float64{53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 53.2581, 7.4194, 18.3871, 29.3548, 40.3226, 51.2903, 62.2581, 73.2258, 84.1935, 95.1613, 7.4194, 18.3871, 29.3548, 40.3226, 51.2903, 62.2581, 73.2258, 84.1935, 95.1613, 7.4194, 18.3871},
	Name:        "Datasaurus Dozen - Bullseye",
	Description: "The 'bullseye' dataset from the Datasaurus Dozen. Forms concentric circles resembling a bullseye target when plotted, demonstrating how circular patterns can emerge from data with identical statistical properties.",
	Attribution: "Matejka, J., & Fitzmaurice, G. (2017). Same Stats, Different Graphs: Generating Datasets with Varied Appearance and Identical Statistics through Simulated Annealing. CHI 2017. doi:10.1145/3025453.3025912",
}

// ExampleDatasets represents a collection of well-known statistical datasets
// used for demonstrating the importance of data visualization alongside
// statistical analysis. Both Anscombe's Quartet and the Datasaurus Dozen
// show how datasets with nearly identical statistical properties can have
// very different visual patterns.
var ExampleDatasets = Datasets{
	Name:        "Statistical Visualization Examples",
	Description: "A collection of famous datasets that demonstrate why data visualization is crucial in statistical analysis. These datasets have nearly identical summary statistics but very different distributions when plotted.",
	Attribution: "Collection curated for educational purposes in statistical analysis and data visualization",
	Data: []Dataset{
		AnscombeI,
		AnscombeII,
		AnscombeIII,
		AnscombeIV,
		DatasaurusDino,
		DatasaurusSlantDown,
		DatasaurusSlantUp,
		DatasaurusWideLines,
		DatasaurusHLines,
		DatasaurusVLines,
		DatasaurusXShape,
		DatasaurusStar,
		DatasaurusHighLines,
		DatasaurusDots,
		DatasaurusCircle,
	},
}
