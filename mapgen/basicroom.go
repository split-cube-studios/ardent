package mapgen

import "image"

// BasicRoom is a simple rectangular implementation
// of a Room. The floor tiles are filled with a single tile,
// while the level above is left empty.
type BasicRoom struct {
	w, h   int
	data   [2]map[image.Point]int
	policy RoomPolicy
}

// NewBasicRoom returns an instantiated *BasicRoom with a specified
// size, tile, and policy.
func NewBasicRoom(w, h, floorTile, wallTile int, policy RoomPolicy) *BasicRoom {

	br := &BasicRoom{
		w:      w,
		h:      h,
		policy: policy,
	}

	br.data[0] = make(map[image.Point]int)
	br.data[1] = make(map[image.Point]int)

	// populate data
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			pt := image.Pt(x, y)
			br.data[0][pt] = floorTile
			br.data[1][pt] = 0

			br.data[1][image.Pt(0, y)] = wallTile
			br.data[1][image.Pt(w-1, y)] = wallTile
		}

		br.data[1][image.Pt(x, 0)] = wallTile
		br.data[1][image.Pt(x, h-1)] = wallTile
	}

	return br
}

// Policy implements the Policy method of Room.
func (br *BasicRoom) Policy() RoomPolicy {
	return br.policy
}

// Data implements the Data method of Room.
func (br *BasicRoom) Data() [2]map[image.Point]int {
	return br.data
}

// Bounds implements the Bounds method of Room.
func (br *BasicRoom) Bounds() image.Rectangle {
	return image.Rect(0, 0, br.w, br.h)
}

// Hallways returns all possible hallways for the Room.
func (br *BasicRoom) Hallways() map[image.Point]Hallway {

	hallways := make(map[image.Point]Hallway)

	const doorWidth, hallWidth = 3, 1

	for x := 1; x < br.w-1; x++ {
		hNorth := NewBasicHallway(doorWidth, hallWidth, 1, HallwayNorth)
		hallways[image.Pt(x, 0)] = hNorth

		hSouth := NewBasicHallway(doorWidth, hallWidth, 1, HallwaySouth)
		hallways[image.Pt(x, br.h-1)] = hSouth
	}
	for y := 1; y < br.h-1; y++ {
		hWest := NewBasicHallway(doorWidth, hallWidth, 1, HallwayWest)
		hallways[image.Pt(0, y)] = hWest

		hEast := NewBasicHallway(doorWidth, hallWidth, 1, HallwayEast)
		hallways[image.Pt(br.w-1, y)] = hEast
	}

	return hallways
}
