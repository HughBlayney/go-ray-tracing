package lights

import (
	"image/color"

	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/vectors"
)

type Light struct {
	Color     color.RGBA
	Intensity float64        // [0.0, 1.0]
	Position  vectors.Vector // temporary
}
