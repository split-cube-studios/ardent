//+build headless

package headless

import "github.com/split-cube-studios/ardent/engine"

// Image is a headless implementation of engine.Image.
type Image struct{}

// Translate sets the image translation.
func (i Image) Translate(x float64, y float64) {
	// NOOP
}

// Offset implements engine.Image.
func (i Image) Offset(x, y float64) {}

// Scale sets the image scale.
func (i Image) Scale(x float64, y float64) {
	// NOOP
}

// Rotate sets the image rotation.
func (i Image) Rotate(d float64) {
	// NOOP
}

// Origin implements engine.Image.
func (i Image) Origin(x, y float64) {}

// SetZDepth implements engine.Image.
func (i Image) SetZDepth(z int) {}

// Tint implements engine.Image.
func (i Image) Tint(r, g, b float64) {}

// SetRenderable implements engine.Image.
func (i Image) SetRenderable(r bool) {}

// IsRenderable implements engine.Image.
func (i Image) IsRenderable() bool {
	return true
}

// Alpha implements engine.Image.
func (i Image) Alpha(alpha float64) {}

// RoundTranslations implements engine.Image.
func (i Image) RoundTranslations(round bool) {}

// TriggersTileOverlapEvent implements engine.Image.
func (i Image) TriggersTileOverlapEvent(triggers bool) {}

// Size returns the image size.
func (i Image) Size() (int, int) {
	return 0, 0
}

// Dispose implements engine.Image.
func (i Image) Dispose() {}

// IsDisposed implements engine.Image.
func (i Image) IsDisposed() bool {
	return false
}

// Position implements engine.Image.
func (i Image) Position() engine.Vec2 {
	return engine.Vec2{}
}

// Class implements engine.Image.
func (i Image) Class() string {
	return "image"
}
