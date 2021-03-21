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
}
