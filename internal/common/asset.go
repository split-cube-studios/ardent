// Package common contains basic structures for use in engine backends.
package common

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"image"
	"image/png"
)

// AssetSignature is the signature prepended to all Ardent asset files.
const AssetSignature = "Ardent"

// ErrInvalidFiletype occurs when an asset file is of an invalid type.
var ErrInvalidFiletype = errors.New("invalid filetype")

// InvalidAssetType occurs when an invalid AssetType value is encountered.
type InvalidAssetType AssetType

// Error implements error.
func (i InvalidAssetType) Error() string {
	return fmt.Sprintf("invalid asset type: %X", AssetType(i))
}

// AssetType indicates a certain type of asset.
type AssetType byte

const (
	// AssetTypeImage indicates a static image asset.
	AssetTypeImage AssetType = 1 << iota

	// AssetTypeAtlas indicates an image atlas asset.
	AssetTypeAtlas

	// AssetTypeAnimation indicates an animated image asset.
	AssetTypeAnimation

	// AssetTypeSound indicates an audio asset.
	AssetTypeSound
)

// Asset is a basic implementation of engine.Asset.
type Asset struct {
	Img      Image
	AtlasMap map[string]AtlasRegion

	AnimationMap map[string]Animation
	AnimWidth    uint16
	AnimHeight   uint16

	Type AssetType
}

// NewAsset creates an empty Asset.
func NewAsset() *Asset {
	return &Asset{
		AtlasMap:     make(map[string]AtlasRegion),
		AnimationMap: make(map[string]Animation),
	}
}

// Marshal marshals the asset as a []byte.
// It is purposefully named Marshal instead of MarshalBinary to prevent a never-ending loop of gob calling Marshal
// and Marshal calling on gob.
func (a Asset) Marshal() ([]byte, error) {
	switch a.Type {
	case AssetTypeImage, AssetTypeAtlas, AssetTypeAnimation, AssetTypeSound:
	default:
		return nil, InvalidAssetType(a.Type)
	}

	// Format is "Ardent", null byte then gob-encoded data.
	// The "ardent" signature is kept to verify that it's not just random gob-encoded data.
	buf := new(bytes.Buffer)
	buf.WriteString(AssetSignature)
	buf.WriteByte(0)
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(a)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Unmarshal unmarshals the provided []byte as a binary.
// It is purposefully named Unmarshal instead of UnmarshalBinary to prevent a never-ending loop of gob calling Unmarshal
// and Unmarshal calling on gob.
func (a *Asset) Unmarshal(data []byte) error {
	buf := bytes.NewBuffer(data)

	magic, err := buf.ReadString(0)
	if err != nil {
		return err
	}

	if magic[:len(magic)-1] != AssetSignature {
		return ErrInvalidFiletype
	}

	decoder := gob.NewDecoder(buf)
	err = decoder.Decode(a)
	switch a.Type {
	case AssetTypeImage, AssetTypeAtlas, AssetTypeAnimation, AssetTypeSound:
	default:
		return InvalidAssetType(a.Type)
	}
	return err
}

// Image is a wrapper around image.Image that serializes as a PNG.
type Image struct {
	image.Image
}

// MarshalBinary marshals the image as a PNG or just a null byte if there's no image.
func (i Image) MarshalBinary() ([]byte, error) {
	if i.Image == nil {
		// If there's no image, don't bother marshalling it.
		return []byte{0}, nil
	}

	// Otherwise, encode it as a PNG.
	buf := new(bytes.Buffer)
	err := png.Encode(buf, i.Image)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalBinary sets the image's content to nil if it's a null byte or the decoded PNG.
func (i *Image) UnmarshalBinary(data []byte) error {
	if data[0] == 0 {
		// No image.
		i.Image = nil
		return nil
	}

	// Decode it as a PNG.
	var err error
	buf := bytes.NewBuffer(data)
	i.Image, err = png.Decode(buf)
	return err
}
