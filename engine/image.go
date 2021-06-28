package engine

// An Image represents a single unchanging
// image that can be applied to a renderer.
type Image interface {
	// Translate sets the x y translation of the image relative to the origin.
	Translate(float64, float64)

	// Offset applies an offset to the image translation.
	// This can be useful for relative positioning to a parent translation.
	Offset(float64, float64)

	// Scale sets the x y scale of the image relative to the origin.
	Scale(float64, float64)

	// Rotate sets the rotation in radians relative to the origin.
	Rotate(float64)

	// Origin sets the coordinate origin of the image in percent ranging from 0.0 to 1.0
	Origin(float64, float64)

	// SetZDepth sets a z value to override draw order.
	SetZDepth(int)

	// Tint scales the image colors by a factor of each value.
	Tint(float64, float64, float64)

	// Alpha sets the image's alpha channel with a range of 0.0 to 1.0
	Alpha(float64)

	// Layers returns all layers of an image, excluding the root layer.
	Layers() []Image

	// SetRenderable sets whether or not an image should be rendered.
	SetRenderable(bool)

	// IsRenderable indicates whether or not an image should be rendered.
	IsRenderable() bool

	// TriggersTileOverlapEvent determines whether tile overlap events will occur.
	// A tile overlap is when an image is behind a tile in the isometric renderer.
	TriggersTileOverlapEvent(bool)

	// Size returns the size of the image.
	Size() (int, int)

	// Dispose marks the image to be disposed.
	Dispose()

	// IsDisposed indicates if the image has been disposed.
	IsDisposed() bool

	// Position returns the translated position of the image.
	Position() Vec2

	// Class returns the image class.
	Class() string
}
