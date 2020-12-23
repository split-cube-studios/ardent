//+build headless

package headless

import (
	"image"

	"github.com/split-cube-studios/ardent/engine"
)

// Renderer is a headless renderer.
type Renderer struct{}

// AddImage adds images to the draw stack.
func (r Renderer) AddImage(images ...engine.Image) {
	// NOOP
}

// SetCamera implements engine.Renderer.
func (r Renderer) SetCamera(camera engine.Camera) {
}

// ScreenToWorld implements engine.Renderer.
func (r Renderer) ScreenToWorld(screen engine.Vec2) engine.Vec2 {
	return engine.Vec2{}
}

// SetViewport implements engine.Renderer.
func (r Renderer) SetViewport(w, h int) {}

// Tick implements engine.Renderer.
func (r Renderer) Tick() {}

// Viewport implements engine.Renderer.
func (r Renderer) Viewport() image.Rectangle {
	return image.Rect(0, 0, 0, 0)
}
