package shapes

import (
	"github.com/jamrig/lumen/internal/lumen/material"
	"github.com/jamrig/lumen/internal/lumen/maths"
)

type HitResult struct {
	Intersection maths.Intersection
	Material     material.Material
}

func NewHitResult(i maths.Intersection, m material.Material) *HitResult {
	return &HitResult{
		Intersection: i,
		Material:     m,
	}
}

type Hittable interface {
	Hit(r maths.Ray, t maths.Interval) *HitResult
}
