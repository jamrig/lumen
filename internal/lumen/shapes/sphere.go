package shapes

import (
	"math"

	"github.com/jamrig/lumen/internal/lumen/material"
	"github.com/jamrig/lumen/internal/lumen/maths"
)

type Sphere struct {
	Center   maths.Ray
	Radius   float64
	Material material.Material
}

func NewSphere(origin maths.Vec3, radius float64, mat material.Material) Sphere {
	return Sphere{
		Center:   maths.NewRay(origin, maths.NewVec3(0, 0, 0)),
		Radius:   radius,
		Material: mat,
	}
}

func NewMovingSphere(origin maths.Vec3, end maths.Vec3, radius float64, mat material.Material) Sphere {
	return Sphere{
		Center:   maths.NewRay(origin, end.Sub(origin)),
		Radius:   radius,
		Material: mat,
	}
}

func (s Sphere) Hit(r maths.Ray, t maths.Interval) *HitResult {
	currCenter := s.Center.At(r.Time)
	oc := currCenter.Sub(r.Origin)
	a := r.Direction.LengthSquared()
	h := r.Direction.Dot(oc)
	c := oc.LengthSquared() - s.Radius*s.Radius
	discriminant := h*h - a*c

	if discriminant < 0 {
		return nil
	}

	sqrtD := math.Sqrt(discriminant)
	root := (h - sqrtD) / a

	if !t.Surrounds(root) {
		root = (h + sqrtD) / a
		if !t.Surrounds(root) {
			return nil
		}
	}

	p := r.At(root)

	return NewHitResult(
		maths.NewIntersection(r, p, root, p.Sub(currCenter).Div(s.Radius)),
		s.Material,
	)
}
