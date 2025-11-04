package material

import (
	"math"
	"math/rand/v2"

	"github.com/jamrig/lumen/internal/lumen/maths"
)

type DielectricMaterial struct {
	RefractionIndex float64
}

func NewDielectricMaterial(refractionIndex float64) DielectricMaterial {
	return DielectricMaterial{
		RefractionIndex: refractionIndex,
	}
}

func (m DielectricMaterial) Scatter(hit *maths.Intersection) *maths.ScatteredRay {
	ri := m.RefractionIndex
	attentuate := maths.NewColor(1.0, 1.0, 1.0)

	if hit.FrontFace {
		ri = 1 / ri
	}

	unit := hit.Ray.Direction.Unit()
	cosTheta := math.Min(unit.Mul(-1).Dot(hit.Normal), 1)
	sinTheta := math.Sqrt(1 - cosTheta*cosTheta)

	if ri*sinTheta > 1 || m.Reflectance(cosTheta, ri) > rand.Float64() {
		r := maths.NewScatteredRay(maths.NewRay(hit.Point, unit.Reflect(hit.Normal)), attentuate)
		return &r
	}

	r := maths.NewScatteredRay(maths.NewRay(hit.Point, unit.Refract(hit.Normal, ri)), attentuate)

	return &r
}

func (m DielectricMaterial) Reflectance(cos float64, refractionIndex float64) float64 {
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	r0 *= r0
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
