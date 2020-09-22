package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
)

// Transition is a series of parallel
// actions organised as tracks with
// instruction nodes.
type Transition struct {
	transform *pixel.Matrix
	tracks    []*transitionTrack
}

// transitionTrack is a queue of
// transition nodes.
//
// When execution of a certain
// node is complete, it's removed
// from the track and execution of
// the next node is started.
type transitionTrack struct {
	name  string
	nodes []TransitionNode
}

// TransitionNode contains a single instruction
// that must be performed over a period of time.
type TransitionNode interface {
	Update(*pixel.Matrix, time.Duration)
	Complete() bool
}

// Update updates the state of all tracks of the transition.
func (tr *Transition) Update(deltaTime time.Duration) {
	for _, track := range tr.tracks {
		if len(track.nodes) <= 0 {
			continue
		}

		if track.nodes[0].Complete() {
			track.nodes = track.nodes[1:]
			continue
		}

		track.nodes[0].Update(tr.transform, deltaTime)
	}
}

// AddTrack adds a new track in the end of the track list.
//
// It returns an error if a track with the same name already exists.
func (tr *Transition) AddTrack(name string) error {
	for _, track := range tr.tracks {
		if track.name == name {
			return fmt.Errorf("track with name '%s' already exists", name)
		}
	}

	track := &transitionTrack{
		name:  name,
		nodes: []TransitionNode{},
	}

	tr.tracks = append(tr.tracks, track)

	return nil
}

// AddTransitionNode adds a new node in the track with the specified name.
//
// It returns an error if there's no track with the specified name.
func (tr *Transition) AddTransitionNode(trackName string, node TransitionNode) error {
	ind := -1

	for i, track := range tr.tracks {
		if track.name == trackName {
			ind = i
			break
		}
	}

	if ind < 0 {
		return fmt.Errorf("track '%s' doesn't exist", trackName)
	}

	track := tr.tracks[ind]
	track.nodes = append(track.nodes, node)

	return nil
}
