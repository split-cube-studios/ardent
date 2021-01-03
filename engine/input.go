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
