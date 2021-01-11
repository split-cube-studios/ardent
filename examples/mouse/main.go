package main

import (
	"log"

	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/engine/input"
)

var (
	game    engine.Game
	stripes engine.Image
	mouse   input.MouseSource
)

// tick function.
func tick() {
	cx, cy := mouse.CursorPosition()
	w, h := stripes.Size()
	stripes.Translate(float64(cx-w/2), float64(cy-h/2))
}

func main() {
	// create new game instance
	game = ardent.NewGame(
		"Mouse",
		854,
		480,
		engine.FlagResizable,
		// tick function
		tick,
		// layout function
		func(ow, oh int) (int, int) {
			return ow, oh
		},
	)

	// create new renderer
	renderer := game.NewRenderer()

	// create new atlas from asset file
	atlas, err := game.NewAtlasFromAssetPath("./examples/atlas/atlas.asset")

	if err != nil {
		panic(err)
	}

	// get atlas subimage
	stripes = atlas.GetImage("stripes")

	// add image to renderer
	renderer.AddImage(stripes)

	// add renderer to game and start game
	game.AddRenderer(renderer)
	mouse = game.NewMouseSource()

	err = game.Run()
	if err != nil {
		log.Fatal(err)
	}
}
