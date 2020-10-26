package ebiten

import "github.com/hajimehoshi/ebiten"

// Image is an ebiten implementation of engine.Image
type Image struct {
	img *ebiten.Image

	tx, ty float64
	ox, oy float64
	sx, sy float64
	d      float64

	originX, originY float64

	z int

	disposed bool
}

// Translate sets the image translation.
func (i *Image) Translate(x, y float64) {
	i.tx, i.ty = x, y
}

// Offset applies the translation offset
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

// Size returns the image size.
func (i *Image) Size() (int, int) {
	return i.img.Size()
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

// disposable describes behavior for disposable resources.
type disposable interface {
	Dispose()
	Undispose()
	IsDisposed() bool
}
