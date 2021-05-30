package mapgen

import (
	"image"

	"github.com/split-cube-studios/ardent/engine"
)

// Path is used by the Generator to fill a tilemap
// with paths around the rooms. In general, a path should
// be a perfect maze, allowing every point to reach every point.
type Path interface {
	// Flood fills the provided *engine.Tilemap
	// with the implemented path algorithm.
	//
	// Flood also returns coordinates of all the placed
	// tiles to be used in generation logic.
	Flood(*engine.Tilemap) []image.Point

	// PostProcess allows a Path to perform edits
	// to the tilemap after rooms and hallways have been placed.
	PostProcess(*engine.Tilemap)
}
