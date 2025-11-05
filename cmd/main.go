package main

import (
	"flag"
	"image/png"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/jamrig/lumen/internal/lumen"
	"github.com/jamrig/lumen/internal/lumen/material"
	"github.com/jamrig/lumen/internal/lumen/maths"
	"github.com/jamrig/lumen/internal/lumen/shapes"
)

var samples = flag.Int("samples", 1, "number of sampling rays")
var aspectRatio = flag.Float64("aspect", 16.0/9.0, "aspect ratio of the render")
var width = flag.Int("width", 1000, "width of the render")
var maxDepth = flag.Int("maxdepth", 50, "max ray depth")

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var heapprofile = flag.String("heapprofile", "", "write heap profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if *heapprofile != "" {
		f, err := os.Create(*heapprofile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		runtime.GC()
		if err := pprof.Lookup("allocs").WriteTo(f, 0); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}

	camera := lumen.NewCamera(*width, *aspectRatio, *samples, *maxDepth)

	// scene
	scene := lumen.NewScene()
	materialGround := material.NewLambertianMaterial(maths.NewColor(0.5, 0.5, 0.5))
	scene.Add(shapes.NewSphere(maths.NewVec3(0.0, -1000, 0), 1000, materialGround))
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := maths.RandomDouble(0, 1)
			center := maths.NewVec3(float64(a)+0.9*maths.RandomDouble(0, 1), 0.2, float64(b)+0.9*maths.RandomDouble(0, 1))

			if (center.Sub(maths.NewVec3(4, 0.2, 0))).Length() > 0.9 {
				var mat material.Material

				if chooseMat < 0.8 {
					albedo := maths.NewColor(maths.RandomDouble(0, 1), maths.RandomDouble(0, 1), maths.RandomDouble(0, 1))
					mat = material.NewLambertianMaterial(albedo)
				} else if chooseMat < 0.95 {
					albedo := maths.NewColor(maths.RandomDouble(0.5, 1), maths.RandomDouble(0.5, 1), maths.RandomDouble(0.5, 1))
					fuzz := maths.RandomDouble(0, 0.5)
					mat = material.NewMetalMaterial(albedo, fuzz)
				} else {
					mat = material.NewDielectricMaterial(1.5)
				}

				scene.Add(shapes.NewSphere(center, 0.2, mat))
			}
		}
	}

	material1 := material.NewDielectricMaterial(1.5)
	scene.Add(shapes.NewSphere(maths.NewVec3(0, 1, 0), 1.0, material1))

	material2 := material.NewLambertianMaterial(maths.NewColor(0.4, 0.2, 0.1))
	scene.Add(shapes.NewSphere(maths.NewVec3(-4, 1, 0), 1.0, material2))

	material3 := material.NewMetalMaterial(maths.NewColor(0.7, 0.6, 0.5), 0.0)
	scene.Add(shapes.NewSphere(maths.NewVec3(4, 1, 0), 1.0, material3))

	img := camera.RenderParallel(scene)

	file, err := os.Create("output.png")
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatalf("failed to encode image: %v", err)
	}
}
