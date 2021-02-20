package engine

import (
	"image"
	"reflect"
	"testing"
)

var testTilemapData = [2][][]int{
	{
		{2, 1, 1, 1, 1, 1},
		{1, 1, 1, 2, 1, 2},
		{1, 1, 2, 1, 1, 1},
		{2, 1, 1, 1, 1, 2},
		{1, 1, 1, 2, 2, 2},
		{1, 1, 2, 1, 1, 1},
	},
	{
		{0, 0, 3, 0, 0, 0},
		{0, 0, 3, 0, 0, 0},
		{0, 0, 0, 0, 3, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 0},
	},
}

var testTilemap = NewTilemap(128, testTilemapData, nil, nil)

func TestIsoToIndex(t *testing.T) {

	tests := map[Vec2]image.Point{
		Vec2{0, 0}:    image.Pt(0, 0),
		Vec2{64, 32}:  image.Pt(1, 0),
		Vec2{-64, 32}: image.Pt(0, 1),
		Vec2{0, 64}:   image.Pt(1, 1),
	}

	for scale := 1; scale < 3; scale++ {
		for input, expected := range tests {

			ix, iy := input.X*float64(scale), input.Y*float64(scale)
			ex, ey := expected.X*scale, expected.Y*scale

			x, y := testTilemap.IsoToIndex(ix, iy)
			if x != ex || y != ey {
				t.Fatalf(
					"Inputs: %f %f. Expected %d %d, got %d %d",
					iy, ix,
					ex, ey,
					x, y,
				)
			}
		}
	}
}

func TestIndexToIso(t *testing.T) {

	tests := map[image.Point]Vec2{
		image.Pt(0, 0): Vec2{0, 0},
		image.Pt(1, 0): Vec2{64, 32},
		image.Pt(0, 1): Vec2{-64, 32},
		image.Pt(1, 1): Vec2{0, 64},
	}

	for scale := 1; scale < 3; scale++ {
		for input, expected := range tests {

			ix, iy := input.X*scale, input.Y*scale
			ex, ey := expected.X*float64(scale), expected.Y*float64(scale)

			x, y := testTilemap.IndexToIso(ix, iy)
			if x != ex || y != ey {
				t.Fatalf(
					"Inputs: %d %d. Expected %f %f, got %f %f",
					ix, iy,
					ex, ey,
					x, y,
				)
			}
		}
	}
}

func TestGetTileValue(t *testing.T) {

	// hadouken test
	for z := 0; z < 2; z++ {
		for y := 0; y < len(testTilemapData[z]); y++ {
			for x := 0; x < len(testTilemapData[z][y]); x++ {
				expected := testTilemapData[z][y][x]
				actual := testTilemap.GetTileValue(x, y, z)
				if expected != actual {
					t.Fatalf(
						"Inputs: x %d y %d z %d. Expected %d, got %d",
						x, y, z,
						expected, actual,
					)
				}
			}
		}
	}

	// out of z upper bounds
	actual := testTilemap.GetTileValue(0, 0, 3)
	if actual != 0 {
		t.Fatalf(
			"Expected 0 for out of z upper bounds. Got %d",
			actual,
		)
	}

	// out of z lower bounds
	actual = testTilemap.GetTileValue(0, 0, -1)
	if actual != 0 {
		t.Fatalf(
			"Expected 0 for out of z lower bounds. Got %d",
			actual,
		)
	}

	// out of horizontal bounds
	actual = testTilemap.GetTileValue(-1, -1, 0)
	if actual != 0 {
		t.Fatalf(
			"Expected 0 for out of horizontal bounds. Got %d",
			actual,
		)
	}
}

