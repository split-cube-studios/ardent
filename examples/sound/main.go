package main

import (
	"log"

	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/assetutil"
	"github.com/split-cube-studios/ardent/engine"
)

var (
	game  engine.Game
	sound engine.Sound
)

func tick() {

	if game.IsKeyJustPressed(engine.KeySpace) {
		sound.Play()
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
			return ow, oh
		},
	)

	assetutil.CreateAssets("./")

	sound, _ = game.NewSoundFromAssetPath("hit.asset")

	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}