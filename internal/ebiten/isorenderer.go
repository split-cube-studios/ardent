package ebiten

import (
	"math"
	"sort"

	"github.com/hajimehoshi/ebiten"
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/common"
)

type IsoRenderer struct {
	Renderer

	drawQueue []*isoRendererImage

	tilemap         *common.Tilemap
	tileEventStates map[[3]int]tileEventState
}

type isoRendererImage struct {
	img        *Image
	isTile     bool
	tileHeight int
	tilePos    [3]int
}

type tileEventState struct {
	complete bool
	state    interface{}
}

func NewIsoRenderer() *IsoRenderer {
	r := &IsoRenderer{
		Renderer: *NewRenderer(),
	}
	r.partitionMap = engine.NewPartitionMap(250, 1000)

	return r
}

func (r *IsoRenderer) SetTilemap(tilemap engine.Tilemap) {
	r.tilemap = tilemap.(*common.Tilemap)
	r.tileEventStates = make(map[[3]int]tileEventState)
}

func (r *IsoRenderer) tilemapToIsoLayers(cx, cy float64) [][]*isoRendererImage {
	if r.tilemap == nil {
		return make([][]*isoRendererImage, 1)
	}

	tw := r.tilemap.Width
	data := r.tilemap.Data
	mapper := r.tilemap.Mapper

	layers := make([][]*isoRendererImage, len(data))

	centerX, centerY := r.tilemap.IsoToIndex(
		cx+float64(r.w/2),
		cy+float64(r.h/2),
	)
	centerX, centerY = int(math.Max(float64(centerX), 0)), int(math.Max(float64(centerY), 0))

	vdim := math.Max(
		float64(r.w),
		float64(r.h),
	) / (float64(tw) * 0.55)

	jmin := int(math.Max(
		math.Max(float64(centerX), 0)-vdim,
		0,
	))
	kmin := int(math.Max(
		math.Max(float64(centerY), 0)-vdim,
		0,
	))

	for i := 0; i < len(data); i++ {
		jmax := int(math.Min(
			float64(len(data[i])),
			vdim+float64(centerX),
		))

		for j := jmin; j < jmax; j++ {
			kmax := int(math.Min(
				float64(len(data[i][j])),
				vdim+float64(centerY),
			))

			for k := kmin; k < kmax; k++ {
				x, y := r.tilemap.IndexToIso(j, k)
				x -= float64(tw / 2)
				y -= float64(tw)

				img := mapper[data[i][k][j]]
				if img == nil {
					continue
				}

				_, h := img.Size()
				if i != 0 {
					y -= float64(h - tw/4)
				}

				layers[i] = append(layers[i], &isoRendererImage{
					img: &Image{
						img:        img.(*Image).img,
						tx:         x,
						ty:         y,
						sx:         1,
						sy:         1,
						r:          1,
						g:          1,
						b:          1,
						alpha:      1,
						renderable: true,
					},
					isTile:     true,
					tileHeight: tw / 2,
					tilePos:    [3]int{i, j, k},
				})
			}
		}
	}

	return layers
}

