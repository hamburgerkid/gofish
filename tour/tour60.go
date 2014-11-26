package main

import (
	// https://code.google.com/p/go-tour/source/browse/pic/pic.go
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width  int
	height int
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	img_func := func(x, y int) uint8 {
		return uint8(y ^ x)
	}
	v := img_func(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
