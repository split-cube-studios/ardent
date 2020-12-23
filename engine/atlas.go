package engine

// Atlas is a set of named Images.
type Atlas interface {
	GetImage(string) Image
}
