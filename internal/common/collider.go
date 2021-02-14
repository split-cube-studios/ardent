package common

import (
	"math"

	"github.com/split-cube-studios/ardent/engine"
)

// Collider is a basic implementation of engine.Collider.
type Collider struct {
	m *engine.Tilemap
}

// SetTilemap sets the Collider's Tilemap.
func (c *Collider) SetTilemap(m *engine.Tilemap) {
	c.m = m
}

// Resolve handles a collision.
func (c *Collider) Resolve(src, dst engine.Vec2) engine.Vec2 {
	if c.m == nil {
		return dst
	}

	ix, iy := c.m.IsoToIndex(dst.X, dst.Y)

	if c.m.GetTileValue(ix, iy, 1) == 0 {
		return dst
	}

	tileX, tileY := c.m.IndexToIso(ix, iy)
	centerX, centerY := tileX, tileY-float64(c.m.TileWidth-c.m.TileWidth/4)

	// tile edge
	tp1 := engine.Vec2{X: centerX - float64(c.m.TileWidth/2), Y: centerY}
	tp2 := engine.Vec2{X: centerX, Y: centerY - float64(c.m.TileWidth/4)}

	var right, bottom bool

	// right corner
	if src.X > centerX {
		tp1.X += float64(c.m.TileWidth)
		right = true
	}

	// bottom corner
	if src.Y > centerY {
		tp2.Y += float64(c.m.TileWidth / 2)
		bottom = true
	}

	nix, niy := ix, iy

	switch {
	case !right && !bottom:
		nix--
	case right && !bottom:
		niy--
	case !right && bottom:
		niy++
	case right && bottom:
		nix++
	}

	// check secondary collision
	if c.m.GetTileValue(nix, niy, 1) != 0 {
		tileX, tileY = c.m.IndexToIso(nix, niy)
		centerX, centerY = tileX, tileY-float64(c.m.TileWidth-c.m.TileWidth/4)

		// tile edge
		tp1 = engine.Vec2{X: centerX - float64(c.m.TileWidth/2), Y: centerY}
		tp2 = engine.Vec2{X: centerX, Y: centerY - float64(c.m.TileWidth/4)}

		right, bottom = false, false

		// right corner
		if src.X > centerX {
			tp1.X += float64(c.m.TileWidth)
			right = true
		}

		// bottom corner
		if src.Y > centerY {
			tp2.Y += float64(c.m.TileWidth / 2)
			bottom = true
		}

		switch {
		case !right && !bottom:
			nix--
		case right && !bottom:
			niy--
		case !right && bottom:
			niy++
		case right && bottom:
			nix++
		}
	}

	// dst right angle to (tp1,tp2)

	atp := engine.Vec2{X: dst.X - tp1.X, Y: dst.Y - tp1.Y}
	atb := engine.Vec2{X: tp2.X - tp1.X, Y: tp2.Y - tp1.Y}

	atb2 := math.Pow(atb.X, 2) + math.Pow(atb.Y, 2)

	atpdotatb := atp.X*atb.X + atp.Y*atb.Y

	t := atpdotatb / atb2

	point := engine.Vec2{
		X: tp1.X + atb.X*t,
		Y: tp1.Y + atb.Y*t,
	}

	// FIXME
	// check tertiary collison
	if c.m.GetTileValue(nix, niy, 1) != 0 {
		tileX, tileY = c.m.IndexToIso(nix, niy)
		centerX, centerY = tileX, tileY-float64(c.m.TileWidth-c.m.TileWidth/4)

		var xMod, yMod float64

		if point.X < centerX-float64(c.m.TileWidth/4) ||
			point.X > centerX+float64(c.m.TileWidth/4) {
			yMod = 1
		}

		if point.Y < centerY-float64(c.m.TileWidth/8) ||
			point.Y > centerY+float64(c.m.TileWidth/8) {
			xMod = 1
		}

		if point.X > centerX {
			xMod *= -1
		}

		if point.Y > centerY {
			yMod *= -1
		}

		point.X += math.Abs(centerX-point.X) * xMod
		point.Y += math.Abs(centerY-point.Y) * yMod
	}

	return point
}
