package shapes

import (
	"math"

	"github.com/jamrig/lumen/internal/lumen/material"
	"github.com/jamrig/lumen/internal/lumen/maths"
)

type Sphere struct {
	Center      maths.Ray
	Radius      float64
	Material    material.Material
	BoundingBox maths.AABB
}

func NewSphere(origin maths.Vec3, radius float64, mat material.Material) Sphere {
	rVec := maths.NewVec3(radius, radius, radius)

	return Sphere{
		Center:      maths.NewRay(origin, maths.NewVec3(0, 0, 0)),
		Radius:      radius,
		Material:    mat,
		BoundingBox: maths.NewAABBFromPoints(origin.Sub(rVec), origin.Add(rVec)),
	}
}

func NewMovingSphere(origin maths.Vec3, end maths.Vec3, radius float64, mat material.Material) Sphere {
	center := maths.NewRay(origin, end.Sub(origin))
	rVec := maths.NewVec3(radius, radius, radius)

	return Sphere{
		Center:   center,
		Radius:   radius,
		Material: mat,
		BoundingBox: maths.NewAABBFromAABBs(
			maths.NewAABBFromPoints(center.At(0).Sub(rVec), center.At(0).Add(rVec)),
			maths.NewAABBFromPoints(center.At(1).Sub(rVec), center.At(1).Add(rVec)),
		),
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

func (s Sphere) GetBoundingBox() maths.AABB {
	return s.BoundingBox
}
