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
	R := L.Reflect(surface_normal)
	R.Normalise()

	var diffuse [3]float64
	var specular [3]float64
	for i := range diffuse {
		diffuse[i] = Diffuse_const[i] * L.Dot(surface_normal)
		specular[i] = Specular_const[i] * math.Pow(R.Dot(viewer_direction), alpha)
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

type Material struct {
	Color           color.RGBA
	Specular_const  [3]float64
	Diffuse_const   [3]float64
	Ambient_const   [3]float64
	Shininess_const float64

	ambient_color color.RGBA // Only needs to be computed once per scene
}

func (m Material) ComputePhong(
	lights []lights.Light, ambient_color color.RGBA, surface_position *vectors.Vector, surface_normal *vectors.Vector, viewer_direction *vectors.Vector,
) (illumination color.RGBA) {
	illumination.A = 0xff

	if m.ambient_color.A == 0 {
		m.ambient_color.A = 0xff
		m.ambient_color.R = clipFloat(float64(ambient_color.R) * m.Ambient_const[0])
		m.ambient_color.G = clipFloat(float64(ambient_color.B) * m.Ambient_const[1])
		m.ambient_color.B = clipFloat(float64(ambient_color.B) * m.Ambient_const[2])
	}
	var light_totals [3]float64
	for _, light := range lights {
		diffuse, specular := computeDiffuseSpecular(
			m.Diffuse_const,
			m.Specular_const,
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
