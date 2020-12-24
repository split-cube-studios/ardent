package common

import (
	"math"

	"github.com/split-cube-studios/ardent/engine"
)

// Tilemap is a basic implementation of engine.Tilemap.
type Tilemap struct {
	Width        int
	Data         [2][][]int
	Mapper       map[int]engine.Image
	OverlapEvent engine.TileOverlapEvent
}

// IsoToIndex converts isometric coordinates to a tile index.
func (t *Tilemap) IsoToIndex(x, y float64) (int, int) {
	ix := int(math.Ceil((x/float64(t.Width/2) + y/float64(t.Width/4)) / 2))
	iy := int(math.Ceil((y/float64(t.Width/4) - x/float64(t.Width/2)) / 2))

	return ix + 1, iy + 1
}

// IndexToIso converts a tile index to isometric coordinates.
func (t *Tilemap) IndexToIso(i, j int) (float64, float64) {
	x := (i - j) * (t.Width / 2)
	y := (i + j) * (t.Width / 4)

	return float64(x), float64(y)
}

// GetTileValue returns the value associated with a tile.
func (t *Tilemap) GetTileValue(x, y, z int) int {
	if z < 0 || x < 0 || y < 0 ||
		z >= len(t.Data) ||
		x >= len(t.Data[z]) ||
		y >= len(t.Data[z][y]) {
		return 0
	}

	return t.Data[z][y][x]
}
