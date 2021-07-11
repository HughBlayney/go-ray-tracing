package scenes

import (
	"image/color"

	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/lights"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/objects"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/rays"
)

type Scene struct {
	Objects []objects.Object
	Lights  []lights.Light
}

func (s Scene) Render(ray_matrix [][]rays.Ray) (colour_matrix [][]color.RGBA) {
	// Given a matrix of rays, return a matrix of colour values.
	for _, ray_row := range ray_matrix {
		var colour_row []color.RGBA
		for _, ray := range ray_row {
			var colour color.RGBA
			closest_obj, _ := s.ClosestObject(ray)
			if closest_obj != nil {
				//surface_vector := ray.Origin.Add(ray.Direction.MultiplyScalar(dist))
				colour = closest_obj.GetColor()
			}
			colour_row = append(colour_row, colour)
		}
		colour_matrix = append(colour_matrix, colour_row)
	}
	return
}

func (s Scene) ClosestObject(ray rays.Ray) (objects.Object, float64) {
	found_obj := false
	var closest_dist float64
	var closest_obj objects.Object
	for _, obj := range s.Objects {
		dists := obj.CollideDistances(ray)
		for _, d := range dists {
			if d > 0.0 {
				if !found_obj || d < closest_dist {
					closest_dist = d
					closest_obj = obj
				}
			}
		}
	}
	return closest_obj, closest_dist
}
