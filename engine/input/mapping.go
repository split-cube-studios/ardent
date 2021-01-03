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

	// Should the mapper perform actions for all current inputs or just the first input
	allowSimultaneousInput bool
}

func NewMapping(Type InputType) Mapping {
	m := Mapping{
		Type: Type,
	}

	// TODO - Switch on the inputtype and set the correct input
	return m
}

// Bind adds a new binding to the mapper if it does not exist, otherwise updates the binding.
func (m *Mapping) Bind(a Action, input Input) {
	for _, binding := range m.bindings {
		if binding.Action == a {
			binding.Input = input
			return
		}
	}

	m.bindings = append(m.bindings, binding{
		Input:  input,
		Action: a,
	})
}

func (m *Mapping) Poll() {
	for _, b := range m.bindings {
		if m.input.IsPressed(b.Input) {
			state := State{
				Type:  m.Type,
				Input: b.Input,
				Value: 1.0,
			}
			m.perform(b.Action, state)

			if !m.allowSimultaneousInput {
				return
			}
		}
	}
}

func (m *Mapping) perform(a Action, s State) {
	for _, c := range m.contexts {
		consumedInput := c.perform(a, s)

		if consumedInput {
			return
		}
	}
}
