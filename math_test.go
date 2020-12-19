package gofinstats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStdDev(t *testing.T) {
	testCases := map[string]struct {
		vals   []float64
		result float64
	}{
		"standard case": {
			vals:   []float64{1, 2, 3, 4, 5},
			result: 1.4142135623730951,
		},
		"small variance": {
			vals:   []float64{1, 2, 1, 2, 1, 2, 1, 2},
			result: 0.5,
		},
		"no deviation": {
			vals:   []float64{5, 5, 5},
			result: 0,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.result, StdDev(tc.vals))
		})
	}

}

func TestAvg(t *testing.T) {
	testCases := map[string]struct {
		vals   []float64
		result float64
	}{
		"simple average": {
			vals:   []float64{5, 5, 5},
			result: 5,
		},
		"float average": {
			vals:   []float64{1.2, 3, 4.8},
			result: 3,
		},
		"repeating decimal": {
			vals:   []float64{0.4, 0.4, 0.2},
			result: 0.3333333333333333,
		},
		"one val": {
			vals:   []float64{15},
			result: 15,
		},
		"zeros": {
			vals:   []float64{0, 0, 0, 0},
			result: 0,
		},
		"with negatives": {
			vals:   []float64{-0.4, -0.6, -10, 5},
			result: -1.5,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.result, Avg(tc.vals))
		})
	}
}

func TestSum(t *testing.T) {
	testCases := map[string]struct {
		vals   []float64
		result float64
	}{
		"simple sum": {
			vals:   []float64{1, 2, 3, 4},
			result: 10,
		},
		"float sum": {
			vals:   []float64{1.2, 2.4, 4.6, 6.8},
			result: 15,
		},
		"one": {
			vals:   []float64{1},
			result: 1,
		},
		"zeros": {
			vals:   []float64{0, 0, 0},
			result: 0,
		},
		"with negatives": {
			vals:   []float64{-2.4, 1.2},
			result: -1.2,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.result, Sum(tc.vals))
		})
	}
}
