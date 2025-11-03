package lumen

import (
	"image/color"
)

type Color struct {
	R float64
	G float64
	B float64
}

func NewColor(r, g, b float64) Color {
	return Color{
		R: r,
		G: g,
		B: b,
	}
}

func (c Color) ToRGBA() color.RGBA {
	return color.RGBA{
		R: uint8(c.R * 255.999),
		G: uint8(c.G * 255.999),
		B: uint8(c.B * 255.999),
		A: 255,
	}
}
