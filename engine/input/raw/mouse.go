package raw

const (
	MouseButtonLeft int = iota
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
