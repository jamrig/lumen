package lumen

import (
	"image"
	"math"
	"math/rand"

	"github.com/jamrig/lumen/internal/lumen/maths"
)

const aspectRatio = 16.0 / 9.0
const imageWidth = 1000
const samplesPerPixel = 100
const maxDepth = 50
const intersectionThreshold = 0.001
const verticalFOV = 50.0

var lookFrom = maths.NewVec3(-2, 2, 1)
var lookAt = maths.NewVec3(0, 0, -1)
var viewUp = maths.NewVec3(0, 1, 0)

type Camera struct {
	Center         maths.Vec3
	FocalLength    float64
	ImageWidth     int
	ImageHeight    int
	ViewportWidth  float64
	ViewportHeight float64
	CameraUp       maths.Vec3
	CameraRight    maths.Vec3
	CameraBack     maths.Vec3
	ViewHorizontal maths.Vec3
	ViewVertical   maths.Vec3
	ViewDeltaDown  maths.Vec3
	ViewDeltaRight maths.Vec3
	ViewTopLeft    maths.Vec3
	ViewPixelStart maths.Vec3
}

func NewCamera() Camera {
	c := Camera{
		Center:      lookFrom,
		FocalLength: lookFrom.Sub(lookAt).Length(),
		ImageWidth:  imageWidth,
		ImageHeight: int(math.Floor(float64(imageWidth) / aspectRatio)),
	}

	// Viewport
	theta := maths.DegreesToRadians(verticalFOV)
	h := math.Tan(theta / 2)
	c.ViewportHeight = 2 * h * c.FocalLength
	c.ViewportWidth = c.ViewportHeight * (float64(c.ImageWidth) / float64(c.ImageHeight))

	// Camera Basis
	c.CameraBack = lookFrom.Sub(lookAt).Unit()            // w
	c.CameraRight = viewUp.Cross(c.CameraBack).Unit()     // u
	c.CameraUp = c.CameraBack.Cross(c.CameraRight).Unit() // v

	c.ViewHorizontal = c.CameraRight.Mul(c.ViewportWidth) // u
	c.ViewVertical = c.CameraUp.Mul(-c.ViewportHeight)    // v

	c.ViewDeltaRight = c.ViewHorizontal.Div(float64(c.ImageWidth))
	c.ViewDeltaDown = c.ViewVertical.Div(float64(c.ImageHeight))

	c.ViewTopLeft = c.Center.Sub(c.CameraBack.Mul(c.FocalLength)).Sub(c.ViewHorizontal.Div(2)).Sub(c.ViewVertical.Div(2))
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

	return maths.NewRay(c.Center, pixelSample.Sub(c.Center))
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
