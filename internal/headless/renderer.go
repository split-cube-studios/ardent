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

func (r Renderer) SetCamera(camera engine.Camera) {

}

func (r Renderer) ScreenToWorld(screen engine.Vec2) engine.Vec2 {
	return engine.Vec2{}
}

func (r Renderer) SetViewport(w, h int) {}

func (r Renderer) Tick() {}

func (r Renderer) Viewport() image.Rectangle {
	return image.Rect(0, 0, 0, 0)
}
