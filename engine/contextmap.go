package engine

import (
	"image"
	"math"
)

// ContextMap handles context aware
// steering behaviors for in-game AI.
type ContextMap struct {
	arc, dist float64
	cmap, buf []float64

	approachBehavior *ContextMapBehavior
	avoidBehavior    *ContextMapBehavior
	wallBehavior     *ContextMapBehavior

	tmap *Tilemap
}

// ContextMapBehavior contains configurations
// for specific steering behaviors.
type ContextMapBehavior struct {
	// Distance indicates the max distance for this behavior.
	Distance float64
	// Mod is a CMapModFunc to be applied to the calculations.
	Mod CMapModFunc
}

// CMapModFunc is a function type that can be applied
// to values calculated by a ContextMapBehavior.
type CMapModFunc func([]float64)

// NewContextMap returns an instantiated *ContextMap.
// The resolution indicates how many discrete directions will be calculated.
// The tilemap is only required if there is a corresponding wallBehavior.
// All behavior parameters are optional.
func NewContextMap(
	resolution int,
	tmap *Tilemap,
	approachBehavior, avoidBehavior, wallBehavior *ContextMapBehavior,
) *ContextMap {
	return &ContextMap{
		arc:              (math.Pi * 2) / float64(resolution),
		cmap:             make([]float64, resolution),
		buf:              make([]float64, resolution),
		tmap:             tmap,
		approachBehavior: approachBehavior,
		avoidBehavior:    avoidBehavior,
		wallBehavior:     wallBehavior,
	}
}

// Angle returns a selected angle to move based on specified inputs.
func (cm *ContextMap) Angle(origin Vec2, excite, inhibit []Vec2) float64 {

	if cm.approachBehavior != nil {
		for _, v := range excite {
			cm.apply(
				origin, v, false,
				cm.approachBehavior,
			)
		}
	}

	if cm.avoidBehavior != nil {
		for _, v := range inhibit {
			cm.apply(
				origin, v, true,
				cm.avoidBehavior,
			)
		}
	}

	if cm.wallBehavior != nil {
		x, y := cm.tmap.IsoToIndex(origin.X, origin.Y)

		for _, wall := range cm.tmap.WallsAround(image.Pt(x+1, y+1), 1) {
			wx, wy := cm.tmap.IndexToIso(wall.X, wall.Y)
			wy -= float64(cm.tmap.TileWidth - cm.tmap.TileWidth/4)

			cm.apply(
				origin, Vec2{X: wx, Y: wy}, true,
				cm.wallBehavior,
			)
		}
	}

	cm.normalize(cm.cmap)

	var maxIndex int
	for i, m := range cm.cmap {
		if m > cm.cmap[maxIndex] {
			maxIndex = i
		}
	}

	// clear cmap
	for i := range cm.cmap {
		cm.cmap[i] = 0
	}

	return float64(maxIndex) * cm.arc
}

func (cm *ContextMap) apply(u, v Vec2, inhibit bool, behavior *ContextMapBehavior) {

	dist := u.Distance(v)
	m := math.Max(0, math.Min(
		1, (behavior.Distance-dist)/behavior.Distance,
	))

	if m == 0 {
		return
	}

	target := v.Sub(u).Normalize()

	for i := 0; i < len(cm.cmap); i++ {
		delta := (target.Dot(Vec2{
			X: math.Cos(float64(i) * cm.arc),
			Y: math.Sin(float64(i) * cm.arc),
		}) + 1) / 2 * m

		if math.IsNaN(delta) {
			continue
		}

		if inhibit {
			delta = -delta
		}

		cm.buf[i] = delta
	}

	if behavior.Mod != nil {
		behavior.Mod(cm.buf)
	}

	for i, delta := range cm.buf {
		cm.cmap[i] += delta
		cm.buf[i] = 0
	}
}

func (cm *ContextMap) normalize(vx []float64) {
	var max float64
	for _, v := range vx {
		if v > max {
			max = v
		}
	}
	for i, v := range vx {
		vx[i] = v / max
	}
}
