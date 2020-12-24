// Package ardent is a cross-platform 2D game engine.
package ardent

import "github.com/split-cube-studios/ardent/engine"

// NewGame creates a new game instance.
//
// The backend can be selected with build tags.
// The default, with no tag, is to use Ebiten.
func NewGame(
	title string,
	w, h int,
	flags byte,
	tickFunc func(),
	layoutFunc func(int, int) (int, int),
) engine.Game {
	return newGame(
		title,
		w, h,
		flags,
		tickFunc,
		layoutFunc,
	)
}
