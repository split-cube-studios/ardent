//+build headless

package headless

import (
	"github.com/split-cube-studios/ardent/engine"
)

// Asset is a headless engine.Asset.
type Asset struct{}

// ToImage implements engine.Asset.
func (a Asset) ToImage() engine.Image {
	return new(Image)
}

// ToAtlas implements engine.Asset.
func (a Asset) ToAtlas() engine.Atlas {
	return new(Atlas)
}

// ToAnimation implements engine.Asset.
func (a Asset) ToAnimation() engine.Animation {
	return new(Animation)
}
