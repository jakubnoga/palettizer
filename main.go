package main

import (
	"fmt"
	"image"
	"image/color"
	"github.com/jakubnoga/processing"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"	
)

var _ = fmt.Print
var _ = color.Alpha16Model

func main() {
	fmt.Print("dupa")
}

func kdtree(input image.Image, palette color.Palette) image.Image {
	return processing.NewKdTreeProcessor(palette).ConvertImage(input)
}

func naive(input image.Image, palette color.Palette) image.Image {
	return processing.NewNaiveProcessor(palette).ConvertImage(input)
}
