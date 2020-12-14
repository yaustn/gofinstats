package gofinstats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExponentialMovingAverage(t *testing.T) {
	period := 10
	testCases := map[string]struct {
		vals   []float64
		result float64
	}{
		"only one value": {
			vals:   []float64{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			result: 10,
		},
		"converging to one value": {
			vals:   []float64{1, 3, 5, 7, 9, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			result: 9.759373172890896,
		},
		"incrementing": {
			vals:   []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			result: 10.77108785032491,
		},
		"decrementing": {
			vals:   []float64{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			result: 5.228912149675083,
		},
		"up and down": {
			vals:   []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			result: 4.142744317401291,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ema := NewExponentialMovingAverage(period)

			for _, val := range tc.vals {
				ema.Add(val)
			}

			assert.Equal(t, tc.result, ema.GetAverage())
		})
	}
}
