package main

import (
	"log"

	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/assetutil"
	"github.com/split-cube-studios/ardent/engine"
)

func main() {
	// create new game instance
	game := ardent.NewGame(
		"Image",
		854,
		480,
		engine.FlagResizable,
		// tick function
		func() {},
		// layout function
		nil,
	)

	// create new renderer and image
	renderer := game.NewRenderer()

	assetutil.CreateAssets("./examples/image")

	image, _ := game.NewImageFromAssetPath("./examples/image/scs.asset")

	// add image to renderer
	renderer.AddImage(image)

	// add renderer to game and start game
	game.AddRenderer(renderer)

	err := game.Run()
	if err != nil {
		log.Fatal(err)
	}
}
