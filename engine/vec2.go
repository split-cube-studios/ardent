package engine

import "math"

// Vec2 represents a point in 2D space.
type Vec2 struct {
	X, Y float64
}

// AngleTo returns the angle in radians
// from the endpoint of v to the endpoint of v2.
func (v Vec2) AngleTo(v2 Vec2) float64 {
	return math.Atan2(v2.Y-v.Y, v2.X-v.X)
}

// Lerp returns the linear interpolation from v to v2
// by proportion t.
func (v Vec2) Lerp(v2 Vec2, t float64) Vec2 {
	return Vec2{
		X: (1-t)*v.X + t*v2.X,
		Y: (1-t)*v.Y + t*v2.Y,
	}
}

// Distance returns the distance between the endpoints of v and v2.
func (v Vec2) Distance(v2 Vec2) float64 {
	return math.Sqrt(
		math.Pow(v2.X-v.X, 2) +
			math.Pow(v2.Y-v.Y, 2),
	)
}

// Translate returns the translation of v
// along angle d in radians by magnitude m.
func (v Vec2) Translate(d float64, m float64) Vec2 {
	return Vec2{
		X: v.X + math.Cos(d)*m,
		Y: v.Y + math.Sin(d)*m,
	}
}
