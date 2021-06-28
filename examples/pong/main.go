package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
)

const (
	paddleWidth, paddleHeight, paddleOffset, velocity = 10, 50, 10, 4.0
)

var (
	xv, yv, w, h = velocity, velocity, 600, 400
)

var (
	game    engine.Game
	ball    engine.Image
	lpaddle engine.Image
	rpaddle engine.Image
)

func collides(a, b engine.Image) bool {
	aw, ah := a.Size()
	bw, bh := b.Size()
	return a.Position().X < b.Position().X+float64(bw) &&
		a.Position().X+float64(aw) > b.Position().X &&
		a.Position().Y < b.Position().Y+float64(bh) &&
		a.Position().Y+float64(ah) > b.Position().Y
}

func foo(paddle engine.Image) float64 {
	return paddle.Position().Y - velocity
}

func bar(paddle engine.Image) float64 {
	return paddle.Position().Y + velocity
}

func checkWallCollisions() {
	if ball.Position().Y < 0 {
		ball.Translate(ball.Position().X, 0)
		yv *= -1
	}
	if ball.Position().Y > float64(h) {
		ball.Translate(ball.Position().X, float64(h))
		yv *= -1
	}
	if ball.Position().X < 0 {
		ball.Translate(0, ball.Position().Y)
		xv *= -1
	}
	if ball.Position().X > float64(w) {
		ball.Translate(float64(w), ball.Position().Y)
		xv *= -1
	}
}

func checkKeyboardInput() {
	if game.IsKeyPressed(engine.KeyUp) && foo(lpaddle) > 0 {
		lpaddle.Translate(lpaddle.Position().X, foo(lpaddle))
	} else if game.IsKeyPressed(engine.KeyDown) && float64(paddleHeight)+bar(lpaddle) < float64(h) {
		lpaddle.Translate(lpaddle.Position().X, bar(lpaddle))
	}
}

func checkPaddleCollisions() {
	if collides(ball, lpaddle) {
		ball.Translate(float64(paddleOffset+paddleWidth), ball.Position().Y)
		xv *= -1
	}
	if collides(ball, rpaddle) {
		fmt.Println(w - paddleOffset - paddleWidth)
		ball.Translate(float64(w-paddleOffset-paddleWidth), ball.Position().Y)
		xv *= -1
	}
}

func main() {
	game = ardent.NewGame("Square",
		w, h,
		engine.FlagResizable,
		func() {
			// move the ball along by adding x and y velocity to its position
			ball.Translate(ball.Position().X+xv, ball.Position().Y+yv)
			checkWallCollisions()
			checkKeyboardInput()
			checkPaddleCollisions()
		},
		func(nw int, nh int) (int, int) {
			w = nw
			h = nh
			return nw, nh
		},
	)

	renderer := game.NewRenderer()
	game.AddRenderer(renderer)

	rpaddle = newRecImage(paddleWidth, paddleHeight)
	rpaddle.Translate(float64(w-paddleWidth-paddleOffset), float64(h-paddleHeight-paddleOffset))

	lpaddle = newRecImage(paddleWidth, paddleHeight)
	lpaddle.Origin(0, 0)
	lpaddle.Translate(float64(paddleOffset), float64(paddleOffset))

	ball = newRecImage(10, 10)
	ball.Origin(0, 0)

	renderer.AddImage(rpaddle)
	renderer.AddImage(lpaddle)
	renderer.AddImage(ball)

	err := game.Run()
	if err != nil {
		panic(err)
	}
}

func newRecImage(x, y int) engine.Image {
	image := image.NewNRGBA(image.Rect(0, 0, x, y))
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			image.Set(i, j, color.White)
		}
	}

	return game.NewImageFromImage(image)
}
