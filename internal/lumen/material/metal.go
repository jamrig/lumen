package material

import "github.com/jamrig/lumen/internal/lumen/maths"

type MetalMaterial struct {
	Albedo maths.Color
}

func NewMetalMaterial(albedo maths.Color) MetalMaterial {
	return MetalMaterial{
		Albedo: albedo,
	}
}

func (m MetalMaterial) Scatter(hit *maths.Intersection) *maths.ScatteredRay {
	dir := hit.Ray.Direction.Reflect(hit.Normal)
	r := maths.NewScatteredRay(maths.NewRay(hit.Point, dir), m.Albedo)

	return &r
}
