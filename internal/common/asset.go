// Package common contains basic structures for use in engine backends.
package common

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"image"
	"image/png"
)

// ErrInvalidFiletype occurs when an asset file is of an invalid type.
var ErrInvalidFiletype = errors.New("invalid filetype")

// WrongNumBytesError occurs when an incorrect number of bytes is read.
type WrongNumBytesError struct {
	Expect, Got int
}

// Error implements error.
func (w WrongNumBytesError) Error() string {
	return fmt.Sprintf("expected %d bytes, got %d", w.Expect, w.Got)
}

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
	Img      image.Image
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

// MarshalBinary implements encoding.BinaryMarshaler.
func (a *Asset) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)

	buf.WriteString("ardent")
	buf.WriteByte(0)
	buf.WriteByte(byte(a.Type))

	switch a.Type {
	case AssetTypeImage:
	case AssetTypeAtlas:
		buf.WriteByte(byte(len(a.AtlasMap)))

		for k, v := range a.AtlasMap {
			buf.WriteString(k)
			buf.WriteByte(0)

			data := make([]byte, 8)
			binary.LittleEndian.PutUint16(data[:2], v.X)
			binary.LittleEndian.PutUint16(data[2:4], v.Y)
			binary.LittleEndian.PutUint16(data[4:6], v.W)
			binary.LittleEndian.PutUint16(data[6:8], v.H)

			buf.Write(data)
		}

	case AssetTypeAnimation:
		buf.WriteByte(byte(len(a.AnimationMap)))

		data := make([]byte, 4)
		binary.LittleEndian.PutUint16(data[:2], a.AnimWidth)
		binary.LittleEndian.PutUint16(data[2:4], a.AnimHeight)

		buf.Write(data)

		for k, v := range a.AnimationMap {
			buf.WriteString(k)
			buf.WriteByte(0)

			data = make([]byte, 7)
			binary.LittleEndian.PutUint16(data[:2], v.Fps)
			binary.LittleEndian.PutUint16(data[2:4], v.Start)
			binary.LittleEndian.PutUint16(data[4:6], v.End)

			if v.Loop {
				data[6] = 1
			}

			buf.Write(data)
		}
	case AssetTypeSound:
		// TODO
	default:
		return nil, InvalidAssetType(a.Type)
	}

	if a.Type != AssetTypeSound {
		if err := png.Encode(buf, a.Img); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (a *Asset) UnmarshalBinary(data []byte) error {
	buf := bytes.NewBuffer(data)

	magic, err := buf.ReadString(0)
	if err != nil {
		return err
	}

	if magic[:len(magic)-1] != "ardent" {
		return ErrInvalidFiletype
	}

	t, err := buf.ReadByte()
	if err != nil {
		return err
	}

	a.Type = AssetType(t)
	switch a.Type {
	case AssetTypeImage:
	case AssetTypeAtlas:
		count, err := buf.ReadByte()
		if err != nil {
			return err
		}

		for i := 0; i < int(count); i++ {
			k, err := buf.ReadString(0)
			if err != nil {
				return err
			}

			regData := make([]byte, 8)
			if n, err := buf.Read(regData); n != len(regData) {
				if err != nil {
					return err
				}

				return WrongNumBytesError{
					Expect: len(regData),
					Got:    n,
				}
			}

			a.AtlasMap[k[:len(k)-1]] = AtlasRegion{
				X: binary.LittleEndian.Uint16(regData[:2]),
				Y: binary.LittleEndian.Uint16(regData[2:4]),
				W: binary.LittleEndian.Uint16(regData[4:6]),
				H: binary.LittleEndian.Uint16(regData[6:8]),
			}
		}

	case AssetTypeAnimation:
		count, err := buf.ReadByte()
		if err != nil {
			return err
		}

		animSize := make([]byte, 4)
		if n, err := buf.Read(animSize); n != len(animSize) {
			if err != nil {
				return err
			}

			return WrongNumBytesError{
				Expect: len(animSize),
				Got:    n,
			}
		}

		a.AnimWidth = binary.LittleEndian.Uint16(animSize[:2])
		a.AnimHeight = binary.LittleEndian.Uint16(animSize[2:4])

		for i := 0; i < int(count); i++ {
			k, err := buf.ReadString(0)
			if err != nil {
				return err
			}

			animData := make([]byte, 7)
			if n, err := buf.Read(animData); n != len(animData) {
				if err != nil {
					return err
				}

				return WrongNumBytesError{
					Expect: len(animData),
					Got:    n,
				}
			}

			anim := Animation{
				Fps:   binary.LittleEndian.Uint16(animData[:2]),
				Start: binary.LittleEndian.Uint16(animData[2:4]),
				End:   binary.LittleEndian.Uint16(animData[4:6]),
			}

			if animData[6] > 0 {
				anim.Loop = true
			}

			a.AnimationMap[k[:len(k)-1]] = anim
		}

	case AssetTypeSound:
		// TODO
	default:
		return InvalidAssetType(a.Type)
	}

	if a.Type == AssetTypeSound {
		// TODO
		return nil
	}

	a.Img, err = png.Decode(buf)

	return err
}
