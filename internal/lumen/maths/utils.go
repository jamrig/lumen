package maths

import (
	"math"
	"math/rand"
)

func RandomDouble(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func LinearToGamma(linearComponent float64) float64 {
	if linearComponent > 0 {
		return math.Sqrt(linearComponent)
	}

	return 0
}

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}
