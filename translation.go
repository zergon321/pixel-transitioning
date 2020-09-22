package main

import (
	"time"

	"github.com/faiface/pixel"
)

// translationNode performs
// translation of the object
// from the starting point to
// the destination point over
// a period of time.
type translationNode struct {
	passed      time.Duration
	duration    time.Duration
	start       pixel.Vec
	destination pixel.Vec
}

// Update updates the position of the game object on the way
// from the starting point to the destination point.
func (node *translationNode) Update(transform *pixel.Matrix, deltaTime time.Duration) {
	t := float64(node.passed) / float64(node.duration)

	if t > 1.0 {
		t = 1.0
	}

	position := pixel.Lerp(node.start, node.destination, t)
	begin := pixel.V(transform[4], transform[5])
	offset := position.Sub(begin)

	*transform = transform.Moved(offset)
	node.passed += deltaTime
}

// Complete returns true if node execution
// is complete, and false otherwise.
func (node *translationNode) Complete() bool {
	return node.passed >= node.duration
}
