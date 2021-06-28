//+build !headless

package ebiten

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg" // jpeg support
	_ "image/png"  // png support
	"io/ioutil"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"

	"github.com/split-cube-studios/ardent/engine"
)

type component struct {
	assetCache map[string]Asset
	sc         *SoundControl
}

func newComponent(sc *SoundControl) *component {
	return &component{
		assetCache: make(map[string]Asset),
		sc:         sc,
	}
}

func (c *component) NewAssetFromPath(path string) (engine.Asset, error) {
	if asset, ok := c.assetCache[path]; ok {
		return &asset, nil
	}

	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to decode asset: %w", err)
	}

	a := new(Asset)
	if err = a.UnmarshalBinary(d); err != nil {
		return nil, err
	}

	c.assetCache[path] = *a

	return a, nil
}

func (c *component) NewImageFromPath(path string) (engine.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open image path: %w", err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	return c.NewImageFromImage(img), nil
}

func (c *component) NewImageFromAssetPath(path string) (engine.Image, error) {
	a, err := c.NewAssetFromPath(path)
	if err != nil {
		return nil, err
	}

	return a.ToImage(), nil
}

func (c *component) NewImageFromImage(img image.Image) engine.Image {
	return &Image{
		img:        ebiten.NewImageFromImage(img),
		sx:         1,
		sy:         1,
		alpha:      1,
		r:          1,
		g:          1,
		b:          1,
		renderable: true,
	}
}

func (c *component) NewTextImage(txt string, w, h int, face font.Face, clr color.Color) engine.Image {
	img := ebiten.NewImage(w, h)
	text.Draw(img, txt, face, 0, face.Metrics().Height.Round(), clr)

	return &Image{
		img:        img,
		sx:         1,
		sy:         1,
		r:          1,
		g:          1,
		b:          1,
		alpha:      1,
		renderable: true,
	}
}

func (c *component) NewImageFromLayers(layers ...engine.Image) engine.Image {

	var baseImg *Image

	for _, layer := range layers {

		if !layer.IsRenderable() {
			continue
		}

		img := engineImageToLocalImage(layer)

		if baseImg == nil {
			baseImg = img
			continue
		}

		baseImg.layers = append(baseImg.layers, layer)
	}

	return baseImg
}

func (c *component) NewAtlasFromAssetPath(path string) (engine.Atlas, error) {
	a, err := c.NewAssetFromPath(path)
	if err != nil {
		return nil, err
	}

	return a.ToAtlas(), nil
}

func (c *component) NewAnimationFromAssetPath(path string) (engine.Animation, error) {
	a, err := c.NewAssetFromPath(path)
	if err != nil {
		return nil, err
	}

	return a.ToAnimation(), nil
}

func (c *component) NewSoundFromAssetPath(path string) (engine.Sound, error) {
	a, err := c.NewAssetFromPath(path)
	if err != nil {
		return nil, err
	}

	sound := a.ToSound().(*Sound)
	sound.sc = c.sc
	sound.volume = c.sc.Volume(sound.group)

	c.sc.addSound(sound)

	return sound, nil
}

func (c *component) NewRenderer() engine.Renderer {
	return NewRenderer()
}

func (c *component) NewIsoRenderer() engine.IsoRenderer {
	return NewIsoRenderer()
}
