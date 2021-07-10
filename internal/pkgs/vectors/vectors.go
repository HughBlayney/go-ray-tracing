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

func (v *Vector) Add(u *Vector) *Vector {
	return &Vector{X: v.X + u.X, Y: v.Y + u.Y, Z: v.Z + u.Z}
}

func (v *Vector) Subtract(u *Vector) *Vector {
	return &Vector{X: v.X - u.X, Y: v.Y - u.Y, Z: v.Z - u.Z}
}

func (v *Vector) MultiplyScalar(scalar float64) *Vector {
	return &Vector{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}

func (v *Vector) Dot(u *Vector) float64 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}
