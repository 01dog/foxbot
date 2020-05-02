package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

// RenderIdenticon draws and renders our identicon
func RenderIdenticon(id Identicon, userID string) (filename string, err error) {
	img := image.NewRGBA(image.Rect(0, 0, 50, 50))
	setBG(img)

	for i, v := range id.bitmap {
		if v == 1 {
			drawRect(img, i, id.color)
		}
	}
	filename = userID + "_id.png"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("error:", err)
	}

	defer f.Close()
	png.Encode(f, img)
	return filename, nil
}

func setBG(img *image.RGBA) {
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
}

func drawRect(rgba *image.RGBA, i int, c color.Color) {
	size := 10
	mRow := 5
	rect := image.Rect(
		(i%mRow)*size,
		(i/mRow)*size,
		(i%mRow)*size+size,
		(i/mRow)*size+size,
	)

	draw.Draw(rgba, rect, &image.Uniform{c}, image.ZP, draw.Src)
}
