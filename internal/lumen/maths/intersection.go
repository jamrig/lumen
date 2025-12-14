package maths

type Intersection struct {
	Ray       *Ray
	Point     *Vec3
	T         float64
	Normal    *Vec3
	FrontFace bool
}

func NewIntersection(r *Ray, p *Vec3, t float64, n *Vec3) *Intersection {
	intersection := Intersection{
		Ray:    r,
		Point:  p,
		T:      t,
		Normal: n,
	}

	intersection.FrontFace = r.Direction.Dot(n) < 0

	if !intersection.FrontFace {
		intersection.Normal = intersection.Normal.Mul(-1.0)
	}

	return &intersection
}
