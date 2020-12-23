package common

import "github.com/split-cube-studios/ardent/engine"

// Camera is a basic implementation of engine.Camera.
type Camera struct {
	engine.Vec2
}

// LookAt moves the Camera toward the point specified.
func (c *Camera) LookAt(x, y, t float64) {
	c.Vec2 = c.Vec2.Lerp(engine.Vec2{X: x, Y: y}, t)
}

// Position returns the Camera's current position.
func (c *Camera) Position() (float64, float64) {
	return c.X, c.Y
}
