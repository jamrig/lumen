package maths

type Ray struct {
	Origin    *Vec3
	Direction *Vec3
	Time      float64
}

func NewRay(origin *Vec3, direction *Vec3) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: direction,
		Time:      0,
	}
}

func NewRayWithTime(origin *Vec3, direction *Vec3, time float64) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: direction,
		Time:      time,
	}
}

func (r *Ray) At(t float64) *Vec3 {
	return r.Origin.Add(r.Direction.Mul(t))
}
