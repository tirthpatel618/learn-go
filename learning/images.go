package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) At(x, y int) color.Color {
	v := uint8((x ^ y)) 
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{w: 256, h: 256}
	pic.ShowImage(m)
}
w