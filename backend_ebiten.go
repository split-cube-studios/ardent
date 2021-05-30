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
	layoutFunc engine.LayoutFunc,
) engine.Game {

	lfunc := layoutFunc
	if lfunc == nil {
		lfunc = engine.LayoutDefault
	}

	return ebiten.NewGame(
		title,
		w,
		h,
		flags,
		tickFunc,
		lfunc,
	)
}
