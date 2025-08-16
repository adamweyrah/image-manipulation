package processing

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

var ErrUnsupportedFileType = errors.New("unsupported file type")

func EncodeImage(file *os.File, img image.Image, format string) error {
	var err error

	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(file, img, nil)
	case "png":
		err = png.Encode(file, img)
	default:
		return ErrUnsupportedFileType
	}

	if err != nil {
		return err
	}

	return nil
}

// func EncodeGif(w io.Writer, img image.Image) error {
// 	p := draw.Quantizer
// }
