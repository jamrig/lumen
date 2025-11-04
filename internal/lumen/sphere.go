package lumen

import (
	"encoding/json"
	"fmt"
	"math"
)

type Sphere struct {
	Origin Vec3
	Radius float64
}

func NewSphere(origin Vec3, radius float64) Sphere {
	return Sphere{
		Origin: origin,
		Radius: radius,
	}
}

func (s Sphere) Hit(r Ray) float64 {
	oc := s.Origin.Sub(r.Origin)
	a := r.Direction.LengthSquared()
	h := r.Direction.Dot(oc)
	c := oc.LengthSquared() - s.Radius*s.Radius
	discriminant := h*h - a*c

	if discriminant < 0 {
		return -1.0
	}

	return (h - math.Sqrt(discriminant)) / a
}

func (s Sphere) String() string {
	pretty, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Sprintf("[ERR] Failed to stringify Sphere: %v", err)
	}

	return fmt.Sprintf("Sphere: %v", string(pretty))
}
