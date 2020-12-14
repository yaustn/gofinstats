package gofinstats

type MovingAverage interface {
	Add(val float64) float64
	GetAverage() float64
}
