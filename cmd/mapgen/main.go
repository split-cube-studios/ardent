package main

import (
	"fmt"
	"image/png"
	"math/rand"
	"os"
	"time"

	"image/color"

	"github.com/split-cube-studios/ardent/mapgen"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	const w, h = 512, 512

	requiredRooms := make([]mapgen.Room, 100)
	for i := 0; i < len(requiredRooms); i++ {
		requiredRooms[i] = mapgen.NewNaturalRoom(
			rand.Intn(5)*6+12, rand.Intn(5)*6+12,
			1, mapgen.RoomPolicy{
				Required: true,
			},
		)
	}

	g := mapgen.NewGenerator(
		mapgen.GeneratorOptions{
			Width:         w,
			Height:        h,
			RequiredRooms: requiredRooms,
			RoomAlign:     4,
			PathAlg:       mapgen.NewBasicPath(6, 1, 2),
			FloorTile:     1,
			WallTile:      2,
		},
	)

	t, err := g.Generate()
	if err != nil {
		fmt.Printf("Failed to generate map: %v\n", err)
		return
	}

	colors := map[int]color.Color{
		1: color.White,
		2: color.Black,
	}

	img := t.Image(colors)

	f, err := os.Create("map.png")
	if err != nil {
		fmt.Printf("Failed to create map.png: %v\n", err)
		return
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		fmt.Printf("Failed to encode PNG: %v\n", err)
		return
	}

	fmt.Println("map.png generated")
}
