//+build !headless

package ebiten

import (
	"fmt"
	"image"
	"math"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/split-cube-studios/ardent/engine"
)

// Renderer is a simple ebiten renderer.
type Renderer struct {
	camera *engine.Camera
	uis    []*engine.UI

	partitionMap *engine.PartitionMap

	w, h int
}

// NewRenderer creates an empty Renderer.
func NewRenderer() *Renderer {
	r := &Renderer{}
	r.partitionMap = engine.NewPartitionMap(1000, 1000)

	return r
}

// AddImage adds images to the draw stack.
func (r *Renderer) AddImage(images ...engine.Image) {
	for _, img := range images {
		img.(disposable).Undispose()
		r.partitionMap.Add(img)
	}
}

// AddUI adds UIs to the renderer.
func (r *Renderer) AddUI(uis ...*engine.UI) {
	r.uis = append(r.uis, uis...)
}

// SetCamera implements engine.Renderer.
func (r *Renderer) SetCamera(camera *engine.Camera) {
	r.camera = camera
}

// ScreenToWorld implements engine.Renderer.
func (r *Renderer) ScreenToWorld(screen engine.Vec2) engine.Vec2 {
	var cx, cy float64

	if r.camera != nil {
		cx, cy = r.camera.Position()
		cx, cy = cx-float64(r.w/2), cy-float64(r.h/2)
	}

	sx := math.Min(
		math.Max(screen.X, 0),
		float64(r.w),
	)
	sy := math.Min(
		math.Max(screen.Y, 0),
		float64(r.h),
	)

	return engine.Vec2{
		X: cx + sx,
		Y: cy + sy,
	}
}

// Tick implements engine.Renderer.
func (r *Renderer) Tick() {}

// draw renders all images in the draw stack.
func (r *Renderer) draw(screen *ebiten.Image) {
	var cx, cy float64

	if r.camera != nil {
		cx, cy = r.camera.Position()
		cx, cy = cx-float64(r.w/2), cy-float64(r.h/2)
	}

	vp := r.Viewport()
	pos := engine.Vec2{
		X: float64(vp.Min.X + (vp.Max.X-vp.Min.X)/2),
		Y: float64(vp.Min.Y + (vp.Max.Y-vp.Min.Y)/2),
	}
	r.partitionMap.Tick(
		pos,
		5,
		func(entries []engine.PartitionEntry) {
			sort.SliceStable(entries, func(i, j int) bool {
				return entries[i].(*Image).z < entries[j].(*Image).z
			})

			for _, entry := range entries {
				op := new(ebiten.DrawImageOptions)
				op.GeoM.Translate(math.Round(-cx), math.Round(-cy))
				r.drawImageAndLayers(entry.(engine.Image), screen, op)
			}
		})

	// draw UIs
	for _, ui := range r.uis {
		for _, img := range ui.Draw() {
			r.drawImageAndLayers(img, screen, nil)
		}
	}
}

func (r *Renderer) drawImageAndLayers(
	img engine.Image,
	screen *ebiten.Image,
	parentOp *ebiten.DrawImageOptions,
) {

	if !img.IsRenderable() {
		return
	}

	eimg, op := engineToEbitenImage(img)

	if parentOp != nil {
		op.GeoM.Concat(parentOp.GeoM)
	}

	screen.DrawImage(eimg, op)

	for _, layer := range img.Layers() {
		r.drawImageAndLayers(layer, screen, op)
	}
}

// SetViewport implements engine.Renderer.
func (r *Renderer) SetViewport(w, h int) {
	r.w, r.h = w, h
}

// Viewport implements engine.Renderer.
func (r *Renderer) Viewport() image.Rectangle {
	var cx, cy float64
	if r.camera != nil {
		cx, cy = r.camera.Position()
	}

	return image.Rect(int(cx), int(cy), r.w+int(cx), r.h+int(cy))
}

func engineImageToLocalImage(img engine.Image) *Image {

	var props *Image

	switch a := img.(type) {
	case *Image:
		props = a
	case *Animation:
		a.tick()
		props = &Image{
			img:                  a.getFrame(),
			tx:                   a.tx,
			ty:                   a.ty,
			ox:                   a.ox,
			oy:                   a.oy,
			sx:                   a.sx,
			sy:                   a.sy,
			originX:              a.originX,
			originY:              a.originY,
			d:                    a.d,
			r:                    a.r,
			g:                    a.g,
			b:                    a.b,
			alpha:                a.alpha,
			renderable:           a.renderable,
			triggersOverlapEvent: a.triggersOverlapEvent,
		}

	default:
		panic(fmt.Sprintf("Invalid image type %T", img))
	}

	return props
}

func engineToEbitenImage(img engine.Image) (*ebiten.Image, *ebiten.DrawImageOptions) {

	props := engineImageToLocalImage(img)

	// typically if an animation frame was not found
	if props.img == nil {
		return nil, nil
	}

	op := new(ebiten.DrawImageOptions)

	w, h := props.img.Size()

	op.GeoM.Scale(props.sx, props.sy)
	op.GeoM.Translate(
		-props.originX*float64(w),
		-props.originY*float64(h),
	)
	op.GeoM.Rotate(props.d)
	x, y := props.tx+props.ox, props.ty+props.oy
	op.GeoM.Translate(x, y)

	op.ColorM.Scale(props.r, props.g, props.b, props.alpha)

	return props.img, op
}
