package gofinstats

const (
	smoothingFactor = 2
	defaultPeriod   = 15
)

type ExponentialMovingAverage struct {
	weight  float64 // Calculated as (SmoothingFactor / (1 + period of EMA))
	average float64 // The running EMA
}

// NewExponentialMovingAverage creates and returns a new EMA
func NewExponentialMovingAverage(period int) *ExponentialMovingAverage {
	if period == 0 {
		period = defaultPeriod
	}
	weight := float64(smoothingFactor) / float64(1+period)

	return &ExponentialMovingAverage{
		average: 0,
		weight:  weight,
	}
}

// Add takes the input value and weights it against the current average. The first N iterations will not be reliable
// due to a initialization period in which the EMA has not fully stablized and requires more observations
func (ema *ExponentialMovingAverage) Add(val float64) float64 {
	if ema.average == 0 {
		ema.average = val
	} else {
		ema.average = (val * ema.weight) + (ema.average * (1 - ema.weight))
	}

	return ema.average
}

// GetAverage returns the EMA in O(1) time
func (ema *ExponentialMovingAverage) GetAverage() float64 {
	return ema.average
}
