package headless

// Animation is a headless engine.Animation.
type Animation struct {
	Image
}

// SetState implements engine.Animation.
func (a Animation) SetState(state string) {}

// SetTickCount implements engine.Animation.
func (a Animation) SetTickCount(count int) {}

// Play implements engine.Animation.
func (a Animation) Play() {}

// Pause implements engine.Animation.
func (a Animation) Pause() {}

// Reset implements engine.Animation.
func (a Animation) Reset() {}
