package ebiten

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/common"
)

type ScalableImage struct {
	img     *ebiten.Image
	regions common.ScalableImage
	cache   map[common.ScalableImageRegion]Image

	engine.ImageComponent
}

func (si *ScalableImage) GetImage(w, h int) engine.Image {

	// align corners
	topLeft := si.getRegion(si.regions.TopLeft)
	topRight := si.getRegion(si.regions.TopRight)
	bottomLeft := si.getRegion(si.regions.BottomLeft)
	bottomRight := si.getRegion(si.regions.BottomRight)

	tlw, tlh := topLeft.Size()

	trw, trh := topRight.Size()
	topRight.Translate(float64(w-trw), 0)

	blw, blh := bottomLeft.Size()
	bottomLeft.Translate(0, float64(h-blh))

	brw, brh := bottomRight.Size()
	bottomRight.Translate(float64(w-brw), float64(h-brh))

	// align and stretch horizontal bars
	top := si.getRegion(si.regions.Top)
	bottom := si.getRegion(si.regions.Bottom)

	tw, th := top.Size()
	top.Translate(float64(tlw), 0)
	top.Scale(float64(w-tlw-trw)/float64(tw), 1)

	bw, bh := bottom.Size()
	bottom.Translate(float64(blw), float64(h-bh))
	bottom.Scale(float64(w-blw-brw)/float64(bw), 1)

	// place and stretch vertical bars
	left := si.getRegion(si.regions.Left)
	right := si.getRegion(si.regions.Right)

	lw, lh := left.Size()
	left.Translate(0, float64(tlh))
	left.Scale(1, float64(h-tlh-blh)/float64(lh))

	rw, rh := right.Size()
	right.Translate(float64(w-rw), float64(trh))
	right.Scale(1, float64(h-trh-brh)/float64(rh))

	// place and stretch center
	center := si.getRegion(si.regions.Center)

	cw, ch := center.Size()
	center.Translate(float64(lw), float64(th))
	center.Scale(float64(w-lw-rw)/float64(cw), float64(h-th-bh)/float64(ch))

	// merge layers
	return si.NewImageFromLayers(
		topLeft, top, topRight,
		left, center, right,
		bottomLeft, bottom, bottomRight,
	)
}

func (si *ScalableImage) getRegion(region common.ScalableImageRegion) *Image {

	if img, ok := si.cache[region]; ok {
		return &img
	}

	img := si.img.SubImage(
		image.Rect(
			region.X,
			region.Y,
			region.X+region.W,
			region.Y+region.H,
		),
	)

	cacheImg := Image{
		img:        img.(*ebiten.Image),
		sx:         1,
		sy:         1,
		r:          1,
		g:          1,
		b:          1,
		alpha:      1,
		renderable: true,
	}
	si.cache[region] = cacheImg

	return &cacheImg
}
