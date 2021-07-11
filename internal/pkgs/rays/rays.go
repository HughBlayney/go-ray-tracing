package rays

import (
	"fmt"

	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/vectors"
)

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
