package lumen

import (
	"encoding/json"
	"fmt"
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

func NewEmptyInterval() Interval {
	return NewInterval(math.Inf(1), math.Inf(-1))
}

func NewUniverseInterval() Interval {
	return NewInterval(math.Inf(-1), math.Inf(1))
}

func (i Interval) String() string {
	pretty, _ := json.MarshalIndent(i, "", "  ")
	return fmt.Sprintf("Interval: %v", string(pretty))
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
