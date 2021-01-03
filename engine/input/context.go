package input

type Context struct {
	actions map[Action][]ActionPerformer
}

func (c *Context) perform(a Action, s State) bool {
	performers, ok := c.actions[a]

	if !ok {
		// This context doesnt have this action
		return false
	}

	var shouldConsumeInput bool
	for _, p := range performers {
		consumedInput := p.Perform(s)

		if consumedInput {
			shouldConsumeInput = true
		}
	}

	return shouldConsumeInput
}
