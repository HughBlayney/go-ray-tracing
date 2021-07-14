package objects

import (
	"math"

	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/materials"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/rays"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/vectors"
)

type Sphere struct {
	Radius   float64
	Center   vectors.Vector
	Material materials.Material
}

func (s Sphere) CollideDistances(test_ray rays.Ray) []float64 {
	// Always returns distances in ascending order
	B := test_ray.Direction
	AC := test_ray.Origin.Subtract(&s.Center)

	a := B.Dot(B)
	b := 2 * B.Dot(AC)
	c := AC.Dot(AC) - s.Radius*s.Radius

	delta := b*b - 4*a*c

	var distances []float64

	if delta > 0.0 {
		// This guarantees in ascending order
		distances = append(distances, (-b-math.Sqrt(delta))/(2*a))
		distances = append(distances, (-b+math.Sqrt(delta))/(2*a))
	} else if delta == 0.0 {
		distances = append(distances, (-b)/(2*a))
	}

	return distances
}

func (s Sphere) Normal(surface_point *vectors.Vector) *vectors.Vector {
	normal := surface_point.Subtract(&s.Center)
	normal.Normalise()
	return normal
}

func (s Sphere) Reflect(incoming_ray rays.Ray, point_of_intersection *vectors.Vector) rays.Ray {
	return ComputeReflectedRay(incoming_ray, point_of_intersection, s.Normal(point_of_intersection))
}

func (s Sphere) GetMaterial() materials.Material {
	return s.Material
}
