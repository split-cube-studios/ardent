package engine

// Camera manages the area being viewed.
type Camera interface {
	LookAt(float64, float64, float64)
	Position() (float64, float64)
}
