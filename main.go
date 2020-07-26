package main

import (
	"time"
	"image/jpeg"
	"image/color"
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"processing"
)
import _ "image/jpeg"
import _ "image/png"
import _ "image/gif"

var _ = fmt.Print
var _ = color.Alpha16Model

func main() {
	path, _ := filepath.Abs("./processing/assets/przedpole.jpg")
	file, e := os.Open(path)
	if e != nil {
		log.Fatalf("%v", e)
	}
	defer file.Close()

	paletteName := os.Args[1]

	palettePath, _ := filepath.Abs("./processing/assets/" + paletteName)
	palettefile, e := os.Open(palettePath)
	if e != nil {
		log.Fatalf("%v", e)
	}
	defer palettefile.Close()

	img, _, e := image.Decode(bufio.NewReader(file))
	if e != nil {
		log.Fatalf("%v", e)
	}

	hexReader := new(processing.HexReader)
	palette, _ := hexReader.Read(palettefile) 

	// log.Printf("%v", palette)

	convertAndSave(img, palette, "kd_output.jpg", kdtree)
	convertAndSave(img, palette, "naive_output.jpg", naive)
}

func convertAndSave(input image.Image, palette color.Palette, outputFileName string, converter func (input image.Image, palette color.Palette) image.Image ) {
	img := converter(input, palette)

	path, _ := filepath.Abs("./processing/assets/" + outputFileName)
	file, e := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if e != nil {
		log.Fatalf("%v", e)
	}
	defer file.Close()

	e = jpeg.Encode(bufio.NewWriter(file), img, nil)
	if e != nil {
		log.Fatalf("%v", e)
	}
}

func kdtree(input image.Image, palette color.Palette) image.Image {
	defer duration(track("KdTree coversion took"))
	p := processing.NewKdTreeProcessor(palette)
	
	return processing.ConvertImage(input, p)
}

func naive(input image.Image, palette color.Palette) image.Image {
	defer duration(track("Naive coversion took"))
	p := processing.NewNaiveProcessor(palette)

	return processing.ConvertImage(input, p)
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}