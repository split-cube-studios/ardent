package aautil

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/split-cube-studios/ardent/internal/common"
)

// InvalidTypeError indicates an invalid type in a config.
type InvalidTypeError string

func (i InvalidTypeError) Error() string {
	return fmt.Sprintf("invalid asset type: %s", string(i))
}

type config struct {
	filepath string

	Version string `yml:"version"`
	Type    string `yml:"type"`

	Atlas map[string]struct {
		X int `yml:"x"`
		Y int `yml:"y"`
		W int `yml:"w"`
		H int `yml:"h"`
	} `yml:"atlas,omitempty"`

	FrameWidth  int `yml:"framewidth,omitempty"`
	FrameHeight int `yml:"frameheight,omitempty"`

	Animations map[string]struct {
		Fps   int  `yml:"fps"`
		Loop  bool `yml:"loop,omitempty"`
		Start int  `yml:"start"`
		End   int  `yml:"end"`
	} `yml:"animations,omitempty"`

	Sounds map[string][]string `yml:"sounds,omitempty"`
}

func (c config) toAsset() (*common.Asset, error) {
	asset := common.NewAsset()

	var err error

	switch c.Type {
	case "image":
		asset.Type = common.AssetTypeImage
		asset.Img.Image, err = c.parseImage()

	case "atlas":
		asset.Type = common.AssetTypeAtlas
		for k, v := range c.Atlas {
			asset.AtlasMap[k] = common.AtlasRegion{
				X: uint16(v.X),
				Y: uint16(v.Y),
				W: uint16(v.W),
				H: uint16(v.H),
			}
		}
		asset.Img.Image, err = c.parseImage()

	case "animation":
		asset.Type = common.AssetTypeAnimation
		asset.AnimWidth = uint16(c.FrameWidth)
		asset.AnimHeight = uint16(c.FrameHeight)

		for k, v := range c.Animations {
			asset.AnimationMap[k] = common.Animation{
				Fps:   uint16(v.Fps),
				Loop:  v.Loop,
				Start: uint16(v.Start),
				End:   uint16(v.End),
			}
		}
		asset.Img.Image, err = c.parseImage()

	case "sound":
		asset.Type = common.AssetTypeSound

		for group, sounds := range c.Sounds {

			asset.Snd.Group = group

			for _, sound := range sounds {
				data, err := ioutil.ReadFile(sound)
				if err != nil {
					return nil, err
				}

				asset.Snd.Options = append(asset.Snd.Options, data)
			}
		}

	default:
		return nil, InvalidTypeError(c.Type)
	}

	return asset, err
}

func (c config) parseImage() (image.Image, error) {

	f, err := os.Open(c.filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return png.Decode(f)
}
