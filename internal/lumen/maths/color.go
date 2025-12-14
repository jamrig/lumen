package maths

import (
	"image/color"
)

type Color struct {
	R float64
	G float64
	B float64
}

func NewColor(r, g, b float64) *Color {
	return &Color{
		R: r,
		G: g,
		B: b,
	}
}

func (c *Color) ToRGBA() color.RGBA {
	intensity := NewInterval(0.000, 0.999)

	return color.RGBA{
		R: uint8(intensity.Clamp(LinearToGamma(c.R)) * 256),
		G: uint8(intensity.Clamp(LinearToGamma(c.G)) * 256),
		B: uint8(intensity.Clamp(LinearToGamma(c.B)) * 256),
		A: 255,
	}
}

func (c *Color) Add(k *Color) *Color {
	return &Color{
		R: c.R + k.R,
		G: c.G + k.G,
		B: c.B + k.B,
	}
}

func (c *Color) Sub(k *Color) *Color {
	return &Color{
		R: c.R - k.R,
		G: c.G - k.G,
		B: c.B - k.B,
	}
}

func (c *Color) Mul(t float64) *Color {
	return &Color{
		R: c.R * t,
		G: c.G * t,
		B: c.B * t,
	}
}

func (c *Color) Div(t float64) *Color {
	return &Color{
		R: c.R / t,
		G: c.G / t,
		B: c.B / t,
	}
}

func (c *Color) Attenuate(a *Color) *Color {
	return &Color{
		R: c.R * a.R,
		G: c.G * a.G,
		B: c.B * a.B,
	}
}
