//+build headless

package headless

import "github.com/split-cube-studios/ardent/engine"

// IsoRenderer is a headless engine.IsoRenderer.
type IsoRenderer struct {
	Renderer
}

// SetTilemap implements engine.IsoRenderer.
func (i IsoRenderer) SetTilemap(tilemap *engine.Tilemap) {}
