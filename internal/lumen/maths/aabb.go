package maths

type AABB struct {
	X *Interval
	Y *Interval
	Z *Interval
}

func NewAABB(x, y, z *Interval) *AABB {
	return &AABB{
		X: x,
		Y: y,
		Z: z,
	}
}

func NewEmptyAABB() *AABB {
	return &AABB{
		X: NewInterval(0, 0),
		Y: NewInterval(0, 0),
		Z: NewInterval(0, 0),
	}
}

func NewAABBFromAABBs(a, b *AABB) *AABB {
	return &AABB{
		X: NewEnclosedInterval(a.X, b.X),
		Y: NewEnclosedInterval(a.Y, b.Y),
		Z: NewEnclosedInterval(a.Z, b.Z),
	}
}

func NewAABBFromPoints(a, b *Vec3) *AABB {
	var xInterval *Interval
	var yInterval *Interval
	var zInterval *Interval

	if a.X <= b.X {
		xInterval = NewInterval(a.X, b.X)
	} else {
		xInterval = NewInterval(b.X, a.X)
	}

	if a.Y <= b.Y {
		yInterval = NewInterval(a.Y, b.Y)
	} else {
		yInterval = NewInterval(b.Y, a.Y)
	}

	if a.Z <= b.Z {
		zInterval = NewInterval(a.Z, b.Z)
	} else {
		zInterval = NewInterval(b.Z, a.Z)
	}

	return &AABB{
		X: xInterval,
		Y: yInterval,
		Z: zInterval,
	}
}

func (a *AABB) AxisInterval(n int) *Interval {
	if n == 1 {
		return a.Y
	}

	if n == 2 {
		return a.Z
	}

	return a.X
}

func (a *AABB) Hit(r *Ray, t *Interval) bool {
	tMin := t.Min
	tMax := t.Max

	for axis := 0; axis < 3; axis++ {
		origAxis := r.Origin.GetAxis(axis)
		ax := a.AxisInterval(axis)
		adInv := 1.0 / origAxis
		t0 := (ax.Min - origAxis) * adInv
		t1 := (ax.Max - origAxis) * adInv

		if t0 < t1 {
			if t0 > tMin {
				tMin = t0
			}

			if t1 < tMax {
				tMax = t1
			}
		} else {
			if t1 > tMin {
				tMin = t1
			}

			if t0 < tMax {
				tMax = t0
			}
		}

		if tMax <= tMin {
			return false
		}
	}

	return true
}
