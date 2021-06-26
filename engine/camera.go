package engine

// Camera is a basic implementation of a viewport camera.
type Camera struct {
	vec2 Vec2
}

// LookAt moves the Camera toward the point specified.
func (c *Camera) LookAt(x, y, t float64) {
	c.vec2 = c.vec2.Lerp(Vec2{X: x, Y: y}, t)
}

// Position returns the Camera's current position.
func (c *Camera) Position() (float64, float64) {
	return c.vec2.X, c.vec2.Y
}
