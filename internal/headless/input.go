//+build headless

package headless

import (
	"github.com/split-cube-studios/ardent/engine/input"
	"github.com/split-cube-studios/ardent/engine/input/raw"
)

var _ input.KeySource = (*KeySource)(nil)
var _ input.MouseSource = (*MouseSource)(nil)

type KeySource struct{}

func (s *KeySource) IsAnyPressed() bool {
	return false
}

func (s *KeySource) IsAnyJustPressed() bool {
	return false
}

func (s *KeySource) StateOf(in input.Input) input.State {
	return input.State{
		Type:        input.Keyboard,
		Input:       in,
		Value:       0.0,
		JustPressed: false,
	}
}

func (s *KeySource) IsPressed(in input.Input) bool {
	return false
}

func (s *KeySource) IsJustPressed(in input.Input) bool {
	return false
}

func (s *KeySource) IsJustReleased(in input.Input) bool {
	return false
}

// Input is an engine.Input.
type MouseSource struct {
	minX, minY, maxX, maxY int
	lcx, lcy               int
	vcx, vcy               int
}

func (s *MouseSource) IsAnyPressed() bool {
	return false
}

func (s *MouseSource) IsAnyJustPressed() bool {
	return false
}

func (s *MouseSource) StateOf(in input.Input) input.State {
	return input.State{
		Type:        input.MouseButton,
		Input:       in,
		Value:       0.0,
		JustPressed: false,
	}
}

func (s *MouseSource) IsPressed(in input.Input) bool {
	return false
}

func (s *MouseSource) IsJustPressed(in input.Input) bool {
	return false
}

func (s *MouseSource) IsJustReleased(in input.Input) bool {
	return false
}

// CursorPosition implements engine.Input.
func (i *MouseSource) CursorPosition() (int, int) {
	return 0, 0
}

// SetCursorBounds implements engine.Input.
func (i *MouseSource) SetCursorBounds(minX, minY, maxX, maxY int) {}

// SetCursorMode implements engine.Input.
func (i *MouseSource) SetCursorMode(mode raw.CursorMode) {}
