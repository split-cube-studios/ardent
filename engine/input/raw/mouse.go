package raw

import "github.com/split-cube-studios/ardent/engine/input"

type MouseButton = int

const (
	MouseButtonLeft MouseButton = iota
	MouseButtonRight
	MouseButtonMiddle
)

// CursorMode indicates a cursor display mode.
type CursorMode byte

const (
	// CursorModeVisible indicates normal cursor display.
	CursorModeVisible CursorMode = 1 << iota

	// CursorModeHidden indicates a hidden cursor that may escape the window.
	CursorModeHidden

	// CursorModeCaptured indicates a hidden cursor that may not escape the window.
	CursorModeCaptured
)

type MouseButtonInput interface {
	input.Source

	CursorPosition() (int, int)
	SetCursorBounds(int, int, int, int)
	SetCursorMode(CursorMode)
}
