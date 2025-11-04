package lumen

import "math/rand"

func RandomDouble(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
