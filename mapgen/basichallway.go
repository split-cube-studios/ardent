package mapgen

import "image"

type BasicHallway struct {
	doorWidth, hallWidth, tile int
	entrance, exit             [2]map[image.Point]int
	orientation                HallwayOrientation
}

func NewBasicHallway(
	doorWidth, hallWidth, tile int,
	orientation HallwayOrientation,
) *BasicHallway {

	bh := &BasicHallway{
		doorWidth:   doorWidth,
		hallWidth:   hallWidth,
		tile:        tile,
		orientation: orientation,
	}

	bh.entrance[0] = make(map[image.Point]int)
	bh.entrance[1] = make(map[image.Point]int)
	bh.exit[0] = make(map[image.Point]int)
	bh.exit[1] = make(map[image.Point]int)

	for x := -doorWidth / 2; x <= doorWidth/2; x++ {
		for y := -doorWidth / 2; y <= doorWidth/2; y++ {
			bh.entrance[0][image.Pt(x, y)] = tile
			bh.entrance[1][image.Pt(x, y)] = 0

			bh.exit[0][image.Pt(x, y)] = tile
			bh.exit[1][image.Pt(x, y)] = 0
		}
	}

	return bh
}

func (bh *BasicHallway) EntranceData() [2]map[image.Point]int {
	return bh.entrance
}

func (bh *BasicHallway) ExitData() [2]map[image.Point]int {
	return bh.exit
}

func (bh *BasicHallway) Width() int {
	return bh.hallWidth
}

func (bh *BasicHallway) Orientation() HallwayOrientation {
	return bh.orientation
}
