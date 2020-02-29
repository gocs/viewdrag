package main

import (
	"errors"
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/gocs/viewdrag"
	"github.com/hajimehoshi/ebiten"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	screenWidth  = 320
	screenHeight = 240
)

func main() {
	emptyImage, _ := ebiten.NewImage(16, 16, ebiten.FilterDefault)
	emptyImage.Fill(color.White)

	w, h := emptyImage.Size()
	vx := []ebiten.Vertex{
		{DstX: 0, DstY: 0, SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
		{DstX: 0, DstY: 100, SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
		{DstX: 100, DstY: 100, SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	}

	v := viewdrag.NewViewWithMesh(
		emptyImage,
		vx, []uint16{0, 1, 2},
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
