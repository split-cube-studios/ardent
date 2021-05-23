package engine

import "math"

// Cardinal directions.
const (
	N = 1 << iota
	E
	S
	W

	NE = N | E
	NW = N | W
	SE = S | E
	SW = S | W
)

const diag = 3.435 * math.Pi / 180

// CardinalToAngle is a map of cardinal directions
// to angles in dimetric space.
var CardinalToAngle = map[byte]float64{
	N:  3 * math.Pi / 2,            // 90
	E:  0,                          // 0
	S:  math.Pi / 2,                // 270
	W:  math.Pi,                    // 180
	SE: math.Pi/6 - diag,           // 26.565
	SW: 5*math.Pi/6 + diag,         // 153.435
	NE: -(math.Pi / 6) + diag,      // 333.435
	NW: math.Pi + math.Pi/6 - diag, // 206.565,
}

var intervalToCardinal = [8]byte{
	E, SE, S, SW, W, NW, N, NE,
}

// AngleToCardinal convert an angle to a cardinal direction.
func AngleToCardinal(angle float64) byte {
	interval := (int(math.Abs(math.Round(angle/(2*math.Pi/8)))) + 8) % 8

	return intervalToCardinal[interval]
}
