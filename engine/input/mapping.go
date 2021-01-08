package input

// Input is an alias for all types of input
type Input = int

// Mapping handles mapping of inputs (ie. KeyE) to actions (ie. Interact).
// These actions are then mapped to performers scoped to a context.
type Mapping struct {
	// The input source to read from
	input Source

	// The bindings that are available for this mapping
	bindings []binding

	// The context stack for this mapping
	contexts []Context

	// Should the mapping read multiple different inputs at once or just the first one found
	allowSimultaneousInput bool
}

type binding struct {
	Input  Input
	Action Action
}

// Create a new mapping using the input source.
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

// Poll the current input state to see if any bound inputs are active.
// If a bound input is active and the bound action can be performed then the action is performed.
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

// Push a context to the top of the stack.
func (m *Mapping) PushContext(ctx Context) {
	m.contexts = append(m.contexts, ctx)
}

// Pop the top context off the stack.
// Does nothing if the stack is empty.
func (m *Mapping) PopContext() {
	n := len(m.contexts) - 1
	if n <= 0 {
		return
	}

	m.contexts = m.contexts[:n]
}
