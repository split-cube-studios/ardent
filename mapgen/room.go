package mapgen

import "image"

// Room is a configurable region of a generated tilemap.
// Rooms can be placed in the world based on the policy
// and maintain full control of their contents.
type Room interface {
	// Policy returns the RoomPolicy.
	Policy() RoomPolicy
	// Data maps two levels of coordinates to tile values.
	Data() [2]map[image.Point]int
	// Bounds returns the bounding box of the Room.
	Bounds() image.Rectangle
	// Hallways returns
	Hallways() map[image.Point]Hallway
}

// RoomPolicy indicates the placement policy of a room.
type RoomPolicy struct {
	// Required indicates if the room is required.
	Required bool
	// CanOverlap indicates if the room may
	// overlap with other rooms.
	CanOverlap bool
	// Alignment indicates a position the room
	// should be placed to the generator.
	Alignment *RoomAlignment
}

// RoomAlignment stores the central alignment
// for a room based on a percent of world size.
type RoomAlignment struct {
	X, Y float64
}

// Default room alignments.
var (
	RoomAlignCenter = &RoomAlignment{
		X: 0.5, Y: 0.5,
	}
	RoomAlignTopLeft = &RoomAlignment{
		X: 0.25, Y: 0.25,
	}
	RoomAlignTopRight = &RoomAlignment{
		X: 0.75, Y: 0.25,
	}
	RoomAlignBottomLeft = &RoomAlignment{
		X: 0.25, Y: 0.75,
	}
	RoomAlignBottomRight = &RoomAlignment{
		X: 0.75, Y: 0.75,
	}
)
