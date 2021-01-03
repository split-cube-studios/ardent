package input

type InputType int

const (
	Keyboard InputType = iota
	MouseButton
	Gamepad
)

type State struct {
	Type  InputType
	Input int
	Value float64
	// TODO - how long has this state been held?
}
