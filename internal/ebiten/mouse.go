//+build !headless

package ebiten

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/split-cube-studios/ardent/engine/input"
	"github.com/split-cube-studios/ardent/engine/input/raw"
)

var toEbitenMouseButton = map[int]ebiten.MouseButton{
	raw.MouseButtonLeft:   ebiten.MouseButtonLeft,
	raw.MouseButtonRight:  ebiten.MouseButtonRight,
	raw.MouseButtonMiddle: ebiten.MouseButtonMiddle,
}

var cursorModes = map[raw.CursorMode]ebiten.CursorModeType{
	raw.CursorModeVisible:  ebiten.CursorModeVisible,
	raw.CursorModeHidden:   ebiten.CursorModeHidden,
	raw.CursorModeCaptured: ebiten.CursorModeCaptured,
}

var _ input.MouseSource = (*MouseSource)(nil)

// Input is an engine.Input.
type MouseSource struct {
	minX, minY, maxX, maxY int
	lcx, lcy               int
	vcx, vcy               int
}

func (s *MouseSource) IsAnyPressed() bool {
	for _, v := range toEbitenMouseButton {
		if ebiten.IsMouseButtonPressed(v) {
			return true
		}
	}

	return false
}

func (s *MouseSource) IsAnyJustPressed() bool {
	for _, v := range toEbitenMouseButton {
		if inpututil.IsMouseButtonJustPressed(v) {
			return true
		}
	}

	return false
}

func (s *MouseSource) StateOf(in input.Input) input.State {
	var value float64

	if s.IsPressed(in) {
		value = 1.0
	}

	return input.State{
		Type:        input.MouseButton,
		Input:       in,
		Value:       value,
		JustPressed: s.IsJustPressed(in),
	}
}

func (s *MouseSource) IsPressed(in input.Input) bool {
	return ebiten.IsMouseButtonPressed(toEbitenMouseButton[in])
}

func (s *MouseSource) IsJustPressed(in input.Input) bool {
	return inpututil.IsMouseButtonJustPressed(toEbitenMouseButton[in])
}

func (s *MouseSource) IsJustReleased(in input.Input) bool {
	return inpututil.IsMouseButtonJustReleased(toEbitenMouseButton[in])
}

// CursorPosition implements engine.Input.
func (i *MouseSource) CursorPosition() (int, int) {
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
func (i *MouseSource) SetCursorBounds(minX, minY, maxX, maxY int) {
	i.minX, i.minY, i.maxX, i.maxY = minX, minY, maxX, maxY
}

// SetCursorMode implements engine.Input.
func (i *MouseSource) SetCursorMode(mode raw.CursorMode) {
	ebiten.SetCursorMode(cursorModes[mode])
}