func TestNeighbors(t *testing.T) {

	// corner no scale
	expected := []image.Point{image.Pt(1, 0), image.Pt(0, 1)}
	actual := testTilemap.Neighbors(image.Pt(0, 0), 1)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf(
			"Incorrect values for corner neighbors. Expected %v got %v",
			expected,
			actual,
		)
	}

	// near corner with scale
	expected = []image.Point{image.Pt(3, 1), image.Pt(1, 3)}
	actual = testTilemap.Neighbors(image.Pt(1, 1), 2)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf(
			"Incorrect values for near corner neighbors with scale 2. Expected %v got %v",
			expected,
			actual,
		)
	}

	// inner no scale
	expected = []image.Point{
		image.Pt(2, 1),
		image.Pt(1, 2),
		image.Pt(0, 1),
		image.Pt(1, 0),
	}
	actual = testTilemap.Neighbors(image.Pt(1, 1), 1)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf(
			"Incorrect values for inner neighbors. Expected %v got %v",
			expected,
			actual,
		)
	}

	// inner with scale
	expected = []image.Point{
		image.Pt(4, 2),
		image.Pt(2, 4),
		image.Pt(0, 2),
		image.Pt(2, 0),
	}
	actual = testTilemap.Neighbors(image.Pt(2, 2), 2)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf(
			"Incorrect values for inner neighbors. Expected %v got %v",
			expected,
			actual,
		)
	}

	// panic with invalid size
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Size < 1 did not panic")
		}
	}()
	testTilemap.Neighbors(image.Pt(0, 0), 0)
}

func TestInBounds(t *testing.T) {

	// basic in bounds
	for y := 0; y < len(testTilemapData[0]); y++ {
		for x := 0; x < len(testTilemapData[0][y]); x++ {
			if !testTilemap.InBounds(image.Pt(x, y), 1) {
				t.Fatalf(
					"Input %d %d, expected in bounds. Declared out of bounds.",
					x, y,
				)
			}
		}
	}

	// basic out of bounds
	actual := testTilemap.InBounds(image.Pt(0, -1), 1)
	if actual {
		t.Fatal("Input 0 -1, expected out of bounds. Declared in bounds.")
	}

	// scaled in bounds
	for y := 0; y < len(testTilemapData[0])/2; y++ {
		for x := 0; x < len(testTilemapData[0][y])/2; x++ {
			if !testTilemap.InBounds(image.Pt(x, y), 2) {
				t.Fatalf(
					"Input %d %d with scale 2, expected in bounds. Declared out of bounds.",
					x, y,
				)
			}
		}
	}

	// scaled out of bounds
	input := len(testTilemapData[0]) - 1
	actual = testTilemap.InBounds(image.Pt(
		len(testTilemapData[0])-1, 0,
	), 2)
	if actual {
		t.Fatalf(
			"Input %d 0 with scale 2, expected in bounds. Declared out of bounds.",
			input,
		)
	}

	// panic with invalid size
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Size < 1 did not panic")
		}
	}()
	testTilemap.InBounds(image.Pt(0, 0), 0)
}

func IsClear(t *testing.T) {

	// no scale
	for z := 0; z < 2; z++ {
		for y := 0; y < len(testTilemapData[z]); y++ {
			for x := 0; x < len(testTilemapData[z][y]); x++ {
				expected := testTilemapData[z][y][x] == 0
				actual := testTilemap.IsClear(image.Pt(x, y), z, 1)
				if expected != actual {
					t.Fatalf(
						"Input %d %d %d, expected IsClear %t. Declared as %t.",
						x, y, z,
						expected, actual,
					)
				}
			}
		}
	}

	// with scale
	if !testTilemap.IsClear(image.Pt(0, 0), 1, 2) {
		t.Fatal("Input 0 0 1, expected IsClear true. Declared as false.")
	}
	if testTilemap.IsClear(image.Pt(1, 1), 1, 2) {
		t.Fatal("Input 1 1 1, expected IsClear false. Declared as true.")
	}

	// out of bounds
	if testTilemap.IsClear(image.Pt(-1, 0), 1, 1) {
		t.Fatal("Input -1 0 1, expected IsClear false. Declared as true.")
	}

	// panic with invalid size
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Size < 1 did not panic")
		}
	}()
	testTilemap.IsClear(image.Pt(0, 0), 0, 0)
}
