package engine

import (
	"image"
	"image/color"

	"golang.org/x/image/font"
)

// Component produces engine components.
type Component interface {
	AssetComponent
	ImageComponent
	SoundComponent
	RendererComponent
	CameraComponent
	PhysicsComponent
}

// AssetComponent produces asset components.
type AssetComponent interface {
	NewAssetFromPath(string) (Asset, error)
	ImageComponent
	SoundComponent
}

// ImageComponent produces image components.
type ImageComponent interface {
	NewImageFromPath(string) (Image, error)
	NewImageFromAssetPath(string) (Image, error)
	NewImageFromImage(image.Image) Image
	NewTextImage(string, int, int, font.Face, color.Color) Image
	NewAtlasFromAssetPath(string) (Atlas, error)
	NewAnimationFromAssetPath(string) (Animation, error)
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

// CameraComponent produces camera components.
type CameraComponent interface {
	NewCamera() Camera
}

// PhysicsComponent produces physics components.
type PhysicsComponent interface {
	NewCollider() Collider
}
