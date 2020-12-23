package main

import (
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
		func(ow, oh int) (int, int) {
			return ow, oh
		},
	)

	// create new renderer and image
	renderer := game.NewRenderer()

	assetutil.CreateAssets("./examples/image")
	image, _ := game.NewImageFromAssetPath("./examples/image/scs.asset")

	// add image to renderer
	renderer.AddImage(image)

	// add renderer to game and start game
	game.AddRenderer(renderer)
	game.Run()
}
