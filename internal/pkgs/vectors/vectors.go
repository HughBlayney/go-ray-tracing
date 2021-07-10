package vectors

import "math"

type Vector struct {
	x float64
	y float64
	z float64
}

func (v *Vector) Magnitude() (magnitude float64) {
	magnitude = math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
	return
}

func (v *Vector) Normalise() {
	m := v.Magnitude()
	if m != 0.0 {
		v.x = v.x / m
		v.y = v.y / m
		v.z = v.z / m
	}
}
