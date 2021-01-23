package engine

import (
	"image"
	"image/color"

	"golang.org/x/image/font"
)

// Component produces engine components.
type Component interface {
	NewAssetFromPath(string) (Asset, error)

	NewImageFromPath(string) (Image, error)
	NewImageFromAssetPath(string) (Image, error)
	NewImageFromImage(image.Image) Image

	NewTextImage(string, int, int, font.Face, color.Color) Image

	NewAtlasFromAssetPath(string) (Atlas, error)

	NewAnimationFromAssetPath(string) (Animation, error)

	NewSoundFromAssetPath(string) (Sound, error)

	NewRenderer() Renderer
	NewIsoRenderer() IsoRenderer

	NewTilemap(int, [2][][]int, map[int]Image, TileOverlapEvent) Tilemap

	NewCamera() Camera

	NewCollider() Collider
}
