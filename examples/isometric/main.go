package main

import (
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/assetutil"
	"github.com/split-cube-studios/ardent/engine"
)

const (
	w, h  = 854, 480
	speed = 2
)

var (
	game      engine.Game
	animation engine.Animation
	x, y      float64
)

func main() {
	// create new game instance
	game = ardent.NewGame(
		"Isometric",
		w,
		h,
		engine.FlagResizable,
		// tick function
		func() {
			if game.IsKeyPressed(engine.KeyW) {
				y -= speed
			} else if game.IsKeyPressed(engine.KeyS) {
				y += speed
			}

			if game.IsKeyPressed(engine.KeyA) {
				x -= speed
			} else if game.IsKeyPressed(engine.KeyD) {
				x += speed
			}

			animation.Translate(x, y)
		},
		// layout function
		func(ow, oh int) (int, int) {
			return w, h
		},
	)

	assetutil.CreateAssets("./examples/isometric")
	atlas, _ := game.NewAtlasFromAssetPath("./examples/isometric/tiles.asset")

	data := [2][][]int{
		{
			{2, 1, 1, 1, 1},
			{1, 1, 1, 2, 1},
			{1, 1, 2, 1, 1},
			{2, 1, 1, 1, 1},
			{1, 1, 1, 2, 2},
		},
		{
			{0, 0, 3, 0, 0},
			{0, 0, 3, 0, 0},
			{0, 0, 0, 0, 3},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
	}

	mapper := map[int]engine.Image{
		1: atlas.GetImage("grass_1"),
		2: atlas.GetImage("grass_2"),
		3: atlas.GetImage("tree"),
	}

	tilemap := game.NewTilemap(128, data, mapper, func(bool, engine.Image, interface{}) interface{} {
		return nil
	})
	camera := game.NewCamera()
	animation, _ = game.NewAnimationFromAssetPath("./examples/animation/animation.asset")
	animation.SetState("sw")

	camera.LookAt(64, 128, 0)

	renderer := game.NewIsoRenderer()
	renderer.SetTilemap(tilemap)
	renderer.SetCamera(camera)
	renderer.AddImage(animation)

	game.AddRenderer(renderer)

	game.Run()
}
