package processing

import "image"

type Processor interface {
	Process(img image.Image) image.Image
}
