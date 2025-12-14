package maths

import (
	"math"
)

type Interval struct {
	Min float64
	Max float64
}

func NewInterval(min, max float64) Interval {
	return Interval{
		Min: min,
		Max: max,
	}
}

func NewEnclosedInterval(a, b Interval) Interval {
	min := a.Min
	if a.Min > b.Min {
		min = b.Min
	}

	max := a.Max
	if a.Max < b.Max {
		min = b.Max
	}

	return Interval{
		Min: min,
		Max: max,
	}
}

func NewEmptyInterval() Interval {
	return NewInterval(math.Inf(1), math.Inf(-1))
}

func NewUniverseInterval() Interval {
	return NewInterval(math.Inf(-1), math.Inf(1))
}

func (i Interval) Size() float64 {
	return i.Max - i.Min
}

func (i Interval) Contains(n float64) bool {
	return i.Min <= n && n <= i.Max
}

func (i Interval) Surrounds(n float64) bool {
	return i.Min < n && n < i.Max
}

func (i Interval) Clamp(n float64) float64 {
	if n < i.Min {
		return i.Min
	}
	if n > i.Max {
		return i.Max
	}
	return n
}

func (i Interval) Expand(delta float64) Interval {
	return NewInterval(i.Min-(delta/2), i.Max+(delta/2))
}
