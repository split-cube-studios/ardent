package main

import (
	"log"

	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/engine/input"
	"github.com/split-cube-studios/ardent/engine/input/raw"
)

var (
	game      engine.Game
	animation engine.Animation
	keyboard  input.KeySource
	x         float64
)

// tick function.
func tick() {
	if keyboard.IsPressed(raw.KeyA) {
		// walk left
		animation.SetState("sw")
		x--
	} else if keyboard.IsPressed(raw.KeyD) {
		// walk right
		animation.SetState("se")
		x++
	}

	animation.Translate(x, 0)
}

func main() {
	// create new game instance
	game = ardent.NewGame(
		"Keyboard",
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

	// create new renderer and animation
	renderer := game.NewRenderer()
	animation, _ = game.NewAnimationFromAssetPath("./examples/animation/animation.asset")
	animation.Scale(4, 4)
	animation.SetState("sw")

	// add animation to renderer
	renderer.AddImage(animation)

	// add renderer to game and start game
	game.AddRenderer(renderer)
	keyboard = game.NewKeySource()

	err := game.Run()
	if err != nil {
		log.Fatal(err)
	}
}
