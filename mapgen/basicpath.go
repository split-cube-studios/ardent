package mapgen

import (
	"image"
	"math/rand"

	"github.com/split-cube-studios/ardent/engine"
)

type BasicPath struct {
	width               int
	floorTile, wallTile int

	coords, points []image.Point
}

func NewBasicPath(width, floorTile, wallTile int) *BasicPath {
	return &BasicPath{
		width:     width,
		floorTile: floorTile,
		wallTile:  wallTile,
	}
}

func (bp *BasicPath) Flood(tmap *engine.Tilemap) []image.Point {

	var stack []image.Point

	// select random starting position
	for {
		pt := image.Pt(
			rand.Intn(len(tmap.Data[0][0])/bp.width)*bp.width,
			rand.Intn(len(tmap.Data[0])/bp.width)*bp.width,
		)
		// TODO check has valid neighbors?
		if tmap.ContainsAll(bp.wallTile, pt, 1, bp.width) {
			stack = append(stack, pt)
			break
		}
	}

	// build and traverse stack
	for len(stack) != 0 {

		pt := stack[len(stack)-1]

		_ = tmap.Fill(bp.floorTile, pt, 0, bp.width)
		clearedTiles := tmap.Fill(0, pt, 1, bp.width)

		// remove from stack when there are no more neighbors,
		// otherwise add random neighbor to stack.
		ns := bp.validNeighbors(tmap, pt)

		if len(ns) == 0 {
			bp.coords = append(bp.coords, pt)
			bp.points = append(bp.points, clearedTiles...)
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, ns[rand.Intn(len(ns))])
	}

	return bp.points
}

func (bp *BasicPath) PostProcess(tmap *engine.Tilemap) {
	bp.trimDeadEnds(tmap)
	bp.smoothCorners(tmap)
}

func (bp *BasicPath) trimDeadEnds(tmap *engine.Tilemap) {

	found := true
	for found {

		found = false

	path:
		for _, pt := range bp.coords {

			if !tmap.IsClear(pt, 1, bp.width) {
				continue
			}

			var hasNeighbor bool
			for _, n := range tmap.Neighbors(pt, bp.width) {

				if tmap.ContainsAny(0, n, 1, bp.width) {
					if hasNeighbor {
						continue path
					}
					hasNeighbor = true
				}
			}

			tmap.Fill(bp.wallTile, pt, 1, bp.width)
			found = true
		}
	}
}

func (bp *BasicPath) smoothCorners(tmap *engine.Tilemap) {

	var fill, clear []image.Point

	for _, pt := range bp.points {

		var count int
		for _, n := range tmap.Neighbors(pt, 1) {
			if tmap.IsClear(n, 1, 1) {
				count++
			} else {
				var ncount int
				for _, nn := range tmap.Neighbors(n, 1) {
					if tmap.IsClear(nn, 1, 1) {
						ncount++
					}
				}

				if ncount > 1 {
					clear = append(clear, n)
				}
			}
		}

		if count < 3 {
			fill = append(fill, pt)
		}
	}

	for _, pt := range fill {
		tmap.Fill(bp.wallTile, pt, 1, 1)
	}
	for _, pt := range clear {
		tmap.Fill(0, pt, 1, 1)
	}
}

func (bp *BasicPath) validNeighbors(tmap *engine.Tilemap, pt image.Point) []image.Point {

	var validNeighbors []image.Point

	for _, n := range tmap.Neighbors(pt, bp.width) {

		if !tmap.ContainsAll(bp.wallTile, n, 1, bp.width) {
			continue
		}

		var count int
		for _, nn := range tmap.Neighbors(n, bp.width) {
			if tmap.ContainsAll(bp.wallTile, nn, 1, bp.width) {
				count++
			}
		}

		if count == 3 {
			validNeighbors = append(
				validNeighbors, n,
			)
		}
	}

	return validNeighbors
}
