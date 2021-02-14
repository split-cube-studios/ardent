package engine

import "math"

// Tilemap contains tile data for
// IsoRenderer to use.
type Tilemap struct {
	// TileWidth is the tiles width in pixels.
	TileWidth int
	// Data contains values representing tiles.
	Data [2][][]int
	// Mapper maps data values to tile images.
	Mapper map[int]Image
	// OverlapEvent is for each tile, allowing custom
	// overlap behavior (Alpha transitions, events, etc).
	OverlapEvent TileOverlapEvent
}

// NewTilemap returns an instantiated *Tilemap.
// All parameters are required except for overlapEvent.
func NewTilemap(
	tileWidth int,
	data [2][][]int,
	mapper map[int]Image,
	overlapEvent TileOverlapEvent,
) *Tilemap {
	return &Tilemap{
		TileWidth:    tileWidth,
		Data:         data,
		Mapper:       mapper,
		OverlapEvent: overlapEvent,
	}
}

// IsoToIndex converts isometric coordinates to a tile index.
func (t *Tilemap) IsoToIndex(x, y float64) (int, int) {
	ix := int(math.Ceil((x/float64(t.TileWidth/2) + y/float64(t.TileWidth/4)) / 2))
	iy := int(math.Ceil((y/float64(t.TileWidth/4) - x/float64(t.TileWidth/2)) / 2))

	return ix + 1, iy + 1
}

// IndexToIso converts a tile index to isometric coordinates.
func (t *Tilemap) IndexToIso(i, j int) (float64, float64) {
	x := (i - j) * (t.TileWidth / 2)
	y := (i + j) * (t.TileWidth / 4)

	return float64(x), float64(y)
}

// GetTileValue returns the value associated with a tile.
func (t *Tilemap) GetTileValue(x, y, z int) int {
	if z >= len(t.Data) || z < 0 ||
		y >= len(t.Data[z]) || y < 0 ||
		x >= len(t.Data[z][y]) || x < 0 {
		return 0
	}

	return t.Data[z][y][x]
}

// TileOverlapEvent updates renderer state in the case of a tile overlap.
// A bool indicates whether the tile is currently overlapping or not. The tile's
// Image is passed, along with arbitrary data returned from the previous call for state.
type TileOverlapEvent func(bool, Image, interface{}) interface{}
