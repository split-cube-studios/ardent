//+build headless

package headless

import (
	"github.com/split-cube-studios/ardent/engine"
)

// Asset is a headless engine.Asset.
type Asset struct{}

// ToImage implements the ToImage method of engine.Asset.
func (a Asset) ToImage() engine.Image {
	return new(Image)
}

// ToAtlas implements the ToAtlas method of engine.Asset.
func (a Asset) ToAtlas() engine.Atlas {
	return new(Atlas)
}

// ToAnimation implements the ToAnimation method of engine.Asset.
func (a Asset) ToAnimation() engine.Animation {
	return new(Animation)
}

func (a Asset) ToScalableImage() engine.ScalableImage {
	return new(ScalableImage)
}

// ToSound implements the ToSound method of engine.Asset.
func (a Asset) ToSound() engine.Sound {
	return new(Sound)
}
