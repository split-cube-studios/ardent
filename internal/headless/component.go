//+build headless

package headless

import (
	"image"
	"image/color"

	"github.com/split-cube-studios/ardent/engine"
	"github.com/split-cube-studios/ardent/internal/common"
	"golang.org/x/image/font"
)

type component struct{}

func (c component) NewAssetFromPath(path string) (engine.Asset, error) {
	return new(Asset), nil
}

func (c component) NewImageFromPath(path string) (engine.Image, error) {
	return new(Image), nil
}

func (c component) NewImageFromAssetPath(path string) (engine.Image, error) {
	return new(Image), nil
}

func (c component) NewImageFromImage(img image.Image) engine.Image {
	return new(Image)
}

func (c component) NewTextImage(text string, w, h int, face font.Face, clr color.Color) engine.Image {
	return new(Image)
}

func (c component) NewAtlasFromAssetPath(path string) (engine.Atlas, error) {
	return new(Atlas), nil
}

func (c component) NewAnimationFromAssetPath(path string) (engine.Animation, error) {
	return new(Animation), nil
}

func (c component) NewSoundFromAssetPath(path string) (engine.Sound, error) {
	return new(Sound), nil
}

func (c component) NewRenderer() engine.Renderer {
	return new(Renderer)
}

func (c component) NewIsoRenderer() engine.IsoRenderer {
	return new(IsoRenderer)
}

func (c component) NewCamera() engine.Camera {
	return new(common.Camera)
}

func (c component) NewCollider() engine.Collider {
	return new(common.Collider)
}
