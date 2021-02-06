package common

import (
	"bytes"
	"image"
	"image/png"
)

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
