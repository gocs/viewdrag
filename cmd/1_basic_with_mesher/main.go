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

	v := viewdrag.NewViewWithMesh(
		emptyImage,
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

var vx = []ebiten.Vertex{
	{DstX: 100, DstY: 100, SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	{DstX: 100, DstY: 500, SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
	{DstX: 500, DstY: 500, SrcX: 0, SrcY: 0, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1},
}

func (g *game) update(scr *ebiten.Image) error {
	if err := g.v.SetMesh(vx, []uint16{0, 1, 2, 1, 2, 3}); err != nil {
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
