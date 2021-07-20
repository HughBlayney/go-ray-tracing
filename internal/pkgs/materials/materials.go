package materials

import (
	"image/color"
)

func MakeMaterial(colour color.RGBA, diffuse float64, specular float64, ambient float64, shininess float64, matte float64) Material {
	Diffuse_consts := [3]float64{float64(colour.R) * diffuse, float64(colour.G) * diffuse, float64(colour.B) * diffuse}
	Specular_consts := [3]float64{float64(colour.R) * specular, float64(colour.G) * specular, float64(colour.B) * specular}
	Ambient_consts := [3]float64{float64(colour.R) * ambient, float64(colour.G) * ambient, float64(colour.B) * ambient}

	return Material{
		Color:           colour,
		Diffuse_const:   diffuse,
		Ambient_const:   ambient,
		Shininess_const: shininess,
		Diffuse_consts:  Diffuse_consts,
		Specular_consts: Specular_consts,
		Ambient_consts:  Ambient_consts,
		Matte:           matte,
	}
}

type Material struct {
	Color           color.RGBA
	Specular_const  float64
	Diffuse_const   float64
	Ambient_const   float64
	Shininess_const float64

	Diffuse_consts  [3]float64
	Specular_consts [3]float64
	Ambient_consts  [3]float64

	Matte float64 // \in [0, 1], higher values = less reflection / refraction

	Ambient_color color.RGBA // Only needs to be computed once per scene
}
