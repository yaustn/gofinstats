package gofinstats

type SimpleMovingAverage struct {
	window  int       // window of the sma calculation
	vals    []float64 // stored values to calculate average
	i       int       // index of the last written val in vals
	n       int       // total number of vals processed
	average float64   // the simple moving average
}

// NewSimpleMovingAverage returns
func NewSimpleMovingAverage(window int) *SimpleMovingAverage {
	return &SimpleMovingAverage{
		window:  window,
		vals:    make([]float64, window),
		i:       0,
		n:       0,
		average: 0,
	}
}

// GetAvg returns the SMA in O(1) time
func (sma *SimpleMovingAverage) GetAvg() float64 {
	return sma.average
}

// Add adds the input value to array of values in the SMA and returns the moving average
func (sma *SimpleMovingAverage) Add(val float64) float64 {
	if sma.n == sma.window { // window is filled
		sma.average += (val - sma.vals[sma.i]) / float64(sma.n)
		sma.vals[sma.i] = val
		sma.i = (sma.i + 1) % sma.window

		// To remove any rounding errors, we brute force an O(n) calculation every time
		// the SMA window is filled.
		if sma.i == 0 {
			sma.Recalc()
		}
	} else { // still on the first loop, filling the window
		sma.average = (val + (sma.average * float64(sma.n))) / float64(sma.n+1)
		sma.n++

		sma.vals[sma.i] = val
		sma.i = (sma.i + 1) % sma.window
	}

	return sma.average
}

// Recalc forces a full O(n) loop to recalculate the moving average's average.
// Due to golang's float64 imprecision, a full recalc from time to time is necessary
func (sma *SimpleMovingAverage) Recalc() float64 {
	sum := float64(0)
	for i := 0; i < sma.n; i++ {
		sum += sma.vals[i]
	}

	sma.average = sum / float64(sma.n)

	return sma.average
}
