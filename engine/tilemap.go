package engine

import (
	"fmt"
	"image"
	"math"
)

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

	bounds image.Rectangle
}

// TileOverlapEvent updates renderer state in the case of a tile overlap.
// A bool indicates whether the tile is currently overlapping or not. The tile's
// Image is passed, along with arbitrary data returned from the previous call for state.
type TileOverlapEvent func(bool, Image, interface{}) interface{}

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
		bounds:       image.Rect(0, 0, len(data[0][0]), len(data[0])),
	}
}

// IsoToIndex converts isometric coordinates to a tile index.
func (t *Tilemap) IsoToIndex(x, y float64) (int, int) {
	ix := int(math.Ceil((x/float64(t.TileWidth/2) + y/float64(t.TileWidth/4)) / 2))
	iy := int(math.Ceil((y/float64(t.TileWidth/4) - x/float64(t.TileWidth/2)) / 2))

	return ix, iy
}

// IndexToIso converts a tile index to isometric coordinates.
func (t *Tilemap) IndexToIso(i, j int) (float64, float64) {
	x := (i - j) * (t.TileWidth / 2)
	y := (i + j) * (t.TileWidth / 4)

	return float64(x), float64(y)
}

// GetTileValue returns the value associated with a tile.
func (t *Tilemap) GetTileValue(x, y, z int) int {

	if z < 0 || z > 1 || !t.InBounds(image.Pt(x, y), 1) {
		return 0
	}

	return t.Data[z][y][x]
}

var ndirs = [4]image.Point{
	image.Pt(1, 0),
	image.Pt(0, 1),
	image.Pt(-1, 0),
	image.Pt(0, -1),
}

// Neighbors returns points adjacent to point p
// that are fully within the tilemap.
// The size parameter can be used as a virtual tile scale.
// A size less than 1 causes a panic.
func (t *Tilemap) Neighbors(p image.Point, size int) (c []image.Point) {

	if size < 1 {
		panic("invalid size")
	}

	for i := 0; i < 4; i++ {

		np := p.Add(ndirs[i].Mul(size))
		fmt.Println(np)
		if !t.InBounds(np, size) {
			continue
		}

		c = append(c, np)
	}

	return
}

// InBounds indicates if a point with a given size is within the tilemap.
func (t *Tilemap) InBounds(p image.Point, size int) bool {

	if size < 1 {
		panic("invalid size")
	}

	return p.In(t.bounds) && p.Add(image.Pt(size-1, size-1)).In(t.bounds)
}

// IsClear indicates whether a point of a given size
// contains only empty tiles.
func (t *Tilemap) IsClear(p image.Point, z, size int) bool {

	if size < 1 {
		panic("invalid size")
	}

	if !t.InBounds(p, size) {
		return false
	}

	for x := p.X; x < p.X+size; x++ {
		for y := p.Y; y < p.Y+size; y++ {
			if t.Data[z][y][x] != 0 {
				return false
			}
		}
	}

	return true
}
