package material

import "github.com/jamrig/lumen/internal/lumen/maths"

type Material interface {
	Scatter(hit *maths.Intersection) *maths.ScatteredRay
}
