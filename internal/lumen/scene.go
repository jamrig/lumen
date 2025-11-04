package lumen

import (
	"github.com/jamrig/lumen/internal/lumen/maths"
	"github.com/jamrig/lumen/internal/lumen/shapes"
)

type Scene struct {
	Objects []shapes.Hittable
}

func NewScene() *Scene {
	return &Scene{}
}

func (s *Scene) Clear() {
	s.Objects = make([]shapes.Hittable, 0)
}

func (s *Scene) Add(object shapes.Hittable) {
	s.Objects = append(s.Objects, object)
}

func (s *Scene) Hit(r maths.Ray, t maths.Interval) *shapes.HitResult {
	closest := maths.NewInterval(t.Min, t.Max)
	var res *shapes.HitResult

	for _, obj := range s.Objects {
		if newRes := obj.Hit(r, closest); newRes != nil {
			res = newRes
			closest.Max = res.Intersection.T
		}
	}

	return res
}
