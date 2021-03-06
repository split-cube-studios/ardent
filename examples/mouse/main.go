package main

import (
	"log"

	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

var (
	game    engine.Game
	stripes engine.Image
)

// tick function.
func tick() {
	cx, cy := game.CursorPosition()
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
		nil,
	)

	// create new renderer
	renderer := game.NewRenderer()

	// create new atlas from asset file
	atlas, _ := game.NewAtlasFromAssetPath("./examples/atlas/atlas.asset")

	// get atlas subimage
	stripes = atlas.GetImage("stripes")

	// add image to renderer
	renderer.AddImage(stripes)

	// add renderer to game and start game
	game.AddRenderer(renderer)

	err := game.Run()
	if err != nil {
		log.Fatal(err)
	}
}
