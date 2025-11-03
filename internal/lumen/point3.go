package lumen

import "fmt"

type Point3 struct {
	X float64
	Y float64
	Z float64
}

func NewPoint3(x, y, z float64) Point3 {
	return Point3{
		X: x,
		Y: y,
		Z: z,
	}
}

func (p Point3) String() string {
	return fmt.Sprintf("Point3(%v, %v, %v)", p.X, p.Y, p.Z)
}

func (p Point3) Add(q Point3) Point3 {
	return Point3{
		X: p.X + q.X,
		Y: p.Y + q.Y,
		Z: p.Z + q.Z,
	}
}

func (p Point3) Sub(q Point3) Point3 {
	return Point3{
		X: p.X - q.X,
		Y: p.Y - q.Y,
		Z: p.Z - q.Z,
	}
}