func (r *IsoRenderer) draw(screen *ebiten.Image) {
	var cx, cy float64
	if r.camera != nil {
		cx, cy = r.camera.Position()
		cx, cy = cx-float64(r.w/2), cy-float64(r.h/2)
	}

	vp := r.Viewport()
	pos := engine.Vec2{
		float64(vp.Min.X),
		float64(vp.Min.Y),
	}

	layers := r.tilemapToIsoLayers(cx, cy)
	topTiles := layers[len(layers)-1]

	r.partitionMap.Tick(
		pos,
		1,
		func(entries []engine.PartitionEntry) {
			for _, entry := range entries {
				img := entry.(engine.Image)
				if !img.IsRenderable() {
					continue
				}

				var tmpImage *isoRendererImage

				// var x, y float64

				switch img.(type) {
				case *Image:
					i := img.(*Image)
					w, h := i.Size()
					tmpImage = &isoRendererImage{
						img: &Image{
							img:                  i.img,
							tx:                   i.tx - i.originX*float64(w),
							ty:                   i.ty - i.originY*float64(h),
							ox:                   i.ox,
							oy:                   i.oy,
							sx:                   i.sx,
							sy:                   i.sy,
							originX:              i.originX,
							originY:              i.originY,
							d:                    i.d,
							z:                    i.z,
							r:                    i.r,
							g:                    i.g,
							b:                    i.b,
							alpha:                i.alpha,
							renderable:           i.renderable,
							roundTranslations:    i.roundTranslations,
							triggersOverlapEvent: i.triggersOverlapEvent,
						},
					}

				//	x, y = i.tx, i.ty

				case *Animation:
					a := img.(*Animation)
					a.tick()
					w, h := a.Size()
					tmpImage = &isoRendererImage{
						img: &Image{
							img:                  a.getFrame(),
							tx:                   a.tx - a.originX*float64(w),
							ty:                   a.ty - a.originY*float64(h),
							ox:                   a.ox,
							oy:                   a.oy,
							sx:                   a.sx,
							sy:                   a.sy,
							originX:              a.originX,
							originY:              a.originY,
							d:                    a.d,
							z:                    a.z,
							r:                    a.r,
							g:                    a.g,
							b:                    a.b,
							alpha:                a.alpha,
							renderable:           a.renderable,
							roundTranslations:    a.roundTranslations,
							triggersOverlapEvent: a.triggersOverlapEvent,
						},
					}

				//	x, y = a.tx, a.ty

				default:
					panic("Invalid image type")
				}

				r.drawQueue = append(r.drawQueue, tmpImage)
			}

			r.drawQueue = append(r.drawQueue, topTiles...)

			// sort queue
			sort.SliceStable(r.drawQueue, func(i, j int) bool {
				var ty1, ty2 float64

				img := r.drawQueue[i].img
				_, h := img.Size()

				if r.drawQueue[i].isTile {
					ty1 = img.ty + float64(h-r.drawQueue[i].tileHeight/8)
				} else {
					ty1 = img.ty + float64(h)
				}

				img = r.drawQueue[j].img
				_, h = img.Size()

				if r.drawQueue[j].isTile {
					ty2 = img.ty + float64(h-r.drawQueue[j].tileHeight/8)
				} else {
					ty2 = img.ty + float64(h)
				}

				return ty1 < ty2
			})

			sort.SliceStable(r.drawQueue, func(i, j int) bool {
				return r.drawQueue[i].img.z < r.drawQueue[j].img.z
			})

			// fix this
			r.drawQueue = append(layers[0], r.drawQueue...)

			for _, isoImage := range r.drawQueue {
				img := isoImage.img

				op := new(ebiten.DrawImageOptions)
				w, h := img.Size()

				op.GeoM.Scale(img.sx, img.sy)
				op.GeoM.Translate(
					-img.originX*float64(w),
					-img.originY*float64(h),
				)
				op.GeoM.Rotate(img.d)
				op.GeoM.Translate(
					img.originX*float64(w),
					img.originY*float64(h),
				)

				x, y := img.tx+img.ox, img.ty+img.oy
				if img.roundTranslations {
					x, y = math.Round(x), math.Round(y)
				}

				op.GeoM.Translate(
					x-cx,
					y-cy,
				)

				op.ColorM.Scale(img.r, img.g, img.b, img.alpha)

				screen.DrawImage(img.img, op)
			}

			r.drawQueue = r.drawQueue[:0]
		},
	)

	/*
		for _, img := range r.imgs {

			less := func(a, b *isoRendererImage) bool {
				var ty1, ty2 float64

				img := a.img
				_, h := img.Size()

				if a.isTile {
					ty1 = img.ty + float64(h-a.tileHeight/8)
				} else {
					ty1 = img.ty + float64(h)
				}

				img = b.img
				_, h = img.Size()

				if b.isTile {
					ty2 = img.ty + float64(h-b.tileHeight/8)
				} else {
					ty2 = img.ty + float64(h)
				}

				return ty1 < ty2
			}

			if r.tilemap.OverlapEvent != nil && tmpImage.img.triggersOverlapEvent {
				for _, tile := range topTiles {
					if !less(tmpImage, tile) {
						continue
					}

					ax := int(tmpImage.img.tx + tmpImage.img.ox)
					ay := int(tmpImage.img.ty + tmpImage.img.oy)
					aw, ah := tmpImage.img.Size()

					bx, by := int(tile.img.tx), int(tile.img.ty)
					bw, bh := tile.img.Size()

					rect1 := image.Rect(ax, ay, ax+aw, ay+ah)
					rect2 := image.Rect(bx, by, bx+bw, by+bh)

					if rect1.Overlaps(rect2) {
						eventState := r.tileEventStates[tile.tilePos]

						if !eventState.complete {
							r.tileEventStates[tile.tilePos] = tileEventState{
								complete: true,
								state:    r.tilemap.OverlapEvent(true, tile.img, eventState.state),
							}
						}
					}
				}
			}

			layers[len(layers)-1] = append(layers[len(layers)-1], tmpImage)
		}

		for i, layer := range layers {

			for _, isoImage := range layer {
				if i == len(layers)-1 && r.tilemap.OverlapEvent != nil && isoImage.isTile {
					event := r.tileEventStates[isoImage.tilePos]
					if !event.complete {
						event.state = r.tilemap.OverlapEvent(false, isoImage.img, event.state)
					}

					event.complete = false
					r.tileEventStates[isoImage.tilePos] = event
				}

				img := isoImage.img

				op := new(ebiten.DrawImageOptions)
				w, h := img.Size()

				op.GeoM.Scale(img.sx, img.sy)
				op.GeoM.Translate(
					-img.originX*float64(w),
					-img.originY*float64(h),
				)
				op.GeoM.Rotate(img.d)
				op.GeoM.Translate(
					img.originX*float64(w),
					img.originY*float64(h),
				)

				x, y := img.tx+img.ox, img.ty+img.oy
				if img.roundTranslations {
					x, y = math.Round(x), math.Round(y)
				}

				op.GeoM.Translate(
					x-cx,
					y-cy,
				)

				op.ColorM.Scale(img.r, img.g, img.b, img.alpha)

				screen.DrawImage(img.img, op)
			}
		}
	*/
}
