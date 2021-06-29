package engine

type ScalableImage interface {
	GetImage(int, int) Image
}
