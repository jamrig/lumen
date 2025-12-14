package material

import "github.com/jamrig/lumen/internal/lumen/maths"

type MetalMaterial struct {
	Albedo *maths.Color
	Fuzz   float64
}

func NewMetalMaterial(albedo *maths.Color, fuzz float64) *MetalMaterial {
	return &MetalMaterial{
		Albedo: albedo,
		Fuzz:   fuzz,
	}
}

func (m *MetalMaterial) Scatter(hit *maths.Intersection) *maths.ScatteredRay {
	reflected := hit.Ray.Direction.Reflect(hit.Normal)

	if m.Fuzz > 0 {
		reflected = reflected.Unit().Add(maths.NewRandomUnitVec3().Mul(m.Fuzz))
	}

	if reflected.Dot(hit.Normal) <= 0 {
		return nil
	}

	r := maths.NewScatteredRay(maths.NewRayWithTime(hit.Point, reflected, hit.Ray.Time), m.Albedo)

	return r
}
