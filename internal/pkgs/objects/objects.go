package objects

import (
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/rays"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/vectors"
)

type Object interface {
	CollideDistances(rays.Ray) []float64 // Returns slice of distances to intersection of ray with object
	Normal(vectors.Vector) vectors.Vector
	Reflect(rays.Ray) rays.Ray
}

// I don't think this is done correctly, but this is the best way I could think of.
// I don't want to be duplicating code, and this function is common among all shapes.
func ComputeReflectedRay(incoming_ray rays.Ray, point_of_intersection *vectors.Vector, surface_normal *vectors.Vector) rays.Ray {
	new_direction := surface_normal.MultiplyScalar(2 * incoming_ray.Dot(surface_normal)).Subtract(incoming_ray.Direction.MultiplyScalar(-1))
	return rays.Ray{Origin: point_of_intersection, Direction: new_direction}
}
