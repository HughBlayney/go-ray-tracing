package objects

import (
	"math"

	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/rays"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/vectors"
)

type Sphere struct {
	Radius float64
	Center vectors.Vector
}

func (s Sphere) CollideDistances(test_ray rays.Ray) []float64 {
	d := test_ray.Direction.Magnitude()
	oc := test_ray.Origin.Subtract(s.Center).Magnitude()

	a := d * d
	b := 2 * d * oc
	c := oc * oc

	delta := b*b - 4*a*c

	var distances []float64

	if delta > 0.0 {
		distances = append(distances, (-b+math.Sqrt(delta))/(2*a))
		distances = append(distances, (-b-math.Sqrt(delta))/(2*a))
	} else if delta == 0.0 {
		distances = append(distances, (-b)/(2*a))
	}

	return distances
}

func (s Sphere) Normal(surface_point *vectors.Vector) *vectors.Vector {
	return surface_point.Subtract(s.Center).Normalise()
}

func (s Sphere) Reflect(incoming_ray rays.Ray, point_of_intersection *vectors.Vector) rays.Ray {
	return ComputeReflectedRay(incoming_ray, point_of_intersection, s.Normal(point_of_intersection))
}
