package rays

import (
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/vectors"
)

type Ray struct {
	Origin    *vectors.Vector
	Direction *vectors.Vector
}
