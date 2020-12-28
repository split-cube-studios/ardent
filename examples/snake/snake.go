package main

import (
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
	"image"
	"image/color"
	"log"
	"time"
)

var (
	w = 500
	h = 500

	blockSize = 10
)

func main() {

	snake := SnakeGame{}
	game := ardent.NewGame("Snake", w, h, 0, snake.Tick, func(_ int, _ int) (int, int) {
		return w, h
	})

	snake.setup(game)

	err := game.Run()
	if err != nil {
		panic(err)
	}
}

// all the bodyparts of the snake are the same, so just create 1 base tile
func createTile() image.Image {
	tile := image.NewNRGBA(image.Rect(0, 0, blockSize, blockSize))
	for x := 0; x < blockSize; x++ {
		for y := 0; y < blockSize; y++ {
			tile.Set(x, y, color.White)
		}
	}

	return tile
}

type SnakeGame struct {
	game     engine.Game
	renderer engine.Renderer

	snake    *Snake
	baseTile image.Image

	lastMove    time.Time
	timePerMove time.Duration
}

func (g *SnakeGame) Tick() {
	now := time.Now()
	if time.Since(g.lastMove) < g.timePerMove {
		return
	}

	log.Println("we ticking!")

	oldHead := g.snake.Head
	newHead := g.snake.End

	// mark previous segment as new end of the snake
	g.snake.End = newHead.Prev

	// TODO: adjust for direction

	// move tail to oldHead and adjust old oldHead
	newHead.Move(oldHead.X-blockSize, oldHead.Y)
	newHead.Prev.Next = nil

	// ensure the new tail knows the previous segment is the head
	oldHead.Prev = newHead

	g.snake.Head = newHead

	g.lastMove = now
}

func (g *SnakeGame) setup(game engine.Game) {
	g.lastMove = time.Now()
	g.timePerMove = 500 * time.Millisecond // SLOW AF snake
	g.game = game
	g.baseTile = createTile()

	g.renderer = game.NewRenderer()
	game.AddRenderer(g.renderer)

	head := g.NewBodyPart(250, 250)
	body := g.NewBodyPart(head.X+blockSize, head.Y)
	end := g.NewBodyPart(body.X+blockSize, body.Y)

	head.Next = &body

	body.Prev = &head
	body.Next = &end

	end.Prev = &body

	g.snake = &Snake{
		Direction: engine.W,
		Head:      &head,
		End:       &end,
	}
}

func (g *SnakeGame) NewBodyPart(x, y int) BodyPart {
	img := g.game.NewImageFromImage(g.baseTile)

	b := BodyPart{
		X:   x,
		Y:   y,
		img: img,
	}

	b.Move(x, y)
	g.renderer.AddImage(img)
	return b
}

type BodyPart struct {
	X, Y int
	img  engine.Image
	Next *BodyPart
	Prev *BodyPart
}

func (b *BodyPart) Move(x, y int) {
	b.X = x
	b.Y = y
	b.img.Translate(float64(b.X), float64(b.Y))
}

type Snake struct {
	Direction int

	Head *BodyPart
	End  *BodyPart
}
