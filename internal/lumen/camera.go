package lumen

import (
	"encoding/json"
	"fmt"
	"image"
	"math"
	"math/rand"
)

const aspectRatio = 16.0 / 9.0
const imageWidth = 800
const viewportHeight = 2.0
const focalLength = 1.0
const samplesPerPixel = 10

type Camera struct {
	Origin         Vec3
	FocalLength    float64
	ImageWidth     int
	ImageHeight    int
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

func NewCamera() Camera {
	c := Camera{
		Origin:         NewVec3(0, 0, 0),
		FocalLength:    focalLength,
		ImageWidth:     imageWidth,
		ImageHeight:    int(math.Floor(float64(imageWidth) / aspectRatio)),
		ViewportHeight: viewportHeight,
	}

	c.ViewportWidth = 2.0 * (float64(c.ImageWidth) / float64(c.ImageHeight))

	c.Right = NewVec3(c.ViewportWidth, 0, 0).Unit()
	c.Down = NewVec3(0, -viewportHeight, 0).Unit()

	c.ViewCenter = c.Origin.Sub(NewVec3(0, 0, c.FocalLength))
	c.ViewDeltaRight = NewVec3(c.ViewportWidth, 0, 0).Div(float64(c.ImageWidth))
	c.ViewDeltaDown = NewVec3(0, -c.ViewportHeight, 0).Div(float64(c.ImageHeight))

	c.ViewTopLeft = c.ViewCenter.Sub(c.Down.Mul(c.ViewportHeight / 2)).Sub(c.Right.Mul(c.ViewportWidth / 2))
	c.ViewPixelStart = c.ViewTopLeft.Add((c.ViewDeltaDown.Add(c.ViewDeltaRight)).Div(2))

	return c
}

func (c *Camera) Render(world HittableList) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, c.ImageWidth, c.ImageHeight))

	for j := range c.ImageHeight {
		for i := range c.ImageWidth {
			pixelColor := NewVec3(0, 0, 0)

			for range samplesPerPixel {
				r := c.GetRay(i, j)
				pixelColor = pixelColor.Add(c.GetRayColor(r, world))
			}

			img.SetRGBA(i, j, pixelColor.Div(samplesPerPixel).ToRGBA())
		}
	}

	return img
}

func (c *Camera) GetRay(i, j int) Ray {
	offset := NewVec3(rand.Float64()-0.5, rand.Float64()-0.5, 0)
	pixelSample := c.ViewPixelStart.Add(c.ViewDeltaRight.Mul(float64(i) + offset.X)).Add(c.ViewDeltaDown.Mul(float64(j) + offset.Y))

	return NewRay(c.Origin, pixelSample.Sub(c.Origin))
}

func (c *Camera) GetRayColor(r Ray, world HittableList) Vec3 {
	if record := world.Hit(r, NewInterval(0, math.Inf(1))); record != nil {
		return NewVec3(1, 1, 1).Add(record.Normal).Mul(0.5)
	}

	a := 0.5 * (r.Direction.Y + 1.0)
	return NewVec3(1.0, 1.0, 1.0).Mul(1.0 - a).Add(NewVec3(0.5, 0.7, 1.0).Mul(a))
}

func (c Camera) String() string {
	pretty, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("Camera: %v", string(pretty))
}
