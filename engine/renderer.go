package engine

import "image"

// A Renderer is a basic context for drawing images.
type Renderer interface {
	// AddImage adds one or more images to the renderer's draw stack.
	// Images are drawn in the order they are added.
	AddImage(...Image)

	// AddUI adds one or more UIs to the renderer.
	// UIs are rendered in the order they are added.
	// A UI is removed when it is disposed.
	AddUI(...*UI)

	// SetCamera optional sets the renderer's camera.
	SetCamera(*Camera)

	// ScreenToWorld takes a set of screen coordinates
	// and returns the coordinates mapped to the scaled viewport.
	ScreenToWorld(Vec2) Vec2

	// SetViewport sets the renderer's viewport size.
	SetViewport(int, int)

	// Viewport returns an image.Rectangle of the renderer's viewport.
	Viewport() image.Rectangle

	// Tick is called by the Game engine each tick. Tick should not be invoked manually
	Tick()
}
