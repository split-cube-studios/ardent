package main

import (
	"log"
	"math"

	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/aautil"
	"github.com/split-cube-studios/ardent/engine"
)

func main() {

	var (
		layerImage engine.Image
		angle      float64
	)

	// create new game instance
	game := ardent.NewGame(
		"Atlas",
		854,
		480,
		engine.FlagResizable,
		// tick function
		func() {
			layerImage.Rotate(angle)
			angle += 0.01
		},
		// layout function
		nil,
	)

	// create new renderer
	renderer := game.NewRenderer()

	// create new atlas from asset file
	aautil.CreateAssets("./examples/atlas")

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
	stripes.Translate(854/2, 480/2)

	// stripes will be the base layer, so everything is relative to its properties
	swirls.Translate(128, 128)
	swirls.Rotate(math.Pi / 5)
	blocks.Translate(128, 128)
	blocks.Scale(0.5, 2)

	layerImage = game.NewImageFromLayers(stripes, swirls, blocks)

	// add images to renderer
	renderer.AddImage(layerImage)

	// add renderer to game and start game
	game.AddRenderer(renderer)

	err = game.Run()
	if err != nil {
		log.Fatal(err)
	}
}
