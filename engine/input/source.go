package input

import "github.com/split-cube-studios/ardent/engine/input/raw"

type Type int

const (
	Keyboard Type = iota
	MouseButton
	Gamepad
)

// Source is a input source such as a gamepad, keyboard, mouse, and more.
type Source interface {
	IsAnyPressed() bool
	IsAnyJustPressed() bool

	StateOf(Input) State

	IsPressed(Input) bool
	IsJustPressed(Input) bool
	IsJustReleased(Input) bool
}

// KeySource is an input source for keyboards.
type KeySource interface {
	Source
}

// MouseSource is an input source for mice.
type MouseSource interface {
	Source

	CursorPosition() (int, int)
	SetCursorBounds(int, int, int, int)
	SetCursorMode(raw.CursorMode)
}

// State is the current input state of a binding
type State struct {
	// The input source type that this state came from.
	Type Type

	// The raw input that was read.
	Input Input

	// The value [0.0, 1.0] of the input.
	// Binary inputs (i.e. keys) will be either 0.0 or 1.0.
	// Range inputs (i.e. Gamepad sticks) will be a normalized value in the range.
	Value float64

	// This key was pressed this frame
	JustPressed bool
}
