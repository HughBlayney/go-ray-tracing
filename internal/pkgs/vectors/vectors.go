package vectors

import "math"

type Vector struct {
	X float64
	Y float64
	Z float64
}

func (v *Vector) Magnitude() (magnitude float64) {
	magnitude = math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	return
}

func (v *Vector) Normalise() {
	m := v.Magnitude()
	if m != 0.0 {
		v.X = v.X / m
		v.Y = v.Y / m
		v.Z = v.Z / m
	}
}
	}
}
