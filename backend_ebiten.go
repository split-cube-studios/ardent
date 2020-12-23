// +build !headless

package ardent

import (
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/ebiten"
)

func newGame(
	title string,
	w, h int,
	flags byte,
	tickFunc func(),
	layoutFunc func(int, int) (int, int),
) engine.Game {
	return ebiten.NewGame(
		title,
		w,
		h,
		flags,
		tickFunc,
		layoutFunc,
	)
}
