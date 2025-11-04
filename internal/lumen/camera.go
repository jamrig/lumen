package lumen

import (
	"encoding/json"
	"fmt"
	"image/color"
	"math"
)

type Camera struct {
	Origin         Vec3
	FocalLength    float64
	ViewportWidth  float64
	ViewportHeight float64
	Down           Vec3
	Right          Vec3
	ViewCenter     Vec3
	ViewDeltaDown  Vec3
	ViewDeltaRight Vec3
	ViewTopLeft    Vec3
	ViewPixelStart Vec3
}

func NewCamera(origin Vec3, focalLength float64, viewportWidth float64, viewportHeight float64, imageWidth int, imageHeight int) Camera {
	c := Camera{
		Origin:         origin,
		FocalLength:    focalLength,
		ViewportWidth:  viewportWidth,
		ViewportHeight: viewportHeight,
	}

	c.Right = NewVec3(viewportWidth, 0, 0).Unit()
	c.Down = NewVec3(0, -viewportHeight, 0).Unit()

	c.ViewCenter = c.Origin.Sub(NewVec3(0, 0, c.FocalLength))
	c.ViewDeltaRight = NewVec3(c.ViewportWidth, 0, 0).Div(float64(imageWidth))
	c.ViewDeltaDown = NewVec3(0, -c.ViewportHeight, 0).Div(float64(imageHeight))

	c.ViewTopLeft = c.ViewCenter.Sub(c.Down.Mul(c.ViewportHeight / 2)).Sub(c.Right.Mul(c.ViewportWidth / 2))
	c.ViewPixelStart = c.ViewTopLeft.Add((c.ViewDeltaDown.Add(c.ViewDeltaRight)).Div(2))

	return c
}

func (c *Camera) ColorAtPixel(screenX, screenY int, world HittableList) color.RGBA {
	pixelCenter := c.ViewPixelStart.Add(c.ViewDeltaRight.Mul(float64(screenX)).Add(c.ViewDeltaDown.Mul(float64(screenY))))
	rayDirection := pixelCenter.Sub(c.Origin).Unit()
	ray := NewRay(c.Origin, rayDirection)

	if record := world.Hit(ray, 0, math.Inf(1)); record != nil {
		return NewVec3(1, 1, 1).Add(record.Normal).Mul(0.5).ToRGBA()
	}

	a := 0.5 * (rayDirection.Y + 1.0)
	return NewVec3(1.0, 1.0, 1.0).Mul(1.0 - a).Add(NewVec3(0.5, 0.7, 1.0).Mul(a)).ToRGBA()
}

func (c Camera) String() string {
	pretty, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Sprintf("[ERR] Failed to stringify Camera: %v", err)
	}

	return fmt.Sprintf("Camera: %v", string(pretty))
}
