package input

// Context is a scoped map of Actions to Action Performers.
type Context struct {
	actions map[Action][]ActionPerformer
}

// Adds a performer to be called when the action is performed.
func (c *Context) AddPerformer(a Action, p ActionPerformer) {
	performers, ok := c.actions[a]
	if !ok {
		performers = make([]ActionPerformer, 0)
	}

	c.actions[a] = append(performers, p)
}

// Removes a performer from the action's performers
func (c *Context) RemovePerformer(a Action, p ActionPerformer) {
	performers, ok := c.actions[a]
	if !ok {
		return
	}

	remaining := make([]ActionPerformer, 0)
	for _, performer := range performers {
		if performer != p {
			remaining = append(remaining, performer)
		}
	}
	c.actions[a] = performers
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
