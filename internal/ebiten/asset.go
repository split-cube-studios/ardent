package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/common"
)

type Asset struct {
	img       Image
	atlas     Atlas
	animation Animation
}

func (a *Asset) ToImage() engine.Image {
	return &a.img
}

func (a *Asset) ToAtlas() engine.Atlas {
	return &a.atlas
}

func (a *Asset) ToAnimation() engine.Animation {
	return &a.animation
}

func (a *Asset) UnmarshalBinary(data []byte) error {
	ca := common.NewAsset()
	if err := ca.UnmarshalBinary(data); err != nil {
		return err
	}

	switch ca.Type {
	case common.AssetTypeImage:
		a.img = Image{
			img:               ebiten.NewImageFromImage(ca.Img),
			sx:                1,
			sy:                1,
			r:                 1,
			g:                 1,
			b:                 1,
			alpha:             1,
			renderable:        true,
			roundTranslations: true,
		}
	case common.AssetTypeAtlas:
		a.atlas = Atlas{
			img:     ebiten.NewImageFromImage(ca.Img),
			regions: ca.AtlasMap,
			cache:   make(map[string]Image),
		}
	case common.AssetTypeAnimation:
		a.animation = Animation{
			Image: Image{
				img:               ebiten.NewImageFromImage(ca.Img),
				sx:                1,
				sy:                1,
				r:                 1,
				g:                 1,
				b:                 1,
				alpha:             1,
				renderable:        true,
				roundTranslations: true,
			},
			w:     ca.AnimWidth,
			h:     ca.AnimHeight,
			anims: ca.AnimationMap,
			cache: make(map[uint16]*ebiten.Image),
		}
	case common.AssetTypeSound:
	default:
		panic("Invalid asset type")
	}

	return nil
}
