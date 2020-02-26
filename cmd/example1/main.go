package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/gocs/viewdrag"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	screenWidth  = 320
	screenHeight = 240
)

func main() {

	img, _, err := image.Decode(bytes.NewReader(images.Ebiten_png))
	if err != nil {
		log.Fatal(err)
	}
	ebitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	w, h := ebitenImage.Size()

	g := viewdrag.NewView(
		ebitenImage,
		rand.Intn(screenWidth-w),
		rand.Intn(screenHeight-h),
		screenWidth,
		screenHeight,
	)

	if err := ebiten.Run(g.Update, screenWidth, screenHeight, 1, "Camera Drag"); err != nil {
		log.Fatal("error while running:", err)
	}
}
