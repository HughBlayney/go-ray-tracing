package scenes

import (
	"image/color"
	"math"

	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/lights"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/materials"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/objects"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/rays"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/vectors"
)

type Scene struct {
	Objects       []objects.Object
	Lights        []lights.Light
	AmbientColour color.RGBA
}

func (s Scene) ObjectsOtherThan(avoid_obj objects.Object) []objects.Object {
	// Return all objects in a scene other than the specified object.
	var return_objects []objects.Object
	for _, obj := range s.Objects {
		if obj != avoid_obj {
			return_objects = append(return_objects, obj)
		}
	}

	return return_objects
}

func (s Scene) Render(ray_matrix [][]rays.Ray) (colour_matrix [][]color.RGBA) {
	// Given a matrix of rays, return a matrix of colour values.
	for _, ray_row := range ray_matrix {
		var colour_row []color.RGBA
		for _, ray := range ray_row {
			var colour color.RGBA
			closest_obj, dist := s.ClosestObject(ray)
			if closest_obj != nil {
				surface_vector := ray.Origin.Add(ray.Direction.MultiplyScalar(dist))
				mat := closest_obj.GetMaterial()
				colour = ComputePhong(
					mat,
					s.Lights,
					s.ObjectsOtherThan(closest_obj),
					s.AmbientColour,
					surface_vector,
					closest_obj.Normal(surface_vector),
					ray.Direction.MultiplyScalar(-1),
				)
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
					found_obj = true
				}
			}
		}
	}
	return closest_obj, closest_dist
}

func computeDiffuseSpecular(
	Diffuse_const [3]float64,
	Specular_const [3]float64,
	alpha float64,
	light_position *vectors.Vector,
	surface_position *vectors.Vector,
	surface_normal *vectors.Vector,
	viewer_direction *vectors.Vector,
	objects []objects.Object,
) ([3]float64, [3]float64) {
	L := (light_position).Subtract(surface_position)
	light_dist := L.Magnitude()
	L.Normalise()
	L_ray := rays.MakeRay(surface_position, L)

	var diffuse [3]float64
	var specular [3]float64
	in_shadow := false
	for _, obj := range objects {
		obj_dists := obj.CollideDistances(L_ray)
		for _, obj_dist := range obj_dists {
			if obj_dist < light_dist {
				in_shadow = true
				break
			}
		}
		if in_shadow {
			break
		}
	}
	if !in_shadow {
		R := L.MultiplyScalar(-1).Reflect(surface_normal)
		R.Normalise()
		diffuse_dot := L.Dot(surface_normal)
		specular_dot := R.Dot(viewer_direction)
		for i := range diffuse {
			if specular_dot > 0.0 {
				specular[i] = Specular_const[i] * math.Pow(specular_dot, alpha)
			} else {
				specular[i] = 0.0
			}
			if diffuse_dot > 0.0 {
				diffuse[i] = Diffuse_const[i] * diffuse_dot
			} else {
				diffuse[i] = 0.0
				specular[i] = 0.0
			}
		}
	}

	return diffuse, specular
}

func clipFloat(float_val float64) uint8 {
	var converted_float_val uint8
	if float_val > 255.0 {
		converted_float_val = 0xff
	} else {
		converted_float_val = uint8(float_val)
	}

	return converted_float_val
}

func ComputePhong(
	m materials.Material, lights []lights.Light, objects []objects.Object, Ambient_color color.RGBA, surface_position *vectors.Vector, surface_normal *vectors.Vector, viewer_direction *vectors.Vector,
) (illumination color.RGBA) {
	illumination.A = 0xff

	if m.Ambient_color.A == 0 {
		m.Ambient_color.A = 0xff
		m.Ambient_color.R = clipFloat(float64(Ambient_color.R) * m.Ambient_consts[0])
		m.Ambient_color.G = clipFloat(float64(Ambient_color.B) * m.Ambient_consts[1])
		m.Ambient_color.B = clipFloat(float64(Ambient_color.B) * m.Ambient_consts[2])
	}
	var light_totals [3]float64
	for _, light := range lights {
		diffuse, specular := computeDiffuseSpecular(
			m.Diffuse_consts,
			m.Specular_consts,
			m.Shininess_const,
			&light.Position,
			surface_position,
			surface_normal,
			viewer_direction,
			objects,
		)
		light_totals[0] += float64(light.Color.R) * (diffuse[0] + specular[0])
		light_totals[1] += float64(light.Color.G) * (diffuse[1] + specular[1])
		light_totals[2] += float64(light.Color.B) * (diffuse[2] + specular[2])
	}
	illumination.R = clipFloat(
		float64(m.Ambient_color.R)*m.Ambient_consts[0] + light_totals[0],
	)
	illumination.G = clipFloat(
		float64(m.Ambient_color.G)*m.Ambient_consts[1] + light_totals[1],
	)
	illumination.B = clipFloat(
		float64(m.Ambient_color.B)*m.Ambient_consts[2] + light_totals[2],
	)

	return
}
