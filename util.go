package main

import (
	"math"

	"github.com/faiface/pixel"
	"gonum.org/v1/plot/vg"
)

const (
	// RadToDeg is a factor to transfrom radians to degrees.
	RadToDeg float64 = 180.0 / math.Pi
	// DegToRad is a factor to transform degrees to radians.
	DegToRad float64 = math.Pi / 180.0
)

// Lerp interpolates between a and b
// using t (0 <= t <= 1).
func Lerp(a, b, t float64) float64 {
	return (1-t)*a + t*b
}

// GonumToPixel transforms the gonum point into
// a Pixel vector.
func GonumToPixel(point vg.Point) pixel.Vec {
	return pixel.V(float64(point.X), float64(point.Y))
}
