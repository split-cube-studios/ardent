package mapgen

import (
	"errors"
	"fmt"
	"image"
	"math/rand"
	"sort"

	"github.com/split-cube-studios/ardent/engine"
)

// Generator handles tilemap generation based on certain rules.
// Rooms are placed into the world based on their RoomPolicy.
// A perfect maze is generated to fill the space between rooms.
// Rooms are then connected to the maze via their configured doors.
// After all rooms are connected, dead-ends are trimmed back.
// Finally, an optional post-processing phase occurs.
type Generator struct {
	GeneratorOptions

	requiredRooms []Room
	optionalRooms []Room

	// cache of room boundaries for quick lookup
	roomBounds map[image.Rectangle]Room
}

// GeneratorOptions is the configuration
// data for a Generator.
type GeneratorOptions struct {
	Width, Height int
	TileWidth     int

	Rooms []Room

	RoomAlign int

	PathAlg Path

	FloorTile, WallTile int

	Mapper map[int]engine.Image

	OverlapEvent engine.TileOverlapEvent
}

// NewGenerator returns an instantiated *Generator
// with the given GeneratorOptions configuration.
func NewGenerator(options GeneratorOptions) *Generator {

	g := &Generator{
		GeneratorOptions: options,
		roomBounds:       make(map[image.Rectangle]Room),
	}

	for _, room := range options.Rooms {
		if room.Policy().Required {
			g.requiredRooms = append(g.requiredRooms, room)
			continue
		}
		g.optionalRooms = append(g.optionalRooms, room)
	}

	sort.Slice(g.requiredRooms, func(i, j int) bool {
		if g.requiredRooms[i].Policy().Alignment != nil {
			return true
		}
		return false
	})
	sort.Slice(g.optionalRooms, func(i, j int) bool {
		if g.optionalRooms[i].Policy().Alignment != nil {
			return true
		}
		return false
	})

	return g
}

// Generate creates a new tilemap based on the
// generator configuration. An error may be returned
// for failure cases of maze generation, or if a required
// room cannot be placed.
func (g *Generator) Generate() (*engine.Tilemap, error) {

	// place required rooms
	for _, room := range g.requiredRooms {
		if err := g.placeRoom(room); err != nil {
			return nil, fmt.Errorf("failed to generate tilemap: %w", err)
		}
	}

	// place optional rooms
	for _, room := range g.optionalRooms {
		// handle errors if other error cases arise
		_ = g.placeRoom(room)
	}

	// initialize tilemap
	var data [2][][]int
	data[0] = make([][]int, g.Height)
	data[1] = make([][]int, g.Height)
	for y := range data[0] {
		data[0][y] = make([]int, g.Width)
		data[1][y] = make([]int, g.Width)

		// skip fill
		if g.FloorTile == 0 && g.WallTile == 0 {
			continue
		}

		// fill tiles
		for x := 0; x < g.Width; x++ {
			data[0][y][x] = g.FloorTile
			data[1][y][x] = g.WallTile
		}
	}

	tmap := engine.NewTilemap(
		g.TileWidth,
		data,
		g.Mapper,
		g.OverlapEvent,
	)

	// fill tmap with room data
	g.fillRoomData(tmap)

	// fill paths
	if g.PathAlg != nil {

		// TODO retry if path is too short

		// create set of exits for hallways
		exits := make(map[image.Point]struct{})
		for _, pt := range g.PathAlg.Flood(tmap) {
			exits[pt] = struct{}{}
		}

		if err := g.placeHallways(tmap, exits); err != nil {
			return nil, err
		}

		// path post-processing
		g.PathAlg.PostProcess(tmap)
	}

	tmap.BuildCache()

	return tmap, nil
}

