package main

import (
	"time"

	"github.com/faiface/pixel"
)

// rotationNode performs rotation
// of the object from the starting
// to the destination angle over a
// period of time.
type rotationNode struct {
	passed     time.Duration
	duration   time.Duration
	startAngle float64
	destAngle  float64
}

// Update updates the game object rotation on the way from
// the starting angle to the destination angle.
func (node *rotationNode) Update(transform *pixel.Matrix, deltaTime time.Duration) {
	t := float64(node.passed) / float64(node.duration)

	if t > 1.0 {
		t = 1.0
	}

	position := pixel.V(transform[4], transform[5])
	angle := Lerp(node.startAngle, node.destAngle, t)
	begin := pixel.V(transform[0], transform[1]).Angle()
	rotation := angle - begin

	*transform = transform.Rotated(position, rotation)
	node.passed += deltaTime
}

// Complete returns true if node execution
// is complete, and false otherwise.
func (node *rotationNode) Complete() bool {
	return node.passed >= node.duration
}
