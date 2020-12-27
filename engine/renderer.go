package engine

import "image"

// A Renderer is a basic context for drawing images.
type Renderer interface {
	// AddImage adds one or more images to the renderer's draw stack.
	// Images are drawn in the order they are added.
	AddImage(...Image)

	SetCamera(Camera)

	ScreenToWorld(Vec2) Vec2

	SetViewport(int, int)

	Viewport() image.Rectangle

	// Tick is called by the Game engine each tick. Tick should not be invoked manually
	Tick()
}
