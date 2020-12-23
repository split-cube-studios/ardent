//+build !headless

package ebiten

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/split-cube-studios/ardent/engine"
)

// Input is an engine.Input.
type Input struct {
	minX, minY, maxX, maxY int
	lcx, lcy               int
	vcx, vcy               int
}

// IsAnyKeyPressed implements engine.Input.
func (i *Input) IsAnyKeyPressed() bool {
	for _, v := range toEbitenKey {
		if ebiten.IsKeyPressed(v) {
			return true
		}
	}

	return false
}

// IsAnyKeyJustPressed implements engine.Input.
func (i *Input) IsAnyKeyJustPressed() bool {
	for _, v := range toEbitenKey {
		if inpututil.IsKeyJustPressed(v) {
			return true
		}
	}

	return false
}

// IsKeyPressed implements engine.Input.
func (i *Input) IsKeyPressed(k int) bool {
	return ebiten.IsKeyPressed(toEbitenKey[k])
}

// IsKeyJustPressed implements engine.Input.
func (i *Input) IsKeyJustPressed(k int) bool {
	return inpututil.IsKeyJustPressed(toEbitenKey[k])
}

// IsKeyJustReleased implements engine.Input.
func (i *Input) IsKeyJustReleased(k int) bool {
	return inpututil.IsKeyJustReleased(toEbitenKey[k])
}

// IsMouseButtonPressed implements engine.Input.
func (i *Input) IsMouseButtonPressed(k int) bool {
	return ebiten.IsMouseButtonPressed(toEbitenMouseButton[k])
}

// IsMouseButtonJustPressed implements engine.Input.
func (i *Input) IsMouseButtonJustPressed(k int) bool {
	return inpututil.IsMouseButtonJustPressed(toEbitenMouseButton[k])
}

// IsMouseButtonJustReleased implements engine.Input.
func (i *Input) IsMouseButtonJustReleased(k int) bool {
	return inpututil.IsMouseButtonJustReleased(toEbitenMouseButton[k])
}

// CursorPosition implements engine.Input.
func (i *Input) CursorPosition() (int, int) {
	x, y := ebiten.CursorPosition()

	if x <= math.MinInt32 {
		x = 0
	}

	if y <= math.MinInt32 {
		y = 0
	}

	if i.minX+i.minY+i.maxX+i.maxY == 0 {
		return x, y
	}

	dx, dy := x-i.lcx, y-i.lcy
	i.lcx, i.lcy = x, y

	nx, ny := i.vcx+dx, i.vcy+dy

	switch {
	case nx < i.minX:
		i.vcx = i.minX
	case nx > i.maxX:
		i.vcx = i.maxX
	default:
		i.vcx = nx
	}

	switch {
	case ny < i.minY:
		i.vcy = i.minY
	case ny > i.maxY:
		i.vcy = i.maxY
	default:
		i.vcy = ny
	}

	return i.vcx, i.vcy
}

// SetCursorBounds implements engine.Input.
func (i *Input) SetCursorBounds(minX, minY, maxX, maxY int) {
	i.minX, i.minY, i.maxX, i.maxY = minX, minY, maxX, maxY
}

var cursorModes = map[engine.CursorMode]ebiten.CursorModeType{
	engine.CursorModeVisible:  ebiten.CursorModeVisible,
	engine.CursorModeHidden:   ebiten.CursorModeHidden,
	engine.CursorModeCaptured: ebiten.CursorModeCaptured,
}

// SetCursorMode implements engine.Input.
func (i *Input) SetCursorMode(mode engine.CursorMode) {
	ebiten.SetCursorMode(cursorModes[mode])
}
