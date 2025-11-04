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

func (s *Scene) Hit(r maths.Ray, t maths.Interval) *maths.Intersection {
	closest := maths.NewInterval(t.Min, t.Max)
	var hit *maths.Intersection

	for _, obj := range s.Objects {
		if newHit := obj.Hit(r, closest); newHit != nil {
			hit = newHit
			closest.Max = hit.T
		}
	}

	return hit
}
