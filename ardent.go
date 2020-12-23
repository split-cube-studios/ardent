package ardent

import "github.com/split-cube-studios/ardent/engine"

// NewGame creates a new game instance.
//
// The backend can be selected with one of the following build tags:
//  ebiten
//  headless
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
