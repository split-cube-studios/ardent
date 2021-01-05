package input

type Input = int

type Source interface {
	IsAnyPressed() bool
	IsAnyJustPressed() bool

	StateOf(Input) State

	IsPressed(Input) bool
	IsJustPressed(Input) bool
	IsJustReleased(Input) bool
}

type binding struct {
	Input  Input
	Action Action
}

// Mapping handles mapping of inputs (ie. KeyE) to actions (ie. Interact).
// These actions are then mapped to performers scoped to a context.
type Mapping struct {
	input                  Source
	bindings               []binding
	contexts               []Context
	allowSimultaneousInput bool
}

func NewMapping(input Source, allowSimultaneousInput bool) Mapping {
	return Mapping{
		input:                  input,
		bindings:               make([]binding, 0),
		contexts:               make([]Context, 0),
		allowSimultaneousInput: allowSimultaneousInput,
	}
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
			state := m.input.StateOf(b.Input)
			// TODO - could just get the state here and check the value != 0 probably

			if !b.Action.CanPerform(state) {
				return
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

func (m *Mapping) PushContext(ctx Context) {
	m.contexts = append(m.contexts, ctx)
}

func (m *Mapping) PopContext() {
	n := len(m.contexts) - 1
	if n <= 0 {
		return
	}

	m.contexts = m.contexts[:n]
}
