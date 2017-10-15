package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"math"
	"os"
	"time"
)

const (
	width, height = 2560, 1440
	marginX       = 200
	marginY       = 100
	paddingX      = 11
	paddingY      = 11
	border        = 1

	// Calculate a lifetime of 90 years.
	deathBy90 = 52 * 90
)

var (
	// Calculate the length for the sides of a square.
	square = int(math.Sqrt((width - 2*marginX - 90*paddingX) * (height - 2*marginY - 49*paddingY) / deathBy90))
)

func main() {
	if err := estimate(); err != nil {
		log.Fatalln(err)
	}
}

func estimate() (err error) {
	f, err := os.Create("out.jpg")
	if err != nil {
		return err
	}
	defer f.Close()

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	borderColor := color.RGBA{110, 110, 110, 255}
	fillColor := white
	num := 0
outer:
	for y := marginY; y < height-marginY; y += square + paddingY {
		for x := marginX; x < width-marginX; x += square + paddingX {
			if num >= deathBy90 {
				break outer
			}
			for i := 0; i < square; i++ {
				for j := 0; j < square; j++ {
					img.Set(x+i, y+j, borderColor)
				}
			}
			for i := border; i < square-border; i++ {
				for j := border; j < square-border; j++ {
					img.Set(x+i, y+j, fillColor)
				}
			}
			num++
		}
	}
	return jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
}

func left() (err error) {
	f, err := os.Create("out.jpg")
	if err != nil {
		return err
	}
	defer f.Close()

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	birth := time.Date(1993, 10, 10, 0, 0, 0, 0, time.UTC)
	today := time.Now()
	lived := today.Sub(birth).Hours() / (24 * 7)

	borderColor := white
	fillColor := white
	num := 0
outer:
	for y := marginY; y < height-marginY; y += square + paddingY {
		for x := marginX; x < width-marginX; x += square + paddingX {
			if num >= int(lived) {
				fillColor = white
				borderColor = color.RGBA{110, 110, 110, 255}
			}
			if num >= deathBy90 {
				break outer
			}
			for i := 0; i < square; i++ {
				for j := 0; j < square; j++ {
					img.Set(x+i, y+j, borderColor)
				}
			}
			for i := border; i < square-border; i++ {
				for j := border; j < square-border; j++ {
					img.Set(x+i, y+j, fillColor)
				}
			}
			num++
		}
	}
	return jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
}
