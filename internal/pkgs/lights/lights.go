package lights

import (
	"image/color"
)

type Light struct {
	Color     color.RGBA
	Intensity float64 // [0.0, 1.0]
}
