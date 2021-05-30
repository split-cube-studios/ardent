package mapgen

import (
	"image"
	"math"
	"math/rand"
)

type NaturalRoom struct {
	w, h      int
	data      [2]map[image.Point]int
	policy    RoomPolicy
	floorTile int
}

func NewNaturalRoom(w, h, floorTile int, policy RoomPolicy) *NaturalRoom {

	nr := &NaturalRoom{
		w:         w,
		h:         h,
		floorTile: floorTile,
		policy:    policy,
	}

	nr.data[0] = make(map[image.Point]int)
	nr.data[1] = make(map[image.Point]int)

	// tilebomb room
	var points, tiles []image.Point

	rx, ry := w/2, h/2

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {

			ex := float64(x-rx) * float64(x-rx) / float64(rx*rx)
			ey := float64(y-ry) * float64(y-ry) / float64(ry*ry)

			if ex+ey <= 1 {

				// randomly skip filling
				if rand.Intn(3) == 0 {
					continue
				}

				points = append(tiles, image.Pt(x, y))
				tiles = append(tiles, image.Pt(x, y))
			}
		}
	}

	bounds := image.Rect(
		tiles[0].X, tiles[0].Y,
		tiles[len(tiles)-1].X, tiles[len(tiles)-1].Y,
	)

	rand.Shuffle(len(tiles), func(i, j int) {
		tiles[i], tiles[j] = tiles[j], tiles[i]
	})

	iters := len(tiles) * 7

	smOffsets := []image.Point{
		image.Pt(-1, 0),
		image.Pt(1, 0),
		image.Pt(0, 1),
		image.Pt(0, -1),
	}

	lgOffsets := []image.Point{
		image.Pt(-2, 0),
		image.Pt(2, 0),
		image.Pt(0, 2),
		image.Pt(0, -2),
		image.Pt(1, 1),
		image.Pt(1, -1),
		image.Pt(-1, 1),
		image.Pt(-1, -1),
	}

	for n := 0; n < iters; n++ {

		var i int
		if rand.Intn(3) == 0 {
			o := int(math.Min(float64(len(tiles)), 15))
			i = rand.Intn(o) + len(tiles) - o
		} else if len(tiles) == 1 {
			i = 0
		} else {
			i = rand.Intn(len(tiles) / 2)
		}

		offsets := smOffsets
		if rand.Intn(20) == 0 {
			offsets = append(offsets, lgOffsets...)
		}

		for _, offset := range offsets {

			pt := tiles[i].Add(offset)

			if pt.X < bounds.Min.X {
				bounds.Min.X = pt.X
			} else if pt.X > bounds.Max.X {
				bounds.Max.X = pt.X
			}
			if pt.Y < bounds.Min.Y {
				bounds.Min.Y = pt.Y
			} else if pt.Y > bounds.Max.Y {
				bounds.Max.Y = pt.Y
			}

			points = append(points, pt)
			tiles = append(tiles, pt)
		}

		if len(tiles) == 1 {
			break
		}

		tiles[i] = tiles[len(tiles)-1]
		tiles = tiles[:len(tiles)-1]
	}

	for _, pt := range points {

		pt = pt.Sub(bounds.Min).Add(image.Pt(1, 1))

		nr.data[0][pt] = floorTile
		nr.data[1][pt] = 0
	}

	nr.w, nr.h = bounds.Dx()+2, bounds.Dy()+2

	return nr
}

func (nr *NaturalRoom) Policy() RoomPolicy {
	return nr.policy
}

func (nr *NaturalRoom) Data() [2]map[image.Point]int {
	return nr.data
}

func (nr *NaturalRoom) Bounds() image.Rectangle {
	return image.Rect(0, 0, nr.w, nr.h)
}

func (nr *NaturalRoom) Hallways() map[image.Point]Hallway {

	hallways := make(map[image.Point]Hallway)

	const doorWidth, hallWidth = 3, 1

	for x := 1; x < nr.w-1; x++ {
		hNorth := NewBasicHallway(doorWidth, hallWidth, nr.floorTile, HallwayNorth)
		hallways[image.Pt(x, 0)] = hNorth

		hSouth := NewBasicHallway(doorWidth, hallWidth, nr.floorTile, HallwaySouth)
		hallways[image.Pt(x, nr.h-1)] = hSouth
	}
	for y := 1; y < nr.h-1; y++ {
		hWest := NewBasicHallway(doorWidth, hallWidth, nr.floorTile, HallwayWest)
		hallways[image.Pt(0, y)] = hWest

		hEast := NewBasicHallway(doorWidth, hallWidth, nr.floorTile, HallwayEast)
		hallways[image.Pt(nr.w-1, y)] = hEast
	}

	return hallways
}
