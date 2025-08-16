package processing

import (
	"image"
	"os"
)

func DecodeImage(file *os.File) (image.Image, string, error) {
	img, format, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}
	return img, format, nil
}
