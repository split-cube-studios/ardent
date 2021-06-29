package engine

import (
	"image"
	"image/color"

	"golang.org/x/image/font"
)

// Component produces backend dependent components.
type Component interface {
	AssetComponent
	ImageComponent
	SoundComponent
	RendererComponent
}

// AssetComponent produces asset components.
type AssetComponent interface {
	NewAssetFromPath(string) (Asset, error)
	ImageComponent
	SoundComponent
}

// ImageComponent produces image components.
type ImageComponent interface {
	// NewImageFromPath returns an image from an image file path.
	NewImageFromPath(string) (Image, error)

	// NewImageFromAssetPath returns an image from an asset file path.
	NewImageFromAssetPath(string) (Image, error)

	// NewImageFromImage returns an image from an image.Image.
	NewImageFromImage(image.Image) Image

	// NewTextImage returns an image containing text.
	NewTextImage(text string, face font.Face, c color.Color) Image

	// NewImageFromLayers merges image layers on in order and returns a new image.
	NewImageFromLayers(...Image) Image

	// NewAtlasFromAssetPath returns an image atlas from an asset file path.
	NewAtlasFromAssetPath(string) (Atlas, error)

	// NewAnimationFromAssetPath returns an animation from an asset file path.
	NewAnimationFromAssetPath(string) (Animation, error)

	// NewScalableImageFromAssetPath returns a scalable image from an asset file path.
	NewScalableImageFromAssetPath(string) (ScalableImage, error)
}

// SoundComponent produces sound components.
type SoundComponent interface {
	NewSoundFromAssetPath(string) (Sound, error)
}

// RendererComponent produces renderer components.
type RendererComponent interface {
	NewRenderer() Renderer
	NewIsoRenderer() IsoRenderer
}