// placeRoom places a Room according to the RoomPolicy.
// If a Room cannot be placed an the RoomPolicy
// requires it, an error is returned.
func (g *Generator) placeRoom(room Room) error {

tries:
	for i := 0; i < 250; i++ {

		bounds := room.Bounds()

		var x, y int

		align := room.Policy().Alignment
		if align != nil {
			x = int(float64(g.Width)*align.X) - bounds.Dx()/2
			y = int(float64(g.Height)*align.Y) - bounds.Dy()/2
		} else {
			// snap rooms to alignment
			if g.RoomAlign > 1 {
				x = rand.Intn(
					(g.Width-bounds.Dx())/g.RoomAlign,
				) * g.RoomAlign
				y = rand.Intn(
					(g.Height-bounds.Dy())/g.RoomAlign,
				) * g.RoomAlign
			} else {
				x = rand.Intn(g.Width - bounds.Dx())
				y = rand.Intn(g.Height - bounds.Dy())
			}
		}

		bounds = bounds.Add(image.Pt(
			x, y,
		))

		// check overlaps
		for bounds2, room2 := range g.roomBounds {
			if bounds.Overlaps(bounds2) {
				// allow the overlap
				if room.Policy().CanOverlap && room2.Policy().CanOverlap {
					break
				}
				continue tries
			}
		}

		// record room boundary and return
		g.roomBounds[bounds] = room
		return nil
	}

	// report no errors if not required
	if !room.Policy().Required {
		return nil
	}

	return errors.New("failed to place required room")
}

// fillRoomData iterates over all placed rooms in the generator
// and adds their tile data to a given *engine.Tilemap.
func (g *Generator) fillRoomData(tmap *engine.Tilemap) {

	for bounds, room := range g.roomBounds {
		for z, data := range room.Data() {
			for pt, v := range data {
				tmap.Data[z][pt.Y+bounds.Min.Y][pt.X+bounds.Min.X] = v
			}
		}
	}
}

func (g *Generator) placeHallways(tmap *engine.Tilemap, exits map[image.Point]struct{}) error {

	for bounds, room := range g.roomBounds {

		var placed bool

		path := []image.Point{}
		var entrance, exit [2]map[image.Point]int

		// shouldn't rely on map order for randomness but oh well
		for pt, hallway := range room.Hallways() {

			var v image.Point
			switch hallway.Orientation() {
			case HallwayNorth:
				v = image.Pt(0, -1)
			case HallwaySouth:
				v = image.Pt(0, 1)
			case HallwayEast:
				v = image.Pt(1, 0)
			case HallwayWest:
				v = image.Pt(-1, 0)
			}

			// TODO handle entrance/exit/path size

			// sub to start exactly at pt + bounds.Min
			current := pt.Add(bounds.Min).Sub(v)

			path = path[:0]

			for {
				current = current.Add(v)
				path = append(path, current)

				if !tmap.InBounds(current, 1) {
					break
				}

				if tmap.IsClear(current, 1, 1) {
					if _, ok := exits[current]; !ok {
						break
					}

					placed = true
					entrance, exit = hallway.EntranceData(), hallway.ExitData()

					break
				}
			}

			if placed {
				break
			}
		}

		if !placed {
			return errors.New("failed to place hallway")
		}

		for pt, tile := range entrance[0] {
			pt = pt.Add(path[0])
			tmap.Data[0][pt.Y][pt.X] = tile
		}
		for pt, tile := range entrance[1] {
			pt = pt.Add(path[0])
			tmap.Data[1][pt.Y][pt.X] = tile
		}
		for pt, tile := range exit[0] {
			pt = pt.Add(path[len(path)-1])
			tmap.Data[0][pt.Y][pt.X] = tile
		}
		for pt, tile := range exit[1] {
			pt = pt.Add(path[len(path)-1])
			tmap.Data[1][pt.Y][pt.X] = tile
		}

		for _, pt := range path {
			if !tmap.InBounds(pt, 1) {
				continue
			}

			tmap.Data[1][pt.Y][pt.X] = 0
		}
	}

	return nil
}
