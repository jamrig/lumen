package shapes

import (
	"math"

	"github.com/jamrig/lumen/internal/lumen/maths"
)

type Sphere struct {
	Origin maths.Vec3
	Radius float64
}

func NewSphere(origin maths.Vec3, radius float64) Sphere {
	return Sphere{
		Origin: origin,
		Radius: radius,
	}
}

func (s Sphere) Hit(r maths.Ray, t maths.Interval) *maths.Intersection {
	oc := s.Origin.Sub(r.Origin)
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
	hit := maths.NewIntersection(r, p, root, p.Sub(s.Origin).Div(s.Radius))

	return &hit
}
