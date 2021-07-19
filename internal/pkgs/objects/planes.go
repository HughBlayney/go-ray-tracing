package objects

import (
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/materials"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/rays"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/vectors"
)

type Plane struct {
	PlaneNormal vectors.Vector
	Point       vectors.Vector
	Material    materials.Material
}

func (p Plane) CollideDistances(test_ray rays.Ray) []float64 {
	var distances []float64

	denominator := test_ray.Direction.Dot(&p.PlaneNormal)

	if denominator != 0 {
		numerator := (&p.Point).Subtract(test_ray.Origin).Dot(&p.PlaneNormal)
		d := numerator / denominator
		distances = append(distances, d)
	}

	return distances
}

func (p Plane) Normal(surface_point *vectors.Vector) *vectors.Vector {
	return &p.PlaneNormal
}

func (p Plane) Reflect(incoming_ray rays.Ray, point_of_intersection *vectors.Vector) rays.Ray {
	return ComputeReflectedRay(incoming_ray, point_of_intersection, p.Normal(point_of_intersection))
}

func (p Plane) GetMaterial() materials.Material {
	return p.Material
}
