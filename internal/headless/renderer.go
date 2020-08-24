package headless

import (
	"github.com/split-cube-studios/ardent/engine"
)

// Renderer is a headless renderer.
type Renderer struct{}

// AddImage adds images to the draw stack.
func (r *Renderer) AddImage(images ...engine.Image) {
	// NOOP
}
