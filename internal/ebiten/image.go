package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/split-cube-studios/ardent/engine"
)

// Image is an ebiten implementation of engine.Image.
type Image struct {
	img *ebiten.Image

	tx, ty float64
	ox, oy float64
	sx, sy float64
	d      float64

	originX, originY float64

	r, g, b float64
	alpha   float64

	z int

	renderable           bool
	roundTranslations    bool
	triggersOverlapEvent bool
	disposed             bool
}

// Translate sets the image translation.
func (i *Image) Translate(x, y float64) {
	i.tx, i.ty = x, y
}

// Offset applies the translation offset.
func (i *Image) Offset(x, y float64) {
	i.ox, i.oy = x, y
}

// Scale sets the image scale.
func (i *Image) Scale(x, y float64) {
	i.sx, i.sy = x, y
}

// Rotate sets the image rotation.
func (i *Image) Rotate(d float64) {
	i.d = d
}

// Origin sets the image origin by percent.
func (i *Image) Origin(x, y float64) {
	i.originX, i.originY = x, y
}

// SetZDepth sets the z order override.
func (i *Image) SetZDepth(z int) {
	i.z = z
}

// Tint sets the color scale of the image.
func (i *Image) Tint(r, g, b float64) {
	i.r, i.g, i.b = r, g, b
}

// Alpha sets the alpha channel of the image.
func (i *Image) Alpha(alpha float64) {
	i.alpha = alpha
}

// SetRenderable sets the render state of the image.
func (i *Image) SetRenderable(r bool) {
	i.renderable = r
}

// IsRenderable returns the render state of the image.
func (i *Image) IsRenderable() bool {
	return i.renderable
}

// Size returns the image size.
func (i *Image) Size() (int, int) {
	return i.img.Size()
}

// RoundTranslations sets whether or not
// image translations will be rounded during rendering.
func (i *Image) RoundTranslations(round bool) {
	i.roundTranslations = round
}

// TriggersTileOverlapEvent determines whether or not
// the tile overlap event will occur when this image
// is behind a tile in the isometric renderer.
func (i *Image) TriggersTileOverlapEvent(triggers bool) {
	i.triggersOverlapEvent = triggers
}

// Dispose marks the image to be disposed.
func (i *Image) Dispose() {
	i.disposed = true
}

// Undispose resets the disposed state of the image.
func (i *Image) Undispose() {
	i.disposed = false
}

// IsDisposed indicates if the image has been dispoed.
func (i *Image) IsDisposed() bool {
	return i.disposed
}

// Position implements engine.Image.
func (i *Image) Position() engine.Vec2 {
	return engine.Vec2{
		X: i.tx,
		Y: i.ty,
	}
}

// Class implements engine.Image.
func (i *Image) Class() string {
	return "image"
}

// disposable describes behavior for disposable resources.
type disposable interface {
	Dispose()
	Undispose()
	IsDisposed() bool
}
