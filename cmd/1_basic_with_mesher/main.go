package main

import (
	"bytes"
	"errors"
	"fmt"
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

	v := viewdrag.NewViewWithMesh(
		ebitenImage,
		rand.Intn(screenWidth-w),
		rand.Intn(screenHeight-h),
		screenWidth,
		screenHeight,
		ebiten.MouseButtonMiddle,
	)

	g := &game{v}

	if err := ebiten.Run(g.update, screenWidth, screenHeight, 1, "View Drag"); err != nil {
		log.Fatal("error while running:", err)
	}
}

type game struct {
	v *viewdrag.View
}

func (g *game) update(scr *ebiten.Image) error {
	vx := []ebiten.Vertex{
		ebiten.Vertex{DstX: 100, DstY: 100, SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
		ebiten.Vertex{DstX: 100, DstY: 200, SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
		ebiten.Vertex{DstX: 200, DstY: 200, SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	}
	if err := g.v.SetMesh(vx, []uint16{0, 1, 2, 1, 0, 2}); err != nil {
		return errors.New(fmt.Sprint("error while SetMesh:", err))
	}

	if err := g.v.Compute(scr); err != nil {
		return errors.New(fmt.Sprint("error while computing:", err))
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if err := g.v.Render(scr); err != nil {
		return errors.New(fmt.Sprint("error while rendering:", err))
	}
	return nil
}
