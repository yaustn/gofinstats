package gofinstats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleMovingAverage(t *testing.T) {
	testCases := map[string]struct {
		window int
		vals   []float64
		result float64
	}{
		"partial window": {
			window: 5,
			vals:   []float64{1, 2, 3, 4},
			result: 2.5,
		},
		"filled window": {
			window: 5,
			vals:   []float64{1, 2, 3, 4, 5},
			result: 3,
		},
		"rolling window": {
			window: 5,
			vals:   []float64{10, 15, 20, 1, 2, 3, 4, 5},
			result: 2.999999999999999, // 3
		},
		"overflow rolling window": {
			window: 3,
			vals:   []float64{10, 15, 25, 1, 2, 3, 4, 5},
			result: 4,
		},
		"no window": {
			window: 1,
			vals:   []float64{0.0, 1.2, 2.4, 3},
			result: 3,
		},
		"small rolling window": {
			window: 2,
			vals:   []float64{0.0, 1.2, 2.4, 3},
			result: 2.7,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			sma := NewSimpleMovingAverage(tc.window)

			for _, val := range tc.vals {
				sma.Add(val)
			}

			assert.Equal(t, tc.result, sma.GetAvg())
		})
	}
}
