package ebiten

import (
	"image"
	"math"
	"sort"

	"github.com/hajimehoshi/ebiten"
	"github.com/split-cube-studios/ardent/engine"
)

// Renderer is a simple ebiten renderer.
type Renderer struct {
	imgs   []engine.Image
	camera engine.Camera

	partitionMap *engine.PartitionMap

	w, h int
}

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

func (r *Renderer) SetCamera(camera engine.Camera) {
	r.camera = camera
}

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

	return engine.Vec2{cx + sx, cy + sy}
}

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
		float64(vp.Min.X + (vp.Max.X-vp.Min.X)/2),
		float64(vp.Min.Y + (vp.Max.Y-vp.Min.Y)/2),
	}
	r.partitionMap.Tick(
		pos,
		5,
		func(entries []engine.PartitionEntry) {
			sort.SliceStable(entries, func(i, j int) bool {
				z1, z2 := entries[i].(*Image).z, entries[j].(*Image).z
				return z1 < z2
			})

			for _, entry := range entries {
				img := entry.(engine.Image)

				if !img.IsRenderable() {
					return
				}

				switch img.(type) {
				case *Image:
					i := img.(*Image)
					eimg = i.img
					tx, ty = i.tx+i.ox, i.ty+i.oy
					sx, sy = i.sx, i.sy
					originX, originY = i.originX, i.originY
					d = i.d
					red, green, blue = i.r, i.g, i.b
					alpha = i.alpha

				case *Animation:
					a := img.(*Animation)
					a.tick()
					eimg = a.getFrame()
					tx, ty = a.tx+a.ox, a.ty+a.oy
					sx, sy = a.sx, a.sy
					originX, originY = a.originX, a.originY
					d = a.d
					red, green, blue = a.r, a.g, a.b
					alpha = a.alpha

				default:
					panic("Invalid image type")
				}

				op := new(ebiten.DrawImageOptions)

				w, h := eimg.Size()

				op.GeoM.Scale(sx, sy)
				op.GeoM.Translate(tx-cx-originX*float64(w), ty-cy-originY*float64(h))
				op.GeoM.Rotate(d)

				op.ColorM.Scale(red, green, blue, alpha)

				screen.DrawImage(eimg, op)
			}
		})
}

func (r *Renderer) SetViewport(w, h int) {
	r.w, r.h = w, h
}

func (r *Renderer) Viewport() image.Rectangle {
	var cx, cy float64
	if r.camera != nil {
		cx, cy = r.camera.Position()
	}

	return image.Rect(int(cx), int(cy), r.w+int(cx), r.h+int(cy))
}
