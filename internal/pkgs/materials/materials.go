package materials

import (
	"image/color"
	"math"

	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/lights"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/vectors"
)

func computeDiffuseSpecular(
	Diffuse_const [3]float64,
	Specular_const [3]float64,
	alpha float64,
	light_position *vectors.Vector,
	surface_position *vectors.Vector,
	surface_normal *vectors.Vector,
	viewer_direction *vectors.Vector,
) ([3]float64, [3]float64) {
	L := (light_position).Subtract(surface_position)
	L.Normalise()
	R := L.MultiplyScalar(-1).Reflect(surface_normal)
	R.Normalise()

	var diffuse [3]float64
	var specular [3]float64
	for i := range diffuse {
		diffuse_dot := L.Dot(surface_normal)
		specular_dot := math.Pow(R.Dot(viewer_direction), alpha)
		if specular_dot > 0.0 {
			specular[i] = Specular_const[i] * specular_dot
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

func MakeMaterial(colour color.RGBA, diffuse float64, specular float64, ambient float64, shininess float64) Material {
	diffuse_consts := [3]float64{float64(colour.R) * diffuse, float64(colour.G) * diffuse, float64(colour.B) * diffuse}
	specular_consts := [3]float64{float64(colour.R) * specular, float64(colour.G) * specular, float64(colour.B) * specular}
	ambient_consts := [3]float64{float64(colour.R) * diffuse, float64(colour.G) * diffuse, float64(colour.B) * diffuse}

	return Material{
		Color:           colour,
		Diffuse_const:   diffuse,
		Ambient_const:   ambient,
		Shininess_const: shininess,
		diffuse_consts:  diffuse_consts,
		specular_consts: specular_consts,
		ambient_consts:  ambient_consts,
	}
}

type Material struct {
	Color           color.RGBA
	Specular_const  float64
	Diffuse_const   float64
	Ambient_const   float64
	Shininess_const float64

	diffuse_consts  [3]float64
	specular_consts [3]float64
	ambient_consts  [3]float64

	ambient_color color.RGBA // Only needs to be computed once per scene
}

func (m Material) ComputePhong(
	lights []lights.Light, ambient_color color.RGBA, surface_position *vectors.Vector, surface_normal *vectors.Vector, viewer_direction *vectors.Vector,
) (illumination color.RGBA) {
	illumination.A = 0xff

	if m.ambient_color.A == 0 {
		m.ambient_color.A = 0xff
		m.ambient_color.R = clipFloat(float64(ambient_color.R) * m.ambient_consts[0])
		m.ambient_color.G = clipFloat(float64(ambient_color.B) * m.ambient_consts[1])
		m.ambient_color.B = clipFloat(float64(ambient_color.B) * m.ambient_consts[2])
	}
	var light_totals [3]float64
	for _, light := range lights {
		diffuse, specular := computeDiffuseSpecular(
			m.diffuse_consts,
			m.specular_consts,
			m.Shininess_const,
			&light.Position,
			surface_position,
			surface_normal,
			viewer_direction,
		)
		light_totals[0] += float64(light.Color.R) * (diffuse[0] + specular[0])
		light_totals[1] += float64(light.Color.G) * (diffuse[1] + specular[1])
		light_totals[2] += float64(light.Color.B) * (diffuse[2] + specular[2])
	}
	illumination.R = clipFloat(
		float64(m.ambient_color.R) + light_totals[0],
	)
	illumination.G = clipFloat(
		float64(m.ambient_color.G) + light_totals[1],
	)
	illumination.B = clipFloat(
		float64(m.ambient_color.B) + light_totals[2],
	)

	return
}
