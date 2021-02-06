package main

import (
	"image/color"
	"log"

	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/assetutil"
	"github.com/split-cube-studios/ardent/engine"
	"golang.org/x/image/font/basicfont"
)

var (
	game  engine.Game
	sound engine.Sound
)

func tick() {

	switch {
	case game.IsKeyJustPressed(engine.KeySpace):
		sound.Play()

	case game.IsKeyJustPressed(engine.KeyL):
		sound.Loop()

	case game.IsKeyJustPressed(engine.KeyP):
		sound.Pause()

	case game.IsKeyJustPressed(engine.KeyR):
		sound.Reset()

	case game.IsKeyJustPressed(engine.KeyUp):
		game.SetVolume("sfx", game.Volume("sfx")+0.1)

	case game.IsKeyJustPressed(engine.KeyDown):
		game.SetVolume("sfx", game.Volume("sfx")-0.1)

	}
}

func main() {
	// create new game instance
	game = ardent.NewGame(
		"Sound",
		854,
		480,
		engine.FlagResizable,
		// tick function
		tick,
		// layout function
		func(ow, oh int) (int, int) {
			return ow / 5, oh / 5
		},
	)

	assetutil.CreateAssets("./")

	sound, _ = game.NewSoundFromAssetPath("sample.asset")

	image := game.NewTextImage(
		"Space: Play\nL: Loop\nP: Pause\nR: Reset",
		250,
		250,
		basicfont.Face7x13,
		color.White,
	)

	renderer := game.NewRenderer()
	renderer.AddImage(image)

	game.AddRenderer(renderer)

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
