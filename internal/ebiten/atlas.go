package ebiten

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/common"
)

// Atlas is an engine.Atlas.
type Atlas struct {
	img     *ebiten.Image
	regions map[string]common.AtlasRegion
	cache   map[string]Image
}

// GetImage implements engine.Atlas.
func (a *Atlas) GetImage(k string) engine.Image {
	region, ok := a.regions[k]
	if !ok {
		return nil
	}

	eImg, ok := a.cache[k]
	if ok {
		return &eImg
	}

	img := a.img.SubImage(
		image.Rect(
			int(region.X),
			int(region.Y),
			int(region.X+region.W),
			int(region.Y+region.H),
		),
	)

	cacheImg := Image{
		img:               img.(*ebiten.Image),
		sx:                1,
		sy:                1,
		r:                 1,
		g:                 1,
		b:                 1,
		alpha:             1,
		renderable:        true,
		roundTranslations: true,
	}
	a.cache[k] = cacheImg

	return &cacheImg
}
