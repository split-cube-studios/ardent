package main

import (
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
	"image"
	"image/color"
	"math"
)

var (
	w, h = 854, 480
)

var (
	game   engine.Game
	square engine.Image
)

func main() {
	game = ardent.NewGame("Square",
		w, h,
		engine.FlagResizable,
		func() {
			square.Rotate(45.0 * 2 * math.Pi / 360)
			square.Translate(float64(w)/2, float64(h)/2)
		},
		func(nw int, nh int) (int, int) {
			w = nw
			h = nh
			return nw, nh
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

	square = game.NewImageFromImage(i)
	square.Origin(0.5, 0.5)
	square.Rotate(45.0 * 2 * math.Pi / 360)
	square.Translate(float64(w)/2, float64(h)/2)
	renderer.AddImage(square)

	err := game.Run()
	if err != nil {
		panic(err)
	}
}
