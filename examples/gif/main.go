package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"

	"github.com/liamg/gca"
)

func main() {

	var images []image.Image

	grid := gca.NewGrid(300, 150)

	images = append(images, createImageFromGrid(grid))

	grid.SetInitialisationChance(0.35)
	grid.Initialise()
	grid.SetMinNeighboursToBirth(4)
	grid.SetMinNeighboursToRemain(2)

	images = append(images, createImageFromGrid(grid))

	for i := 0; i < 3; i++ {
		grid.Step()
		images = append(images, createImageFromGrid(grid))
	}

	createGif(images, "demo.gif")
}

func createImageFromGrid(grid *gca.Grid) image.Image {
	width, height := grid.Size()
	scale := 2

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width * scale, height * scale}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	cyan := color.RGBA{100, 200, 200, 0xff}
	brown := color.RGBA{128, 32, 32, 0xff}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if grid.Read(x, y) {
				for i := 0; i < scale; i++ {
					img.Set(i+(scale*x), i+(scale*y), cyan)
				}
			} else {
				for i := 0; i < scale; i++ {
					img.Set(i+(scale*x), i+(scale*y), brown)
				}
			}
		}
	}

	return img
}

func createGif(images []image.Image, path string) {

	outGif := &gif.GIF{}
	for _, img := range images {
		palettedImage := image.NewPaletted(img.Bounds(), palette.Plan9)
		draw.Draw(palettedImage, palettedImage.Rect, img, img.Bounds().Min, draw.Over)
		outGif.Image = append(outGif.Image, palettedImage)
		outGif.Delay = append(outGif.Delay, 50)
	}

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, outGif)
}
