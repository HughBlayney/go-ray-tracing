package utils

import "math"

const tolerance float64 = 0.0001

func Close_enough(x, y float64) bool {
	return math.Abs(x-y) < tolerance
}

func Slice_close_enough(x, y []float64) bool {
	if len(x) != len(y) {
		return false
	} else {
		for i, xval := range x {
			if !(Close_enough(xval, y[i])) {
				return false
			}
		}
	}
	return true
}
