// Package engine contains generic implementations of game logic and asset management.
package engine

import (
	"github.com/split-cube-studios/ardent/engine/input"
)

const (
	// FlagResizable indicates that the viewport may be resized.
	FlagResizable = 1 << iota

	// FlagRunsInBackground indicates that logic will continue when the game is unfocused.
	FlagRunsInBackground
)

// Game is an engine instance.
type Game interface {
	// Run starts running the game.
	Run() error

	// AddRenderer adds a renderer
	// to the game's draw stack. Renderers will
	// be applied in the order they are added.
	AddRenderer(...Renderer)

	// IsFullscreen returns the fullscreen state of the game.
	IsFullscreen() bool
	// SetFullscreen sets the fullscreen state of the game.
	SetFullscreen(bool)
	// IsVsync returns the vsync state of the game.
	IsVsync() bool
	// SetVsync sets the vsync state of the game.
	SetVsync(bool)
	// IsFocused returns the focused state of the game.
	IsFocused() bool

	// TODO - Should this be its own interface?
	NewKeySource() input.KeySource
	NewMouseSource() input.MouseSource

	Component
}
