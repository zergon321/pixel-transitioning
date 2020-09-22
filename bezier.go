package main

import (
	"time"

	"github.com/faiface/pixel"
	"gonum.org/v1/plot/tools/bezier"
)

// bezierNode makes an object move
// along the given Bezier curve.
type bezierNode struct {
	passed   time.Duration
	duration time.Duration
	curve    bezier.Curve
}

// Update updates the position of the game object on the way
// from the starting point to the destination point.
func (node *bezierNode) Update(transform *pixel.Matrix, deltaTime time.Duration) {
	t := float64(node.passed) / float64(node.duration)

	if t > 1.0 {
		t = 1.0
	}

	position := GonumToPixel(node.curve.Point(t))
	begin := pixel.V(transform[4], transform[5])
	offset := position.Sub(begin)

	*transform = transform.Moved(offset)
	node.passed += deltaTime
}

// Complete returns true if node execution
// is complete, and false otherwise.
func (node *bezierNode) Complete() bool {
	return node.passed >= node.duration
}
