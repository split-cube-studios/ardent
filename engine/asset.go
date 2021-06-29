package engine

// Asset is a generic asset container.
type Asset interface {
	ToImage() Image
	ToAtlas() Atlas
	ToAnimation() Animation
	ToScalableImage() ScalableImage
	ToSound() Sound
}
