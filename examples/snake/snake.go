package main

import (
	"github.com/split-cube-studios/ardent"
	"github.com/split-cube-studios/ardent/engine"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"math/rand"
	"time"
)

const (
	w = 500
	h = 500

	blockSize = 10
)

func main() {
	rand.Seed(time.Now().UnixNano())
	snake := SnakeGame{}
	game := ardent.NewGame(
		"Snake",
		w, h,
		engine.FlagResizable,
		snake.Tick, func(_ int, _ int) (int, int) {
			return w, h
		})

	snake.setup(game)

	if err := game.Run(); err != nil {
		panic(err)
	}
}

// all the bodyparts of the snake are the same, so just create 1 base tile
func createTile() image.Image {
	tile := image.NewNRGBA(image.Rect(0, 0, blockSize, blockSize))
	fillImage(tile, color.White)
	return tile
}

type SnakeGame struct {
	game     engine.Game
	renderer engine.Renderer

	snake    *Snake
	food     *Food
	baseTile image.Image

	lastMove    time.Time
	timePerMove time.Duration

	lastInput int
}

func (g *SnakeGame) Tick() {
	g.checkInput()

	if time.Since(g.lastMove) < g.timePerMove {
		return
	}

	g.lastMove = time.Now()
	g.snake.Tick()
}

func (g *SnakeGame) checkInput() {
	if g.game.IsKeyJustPressed(engine.KeyW) {
		g.snake.Direction = engine.N
	}
	if g.game.IsKeyJustPressed(engine.KeyA) {
		g.snake.Direction = engine.W
	}
	if g.game.IsKeyJustPressed(engine.KeyS) {
		g.snake.Direction = engine.S
	}
	if g.game.IsKeyJustPressed(engine.KeyD) {
		g.snake.Direction = engine.E
	}
}

func (g *SnakeGame) setup(game engine.Game) {
	g.lastMove = time.Now()
	g.timePerMove = 250 * time.Millisecond // SLOW AF snake
	g.game = game
	g.baseTile = createTile()

	g.renderer = game.NewRenderer()
	game.AddRenderer(g.renderer)

	g.DrawBorder()

	g.snake = &Snake{
		game: g,
	}

	g.snake.Reset()
	g.AddFood()
}

func (g *SnakeGame) NewBodyPart(x, y int) *BodyPart {
	img := g.game.NewImageFromImage(g.baseTile)

	b := BodyPart{
		X:   x,
		Y:   y,
		img: img,
	}

	b.Translate(x, y)
	g.renderer.AddImage(img)
	return &b
}

// DrawBorder creates the game border. The border is always `blockSize` wide
func (g *SnakeGame) DrawBorder() {
	border := image.NewNRGBA(image.Rect(0, 0, w, h))
	for x := 0; x <= w; x++ {
		for y := 0; y <= h; y++ {
			if x < blockSize || x > w-blockSize || y < blockSize || y > h-blockSize {
				border.Set(x, y, color.White)
			}
		}
	}

	g.renderer.AddImage(g.game.NewImageFromImage(border))
}

func (g *SnakeGame) AddFood() {
	img := g.game.NewImageFromImage(g.baseTile)
	g.renderer.AddImage(img)

	g.food = &Food{
		img: img,
	}

	g.MoveFood()
}

func (g *SnakeGame) MoveFood() {
	xMax := w/blockSize - 1
	yMax := h/blockSize - 1

	posX := rand.Intn(xMax)
	posY := rand.Intn(yMax)

	for !g.checkValidPosition((posX+1)*blockSize, (posY+1)*blockSize) {
		posX = rand.Intn(xMax)
		posY = rand.Intn(yMax)
	}

	g.food.Translate((posX+1)*blockSize, (posY+1)*blockSize)
}

type Food struct {
	X, Y int
	img  engine.Image
}

func (b *Food) Translate(x, y int) {
	b.X = x
	b.Y = y
	b.img.Translate(float64(b.X), float64(b.Y))
}

type BodyPart struct {
	X, Y int
	img  engine.Image
	Next *BodyPart
	Prev *BodyPart
}

func (b *BodyPart) Translate(x, y int) {
	b.X = x
	b.Y = y
	b.img.Translate(float64(b.X), float64(b.Y))
}

type Snake struct {
	game      *SnakeGame
	Direction int

	Head *BodyPart
	End  *BodyPart
}

func (s *Snake) Tick() {
	oldHead := s.Head

	newX := oldHead.X
	newY := oldHead.Y
	switch s.Direction {
	case engine.W:
		newX -= blockSize
	case engine.N:
		newY -= blockSize
	case engine.E:
		newX += blockSize
	case engine.S:
		newY += blockSize
	}

	if !s.game.checkValidPosition(newX, newY) {
		s.Reset()
		s.game.MoveFood()
		return
	}

	food := s.game.food
	if food.X == newX && food.Y == newY {
		newHead := s.game.NewBodyPart(newX, newY)
		s.Head = newHead
		oldHead.Prev = newHead

		s.game.MoveFood()
		return
	}

	newHead := s.End

	// mark previous segment as new end of the snake
	s.End = newHead.Prev

	// move tail to oldHead and adjust old oldHead
	newHead.Translate(newX, newY)
	newHead.Prev.Next = nil
	newHead.Next = oldHead

	// ensure the new tail knows the previous segment is the head
	oldHead.Prev = newHead

	s.Head = newHead
}

func (g *SnakeGame) checkValidPosition(x int, y int) bool {
	if x < blockSize {
		return false
	}
	if y < blockSize {
		return false
	}
	if x >= w-blockSize {
		return false
	}
	if y >= h-blockSize {
		return false
	}

	bodyPart := g.snake.Head
	for bodyPart.Next != nil {
		if bodyPart.X == x && bodyPart.Y == y {
			return false
		}

		bodyPart = bodyPart.Next
	}
	return true
}

func (s *Snake) Reset() {
	g := s.game

	bodyPart := s.Head
	for bodyPart != nil {
		bodyPart.img.Dispose()

		bodyPart = bodyPart.Next
	}

	// create a 3 part snake using three bodyparts and linking them together
	head := g.NewBodyPart(250, 250)
	body := g.NewBodyPart(head.X+blockSize, head.Y)
	end := g.NewBodyPart(body.X+blockSize, body.Y)

	head.Next = body

	body.Prev = head
	body.Next = end

	end.Prev = body

	s.Head = head
	s.End = end
	s.Direction = engine.W
}

func fillImage(image draw.Image, color color.Color) {
	for x := image.Bounds().Min.X; x < image.Bounds().Max.X; x++ {
		for y := image.Bounds().Min.Y; y < image.Bounds().Max.Y; y++ {
			image.Set(x, y, color)
		}
	}
}
