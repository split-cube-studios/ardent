package main

import (
	"log"

	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

var (
	game      engine.Game
	animation engine.Animation
	x         float64
)

// tick function.
func tick() {
	if game.IsKeyPressed(engine.KeyA) {
		// walk left
		animation.SetState("sw")
		x--
	} else if game.IsKeyPressed(engine.KeyD) {
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
		nil,
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

	err := game.Run()
	if err != nil {
		log.Fatal(err)
	}
}
