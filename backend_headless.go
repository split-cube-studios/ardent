// +build headless

package ardent

import (
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/headless"
)

func newGame(
	title string,
	w, h int,
	flags byte,
	tickFunc func(),
	layoutFunc engine.LayoutFunc,
) engine.Game {
	return headless.NewGame(
		tickFunc,
		nil,
	)
}
