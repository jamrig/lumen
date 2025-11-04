package shapes

import (
	"github.com/jamrig/lumen/internal/lumen/maths"
)

type Hittable interface {
	Hit(r maths.Ray, t maths.Interval) *maths.Intersection
}
