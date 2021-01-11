//+build headless

package headless

import (
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/engine/input"
)

// Game is a headless implementation of engine.Game.
type Game struct {
	component
	Input

	tickFunc func()
}

// NewGame returns an instantiated game.
func NewGame(
	tickFunc func(),
	layoutFunc func(int, int) (int, int),
) *Game {
	return &Game{
		tickFunc: tickFunc,
	}
}

// Run starts up the engine and begins
// running the game.
func (g *Game) Run() error {
	for {
		g.tickFunc()
	}
}

// AddRenderer adds a renderer to the draw stack.
func (g *Game) AddRenderer(renderer ...engine.Renderer) {
	// NOOP
}

// IsFullscreen returns the fullscreen state of the game.
func (g Game) IsFullscreen() bool {
	return false
}

// SetFullscreen sets the fullscreen state of the game.
func (g Game) SetFullscreen(v bool) {
	// NOOP
}

// IsVsync returns the vsync state of the game.
func (g Game) IsVsync() bool {
	return false
}

// SetVsync sets the vsync state of the game.
func (g Game) SetVsync(v bool) {
	// NOOP
}

// IsFocused returns the focused state of the game.
func (g Game) IsFocused() bool {
	return false
}

func (g Game) NewKeySource() input.KeySource {
	return &KeySource{}
}

func (g Game) NewMouseSource() input.MouseSource {
	return &MouseSource{}
}
