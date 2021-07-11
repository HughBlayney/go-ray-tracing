package rays

import (
	"fmt"

	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/vectors"
)

func FireRays(screen_vectors [][]*vectors.Vector, origin *vectors.Vector) [][]Ray {
	var screen_rays [][]Ray
	for _, row := range screen_vectors {
		var screen_row []Ray
		for _, screen_vector := range row {
			screen_row = append(screen_row, MakeRay(
				origin, screen_vector.Subtract(origin),
			))
		}
		screen_rays = append(screen_rays, screen_row)
	}

	return screen_rays
}

type Ray struct {
	Origin    *vectors.Vector
	Direction *vectors.Vector
}

func MakeRay(origin *vectors.Vector, direction *vectors.Vector) Ray {
	// Guarantees that the direction vector will be normalised
	direction.Normalise()
	return Ray{origin, direction}
}

func (r Ray) CloseTo(s Ray) bool {
	return r.Origin.CloseTo(s.Origin) && r.Direction.CloseTo(s.Direction)
}

func (r Ray) Print() string {
	return fmt.Sprintf("Origin = %v, Direction = %v", r.Origin, r.Direction)
}
