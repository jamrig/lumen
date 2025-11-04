package maths

import (
	"math"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}

func NewRandomVec3(min, max float64) Vec3 {
	return NewVec3(RandomDouble(min, max), RandomDouble(min, max), RandomDouble(min, max))
}

func NewRandomUnitVec3() Vec3 {
	for {
		p := NewRandomVec3(-1, 1)
		lensq := p.LengthSquared()
		if math.SmallestNonzeroFloat64 < lensq || lensq <= 1 {
			return p.Div(math.Sqrt(lensq))
		}
	}
}

func NewRandomUnitVec3OnHemisphere(normal Vec3) Vec3 {
	v := NewRandomUnitVec3()
	if v.Dot(normal) > 0.0 {
		return v
	}

	return v.Mul(-1)
}

func (v Vec3) Add(u Vec3) Vec3 {
	return Vec3{
		X: v.X + u.X,
		Y: v.Y + u.Y,
		Z: v.Z + u.Z,
	}
}

func (v Vec3) Sub(u Vec3) Vec3 {
	return Vec3{
		X: v.X - u.X,
		Y: v.Y - u.Y,
		Z: v.Z - u.Z,
	}
}

func (v Vec3) Mul(t float64) Vec3 {
	return Vec3{
		X: v.X * t,
		Y: v.Y * t,
		Z: v.Z * t,
	}
}

func (v Vec3) Div(t float64) Vec3 {
	return Vec3{
		X: v.X / t,
		Y: v.Y / t,
		Z: v.Z / t,
	}
}

func (v Vec3) Dot(u Vec3) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func (v Vec3) Cross(u Vec3) Vec3 {
	return NewVec3(
		v.Y*u.Z-v.Z*u.Y,
		v.Z*u.X-v.X*u.Z,
		v.X*u.Y-v.Y*u.X,
	)
}

func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) Unit() Vec3 {
	return v.Div(v.Length())
}

func (v Vec3) NearZero() bool {
	threshold := 1e-8
	return v.X < threshold && v.Y < threshold && v.Z < threshold
}

func (v Vec3) Reflect(n Vec3) Vec3 {
	return v.Sub(n.Mul(2 * v.Dot(n)))
}

func (v Vec3) Refract(n Vec3, refractionIndex float64) Vec3 {
	cosTheta := math.Min(v.Mul(-1).Dot(n), 1)
	rOutPerp := v.Add(n.Mul(cosTheta)).Mul(refractionIndex)
	rOutParallel := n.Mul(-math.Sqrt(math.Abs(1 - rOutPerp.LengthSquared())))
	return rOutPerp.Add(rOutParallel)
}
