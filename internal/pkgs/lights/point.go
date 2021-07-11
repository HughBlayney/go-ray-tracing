package lights

import (
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/vectors"
)

type PointLight struct {
	Position vectors.Vector
	Light
}
