package lumen

import (
	"image"
	"math"
	"math/rand"

	"github.com/jamrig/lumen/internal/lumen/maths"
)

const aspectRatio = 16.0 / 9.0
const imageWidth = 1000
const viewportHeight = 2.0
const focalLength = 1.0
const samplesPerPixel = 100
const maxDepth = 50
const intersectionThreshold = 0.001

type Camera struct {
	Origin         maths.Vec3
	FocalLength    float64
	ImageWidth     int
	ImageHeight    int
	ViewportWidth  float64
	ViewportHeight float64
	Down           maths.Vec3
	Right          maths.Vec3
	ViewCenter     maths.Vec3
	ViewDeltaDown  maths.Vec3
	ViewDeltaRight maths.Vec3
	ViewTopLeft    maths.Vec3
	ViewPixelStart maths.Vec3
}

func NewCamera() Camera {
	c := Camera{
		Origin:         maths.NewVec3(0, 0, 0),
		FocalLength:    focalLength,
		ImageWidth:     imageWidth,
		ImageHeight:    int(math.Floor(float64(imageWidth) / aspectRatio)),
		ViewportHeight: viewportHeight,
	}

	c.ViewportWidth = 2.0 * (float64(c.ImageWidth) / float64(c.ImageHeight))

	c.Right = maths.NewVec3(c.ViewportWidth, 0, 0).Unit()
	c.Down = maths.NewVec3(0, -viewportHeight, 0).Unit()

	c.ViewCenter = c.Origin.Sub(maths.NewVec3(0, 0, c.FocalLength))
	c.ViewDeltaRight = maths.NewVec3(c.ViewportWidth, 0, 0).Div(float64(c.ImageWidth))
	c.ViewDeltaDown = maths.NewVec3(0, -c.ViewportHeight, 0).Div(float64(c.ImageHeight))

	c.ViewTopLeft = c.ViewCenter.Sub(c.Down.Mul(c.ViewportHeight / 2)).Sub(c.Right.Mul(c.ViewportWidth / 2))
	c.ViewPixelStart = c.ViewTopLeft.Add((c.ViewDeltaDown.Add(c.ViewDeltaRight)).Div(2))

	return c
}

func (c *Camera) Render(scene *Scene) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, c.ImageWidth, c.ImageHeight))

	for j := range c.ImageHeight {
		for i := range c.ImageWidth {
			pixelColor := maths.NewColor(0, 0, 0)

			for range samplesPerPixel {
				r := c.GetRay(i, j)
				pixelColor = pixelColor.Add(c.GetRayColor(r, scene, maxDepth))
			}

			img.SetRGBA(i, j, pixelColor.Div(samplesPerPixel).ToRGBA())
		}
	}

	return img
}

func (c *Camera) GetRay(i, j int) maths.Ray {
	offset := maths.NewVec3(rand.Float64()-0.5, rand.Float64()-0.5, 0)
	pixelSample := c.ViewPixelStart.Add(c.ViewDeltaRight.Mul(float64(i) + offset.X)).Add(c.ViewDeltaDown.Mul(float64(j) + offset.Y))

	return maths.NewRay(c.Origin, pixelSample.Sub(c.Origin))
}

func (c *Camera) GetRayColor(r maths.Ray, scene *Scene, depth int) maths.Color {
	if depth <= 0 {
		return maths.NewColor(0, 0, 0)
	}

	if res := scene.Hit(r, maths.NewInterval(intersectionThreshold, math.Inf(1))); res != nil {
		if scatterRay := res.Material.Scatter(&res.Intersection); scatterRay != nil {
			return c.GetRayColor(scatterRay.Ray, scene, depth-1).Attenuate(scatterRay.Attenuation)
		}

		return maths.NewColor(0, 0, 0)
	}

	a := 0.5 * (r.Direction.Y + 1.0)
	return maths.NewColor(1.0, 1.0, 1.0).Mul(1.0 - a).Add(maths.NewColor(0.5, 0.7, 1.0).Mul(a))
}
