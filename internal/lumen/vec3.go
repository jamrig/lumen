package lumen

import (
	"encoding/json"
	"fmt"
	"image/color"
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

func (v Vec3) String() string {
	pretty, _ := json.MarshalIndent(v, "", "  ")
	return fmt.Sprintf("Vec3: %v", string(pretty))
}

func (v Vec3) ToRGBA() color.RGBA {
	return color.RGBA{
		R: uint8(v.X * 255),
		G: uint8(v.Y * 255),
		B: uint8(v.Z * 255),
		A: 255,
	}
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
		v.Z*u.Y-v.Y*u.Z,
		v.X*u.Z-v.X*u.Z,
		v.Y*u.X-v.X*u.Y,
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
