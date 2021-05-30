package main

import (
	"image/color"
	"log"

	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
	"golang.org/x/image/font/basicfont"
)

func main() {
	// create new game instance
	game := ardent.NewGame(
		"Text",
		854,
		480,
		engine.FlagResizable,
		// tick function
		func() {},
		// layout function
		engine.LayoutFit(854, 450),
	)

	// create new renderer and text image
	renderer := game.NewRenderer()
	image := game.NewTextImage(
		"Hello world!\nThis is a sample text image!",
		400,
		30,
		basicfont.Face7x13,
		color.White,
	)

	image.Scale(2, 2)

	// add image to renderer
	renderer.AddImage(image)

	// add renderer to game and start game
	game.AddRenderer(renderer)

	err := game.Run()
	if err != nil {
		log.Fatal(err)
	}
}
