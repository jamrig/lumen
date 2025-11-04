package lumen

import (
	"encoding/json"
	"fmt"
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

func (s Sphere) Hit(r Ray) bool {
	// fmt.Println(r)
	oc := s.Origin.Sub(r.Origin)
	a := r.Direction.Dot(r.Direction)
	b := r.Direction.Dot(oc) * -2.0
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - 4*a*c
	return discriminant >= 0
}

func (s Sphere) String() string {
	pretty, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Sprintf("[ERR] Failed to stringify Sphere: %v", err)
	}

	return fmt.Sprintf("Sphere: %v", string(pretty))
}
