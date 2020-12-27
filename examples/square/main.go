package main

import (
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
	"image"
	"image/color"
)

const (
	w, h  = 854, 480

)

var (
	game engine.Game
	img engine.Image
)

func main() {
	game = ardent.NewGame("Square",
		w, h,
		0,
		func() {
			
		},
		func(w int, h int) (int, int) {
			return w, h
		},
	)

	renderer := game.NewRenderer()
	game.AddRenderer(renderer)

	i := image.NewNRGBA(image.Rect(0, 0, 20, 20))
	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			i.Set(x, y, color.White)
		}
	}

	gameImage := game.NewImageFromImage(i)
	gameImage.Translate(w/2, h/2)
	gameImage.Origin(0.5, 0.5)
	renderer.AddImage(gameImage)

	err := game.Run()
	if err != nil {
		panic(err)
	}
}
