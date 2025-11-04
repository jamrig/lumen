package lumen

import (
	"encoding/json"
	"fmt"
)

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func NewRay(origin Vec3, direction Vec3) Ray {
	return Ray{
		Origin:    origin,
		Direction: direction,
	}
}

func (r Ray) At(t float64) Vec3 {
	return r.Origin.Add(r.Direction.Mul(t))
}

func (r Ray) String() string {
	pretty, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return fmt.Sprintf("[ERR] Failed to stringify Ray: %v", err)
	}

	return fmt.Sprintf("Ray: %v", string(pretty))
}
