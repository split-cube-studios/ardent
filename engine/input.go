package engine

// Input describes various input detection methods.
type Input interface {
	// Keyboard
	IsAnyKeyPressed() bool
	IsAnyKeyJustPressed() bool
	IsKeyPressed(int) bool
	IsKeyJustPressed(int) bool
	IsKeyJustReleased(int) bool

	// Mouse
	IsMouseButtonPressed(int) bool
	IsMouseButtonJustPressed(int) bool
	IsMouseButtonJustReleased(int) bool
	CursorPosition() (int, int)
	SetCursorBounds(int, int, int, int)
	SetCursorMode(CursorMode)
}

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
