package nobody

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
)

func New(size int) image.Image {
	if size < 1 || size > 512 {
		log.Fatal("size out of bounds")
	}
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	fileName := fmt.Sprintf("images/%d.png", size)
	asset, err := Asset(fileName)
	if err != nil {
		log.Fatal(err)
	}
	assetFull, err := png.Decode(bytes.NewReader(asset))
	if err != nil {
		log.Fatal(err)
	}
	draw.Draw(img, img.Bounds(), assetFull, image.ZP, draw.Over)

	return img
}