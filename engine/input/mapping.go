package input

type Input = int

type input interface {
	IsAnyPressed() bool
	IsAnyJustPressed() bool
	IsPressed(Input) bool
	IsJustPressed(Input) bool
	IsJustReleased(Input) bool
}

type binding struct {
	Input  Input
	Action Action
}

type Mapping struct {
	Type  InputType
	input input

	bindings []binding
	contexts []Context
}

func NewMapping(Type InputType) Mapping {
	m := Mapping{
		Type: Type,
	}

	// TODO - Switch on the inputtype and set the correct input
	return m
}

func (m *Mapping) AddBinding(a Action, defaultBinding Input) {
}

func (m *Mapping) Bind(a Action, binding Input) {

}

func (m *Mapping) BindToNextInput(a Action) {

}
