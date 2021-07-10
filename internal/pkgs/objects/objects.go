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
