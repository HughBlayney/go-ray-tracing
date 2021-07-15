package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/lights"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/materials"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/objects"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/rays"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/scenes"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/vectors"
)

const frame_height float64 = 1.0
const frame_width float64 = 1.0

var viewer_vector *vectors.Vector = &vectors.Vector{
	X: 0.0,
	Y: 0.0,
	Z: -10.0,
}

var top_left *vectors.Vector = &vectors.Vector{
	X: -frame_width / 2.0,
	Y: frame_height / 2.0,
	Z: 0.0,
}

const height_res int = 100
const width_res int = 200

func subdivide(start float64, end float64, num_segments int) []float64 {
	if end == start {
		log.Fatal("Start and end values must be different.")
	} else if end < start {
		log.Fatal("Start value must be less than end value.")
	}
	sub_dist := (end - start) / float64(num_segments)

	segments := make([]float64, num_segments)

	current_value := start

	for i := range segments {
		segments[i] = current_value
		current_value += sub_dist
	}

	return segments
}

func main() {
	width := 2000
	height := 1000

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// As an initial test, we'll define the screen as a plane at 0, 0, 1
	xs := subdivide(-2.0, 2.0, width)
	ys := subdivide(-1.0, 1.0, height)
	z := 1.0
	// We need to reverse the ys to go top-down
	for i := len(ys)/2 - 1; i >= 0; i-- {
		opp := len(ys) - 1 - i
		ys[i], ys[opp] = ys[opp], ys[i]
	}
	var screen_vectors [][]*vectors.Vector
	for _, y := range ys {
		var screen_row []*vectors.Vector
		for _, x := range xs {
			screen_row = append(screen_row, &vectors.Vector{X: x, Y: y, Z: z})
		}
		screen_vectors = append(screen_vectors, screen_row)
	}
	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}
	cyanmat := materials.MakeMaterial(
		cyan,
		0.5,
		0,
		0,
		0.005,
	)
	red := color.RGBA{0xff, 0, 0, 0xff}
	redmat := materials.MakeMaterial(
		red,
		0.5,
		0,
		0,
		0.005,
	)
	// And a sphere with radius 1 at 0, 0, 10
	sphere := objects.Sphere{
		Radius:   1.0,
		Center:   vectors.Vector{X: 0.0, Y: 0.0, Z: 100.0},
		Material: cyanmat,
	}
	sphere2 := objects.Sphere{
		Radius:   1.0,
		Center:   vectors.Vector{X: 10.0, Y: 0.0, Z: 100.0},
		Material: redmat,
	}

	screen_rays := rays.FireRays(screen_vectors, viewer_vector)

	scene := scenes.Scene{
		Objects: []objects.Object{sphere, sphere2},
		Lights: []lights.Light{{
			Color:     color.RGBA{0xff, 0xff, 0xff, 0xff},
			Intensity: 1.0,
			Position:  vectors.Vector{10.0, 0.0, 0.0},
		}},
	}

	colour_matrix := scene.Render(screen_rays)

	// Set color for each pixel. Cyan if hits sphere, transparent otherwise.
	for i, row := range colour_matrix {
		for j, color := range row {
			img.Set(j, i, color)
		}
	}

	f, err := os.Create("./images/output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
}
