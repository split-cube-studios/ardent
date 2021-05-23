package engine

import "math"

// Vec2 represents a point in 2D space.
type Vec2 struct {
	X, Y float64
}

// Add adds two vectors together
func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

// Sub subtracts a vector.
func (v Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

// Angle returns the angle in radians.
func (v Vec2) Angle() float64 {
	return Vec2{}.AngleTo(v)
}

// AngleTo returns the angle from the end of
// vector v to the end of vector v2.
func (v Vec2) AngleTo(v2 Vec2) float64 {
	return math.Atan2(v2.Y-v.Y, v2.X-v.X)
}

// Lerp returns the linear interpolation between
// two vectors by percent t.
func (v Vec2) Lerp(v2 Vec2, t float64) Vec2 {
	return Vec2{
		X: (1-t)*v.X + t*v2.X,
		Y: (1-t)*v.Y + t*v2.Y,
	}
}

// Distance returns the distance between two vectors.
func (v Vec2) Distance(v2 Vec2) float64 {
	return math.Sqrt(
		math.Pow(v2.X-v.X, 2) +
			math.Pow(v2.Y-v.Y, 2),
	)
}

// Translate returns the translation of a vector
// towards direction d by magnitude m.
func (v Vec2) Translate(d float64, m float64) Vec2 {
	return Vec2{
		X: v.X + math.Cos(d)*m,
		Y: v.Y + math.Sin(d)*m,
	}
}

// Scale multiplies vector v by a scalar value s
// and returns the result.
func (v Vec2) Scale(s float64) Vec2 {
	return Vec2{
		X: v.X * s,
		Y: v.Y * s,
	}
}

// Dot returns the dot product between two vectors.
func (v Vec2) Dot(v2 Vec2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

// Length returns the length of a vector.
func (v Vec2) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

// Normalize normalizes a vector magnitude.
func (v Vec2) Normalize() Vec2 {
	return v.Scale(1 / v.Length())
}
