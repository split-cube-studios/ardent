package main

import (
	"log"

	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/assetutil"
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/engine/input"
	"github.com/split-cube-studios/ardent/engine/input/raw"
)

var (
	game                      engine.Game
	camera                    engine.Camera
	stripes                   engine.Image
	keyboard                  input.KeySource
	x, y                      float64
	stripeWidth, stripeHeight int
)

const (
	w, h = 854, 480
)

// tick function.
func tick() {
	if keyboard.IsPressed(raw.KeyW) {
		y -= 2
	} else if keyboard.IsPressed(raw.KeyS) {
		y += 2
	}

	if keyboard.IsPressed(raw.KeyA) {
		x -= 2
	} else if keyboard.IsPressed(raw.KeyD) {
		x += 2
	}

	stripes.Translate(x, y)

	// 0.05 lerp rate
	camera.LookAt(x+float64(stripeWidth/2), y+float64(stripeHeight/2), 0.05)
}

func main() {
	// create new game instance
	game = ardent.NewGame(
		"Camera",
		w,
		h,
		engine.FlagResizable,
		// tick function
		tick,
		// layout function
		func(ow, oh int) (int, int) {
			// preserve virtual res
			return w, h
		},
	)

	// create new renderer
	renderer := game.NewRenderer()

	// create new camera
	camera = game.NewCamera()

	// create new atlas from asset file
	assetutil.CreateAssets("./examples/atlas")

	atlas, _ := game.NewAtlasFromAssetPath("./examples/atlas/atlas.asset")

	// get atlas subimages
	stripes = atlas.GetImage("stripes")
	swirls := atlas.GetImage("swirls")
	blocks := atlas.GetImage("blocks")

	stripeWidth, stripeHeight = stripes.Size()

	// set image positions
	swirls.Translate(128, 0)
	blocks.Translate(128, 128)

	// add images to renderer
	renderer.AddImage(swirls, blocks, stripes)

	// add camera to renderer
	renderer.SetCamera(camera)

	// add renderer to game and start game
	game.AddRenderer(renderer)
	keyboard = game.NewKeySource()
	err := game.Run()
	if err != nil {
		log.Fatal(err)
	}
}
