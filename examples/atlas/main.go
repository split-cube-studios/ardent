package main

import (
	"log"
	"math"

	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/assetutil"
	"github.com/split-cube-studios/ardent/engine"
)

func main() {
	// create new game instance
	game := ardent.NewGame(
		"Atlas",
		854,
		480,
		engine.FlagResizable,
		// tick function
		func() {},
		// layout function
		nil,
	)

	// create new renderer
	renderer := game.NewRenderer()

	// create new atlas from asset file
	assetutil.CreateAssets("./examples/atlas")

	atlas, err := game.NewAtlasFromAssetPath("./examples/atlas/atlas.asset")
	if err != nil {
		log.Fatal(err)
	}

	// get atlas subimages
	stripes := atlas.GetImage("stripes")
	swirls := atlas.GetImage("swirls")
	blocks := atlas.GetImage("blocks")

	// set image positions
	stripes.Rotate(math.Pi / 3)
	swirls.Translate(128, 0)
	blocks.Translate(128, 128)
	blocks.Scale(0.5, 2)

	// add images to renderer
	renderer.AddImage(stripes, swirls, blocks)

	// add renderer to game and start game
	game.AddRenderer(renderer)

	err = game.Run()
	if err != nil {
		log.Fatal(err)
	}
}
