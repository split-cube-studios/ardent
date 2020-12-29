//+build headless

package headless

import "github.com/split-cube-studios/ardent/engine"

// Atlas is a headless engine.Atlas.
type Atlas struct{}

// GetImage implements engine.Atlas.
func (a Atlas) GetImage(k string) engine.Image {
	return new(Image)
}
