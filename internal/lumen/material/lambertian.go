package material

import "github.com/jamrig/lumen/internal/lumen/maths"

type LambertianMaterial struct {
	Albedo *maths.Color
}

func NewLambertianMaterial(albedo *maths.Color) *LambertianMaterial {
	return &LambertianMaterial{
		Albedo: albedo,
	}
}

func (m *LambertianMaterial) Scatter(hit *maths.Intersection) *maths.ScatteredRay {
	dir := hit.Normal.Add(maths.NewRandomUnitVec3())
	if dir.NearZero() {
		dir = hit.Normal
	}

	r := maths.NewScatteredRay(maths.NewRayWithTime(hit.Point, dir, hit.Ray.Time), m.Albedo)

	return r
}
