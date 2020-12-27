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
	camera engine.Camera

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

// SetCamera implements engine.Renderer.
func (r *Renderer) SetCamera(camera engine.Camera) {
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
func (r *Renderer) Tick() {
	/*
		var i int
		for _, img := range r.imgs {
			if img.IsDisposed() {
				continue
			}

			anim, ok := img.(*Animation)
			if ok {
				anim.tick()
			}

			r.imgs[i] = img
			i++
		}

		for j := i; j < len(r.imgs); j++ {
			r.imgs[j] = nil
		}
		r.imgs = r.imgs[:i]
	*/
}

// draw renders all images in the draw stack.
func (r *Renderer) draw(screen *ebiten.Image) {
	var (
		eimg             *ebiten.Image
		tx, ty           float64
		ox, oy           float64
		originX, originY float64
		sx, sy           float64
		d                float64
		cx, cy           float64
		red, green, blue float64
		alpha            float64
	)

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
				img := entry.(engine.Image)

				if !img.IsRenderable() {
					return
				}

				switch a := img.(type) {
				case *Image:
					eimg = a.img
					tx, ty = a.tx+a.ox, a.ty+a.oy
					ox, oy = a.ox, a.oy
					sx, sy = a.sx, a.sy
					originX, originY = a.originX, a.originY
					d = a.d
					red, green, blue = a.r, a.g, a.b
					alpha = a.alpha

				case *Animation:
					a.tick()
					eimg = a.getFrame()
					tx, ty = a.tx+a.ox, a.ty+a.oy
					ox, oy = a.ox, a.oy
					sx, sy = a.sx, a.sy
					originX, originY = a.originX, a.originY
					d = a.d
					red, green, blue = a.r, a.g, a.b
					alpha = a.alpha

				default:
					panic(fmt.Sprintf("Invalid image type %T", img))
				}

				op := new(ebiten.DrawImageOptions)

				w, h := eimg.Size()

				op.GeoM.Scale(sx, sy)
				op.GeoM.Translate(
					-originX*float64(w),
					-originY*float64(h),
				)
				op.GeoM.Rotate(d)
				x, y := tx+ox, ty+oy
				op.GeoM.Translate(
					x-cx,
					y-cy,
				)

				op.ColorM.Scale(red, green, blue, alpha)

				screen.DrawImage(eimg, op)
			}
		})
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
