package processing

import (
	"image"
	"image/color"
	"sync"
)

const (
	sepiaRFromR, sepiaRFromG, sepiaRFromB = 0.393, 0.769, 0.189
	sepiaGFromR, sepiaGFromG, sepiaGFromB = 0.349, 0.686, 0.168
	sepiaBFromR, sepiaBFromG, sepiaBFromB = 0.272, 0.534, 0.131
)

const max16BitColor = 65_535

type InvertedFilter struct{}

func (i InvertedFilter) Process(img image.Image) image.Image {
	return applyColorTransformation(img, toInverted)
}

func toInverted(c color.Color) color.Color {
	r, g, b, a := c.RGBA()

	rNew := uint8(255 - r>>8)
	gNew := uint8(255 - g>>8)
	bNew := uint8(255 - b>>8)

	return color.RGBA{
		R: rNew,
		G: gNew,
		B: bNew,
		A: uint8(a >> 8),
	}
}

type SepiaFilter struct{}

func (s SepiaFilter) Process(img image.Image) image.Image {
	return applyColorTransformation(img, toSepia)
}

func toSepia(c color.Color) color.Color {
	r, g, b, a := c.RGBA()

	rFloat := float64(r)
	gFloat := float64(g)
	bFloat := float64(b)

	rNew := uint8(uint32(min(rFloat*sepiaRFromR+gFloat*sepiaRFromG+bFloat*sepiaRFromB, max16BitColor)) >> 8)
	gNew := uint8(uint32(min(rFloat*sepiaGFromR+gFloat*sepiaGFromG+bFloat*sepiaGFromB, max16BitColor)) >> 8)
	bNew := uint8(uint32(min(rFloat*sepiaBFromR+gFloat*sepiaBFromG+bFloat*sepiaBFromB, max16BitColor)) >> 8)
	aNew := uint8(a >> 8)

	return color.RGBA{
		R: rNew,
		G: gNew,
		B: bNew,
		A: aNew,
	}
}

type GrayscaleFilter struct{}

func (g GrayscaleFilter) Process(img image.Image) image.Image {
	return applyColorTransformation(img, toGrayscale)
}

func toGrayscale(c color.Color) color.Color {
	r, g, b, a := c.RGBA()
	grayValue := uint8(uint32((float64(r)*0.2125)+(float64(g)*0.7154)+(float64(b)*0.0721)) >> 8)
	return color.RGBA{R: grayValue, G: grayValue, B: grayValue, A: uint8(a >> 8)}
}

func applyColorTransformation(img image.Image, transformFunc func(c color.Color) color.Color) image.Image {
	b := img.Bounds()
	rgbaImage := image.NewRGBA(b)

	var wg sync.WaitGroup

	for y := b.Min.Y; y < b.Max.Y; y++ {
		wg.Add(1)

		go func(y int) {
			defer wg.Done()

			for x := b.Min.X; x < b.Max.X; x++ {
				color := img.At(x, y)
				sepia := transformFunc(color)
				rgbaImage.Set(x, y, sepia)
			}
		}(y)
	}

	wg.Wait()

	return rgbaImage
}
