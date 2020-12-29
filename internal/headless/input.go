//+build headless

package headless

import "github.com/split-cube-studios/ardent/engine"

// Input is a headless engine.Input.
type Input struct{}

// IsAnyKeyPressed implements engine.Input.
func (i Input) IsAnyKeyPressed() bool {
	return false
}

// IsAnyKeyJustPressed implements engine.Input.
func (i Input) IsAnyKeyJustPressed() bool {
	return false
}

// IsKeyPressed implements engine.Input.
func (i Input) IsKeyPressed(k int) bool {
	return false
}

// IsKeyJustPressed implements engine.Input.
func (i Input) IsKeyJustPressed(k int) bool {
	return false
}

// IsKeyJustReleased implements engine.Input.
func (i Input) IsKeyJustReleased(k int) bool {
	return false
}

// IsMouseButtonPressed implements engine.Input.
func (i Input) IsMouseButtonPressed(k int) bool {
	return false
}

// IsMouseButtonJustPressed implements engine.Input.
func (i Input) IsMouseButtonJustPressed(k int) bool {
	return false
}

// IsMouseButtonJustReleased implements engine.Input.
func (i Input) IsMouseButtonJustReleased(k int) bool {
	return false
}

// CursorPosition implements engine.Input.
func (i Input) CursorPosition() (int, int) {
	return 0, 0
}

// SetCursorBounds implements engine.Input.
func (i Input) SetCursorBounds(minX, minY, maxX, maxY int) {}

// SetCursorMode implements engine.Input.
func (i Input) SetCursorMode(mode engine.CursorMode) {}
