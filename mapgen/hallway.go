package mapgen

import "image"

// Hallway is a path between to clearings,
// along with customizable entrance and exit ends.
type Hallway interface {
	// EntraceData returns the tilemap data
	// for the entrance end of the hallway.
	EntranceData() [2]map[image.Point]int

	// ExitData returns the tilemap data
	// for the exit end of the hallway
	ExitData() [2]map[image.Point]int

	// Width returns the width in tiles
	// of the hallway between the entrance and exit.
	Width() int

	// Orientation returns the orientation of the hallway.
	Orientation() HallwayOrientation
}

type HallwayOrientation byte

const (
	HallwayNorth HallwayOrientation = iota
	HallwaySouth
	HallwayEast
	HallwayWest
)
