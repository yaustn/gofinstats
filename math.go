package gofinstats

import "math"

// StdDev returns the standard deviation of the input vals slice
func StdDev(vals []float64) float64 {
	mean := Avg(vals)

	var deviationSum float64
	for _, val := range vals {
		deviationSum += math.Pow(val-mean, 2)
	}

	variance := deviationSum / float64(len(vals))

	return math.Sqrt(variance)
}

// Avg returns the average of the input vals slice
func Avg(vals []float64) float64 {
	return Sum(vals) / float64(len(vals))
}

// Sum returns the sum of the input vals slice
func Sum(vals []float64) float64 {
	var sum float64
	for _, val := range vals {
		sum += val
	}

	return sum
}
