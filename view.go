package viewdrag

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// Spriter will match mesh and sprite structs.
// only difference is that mesh will contain triangle vertices nad indeces
type Spriter interface {
	MoveBy(x, y int)
	GetPosition() (int, int)
	Draw(screen *ebiten.Image, dx, dy int, alpha float64)
}

// View controls the state of the game
type View struct {
	spriter Spriter
	stk     *Stroke
	trigger ebiten.MouseButton
}

// NewView gives custom default values
func NewView(ebitenImage *ebiten.Image, x, y, screenWidth, screenHeight int, trigger ebiten.MouseButton) *View {
	return &View{spriter: &Sprite{
		image:        ebitenImage,
		x:            x,
		y:            y,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}, trigger: trigger}
}

// NewViewWithMesh gives custom default values
func NewViewWithMesh(ebitenImage *ebiten.Image, vertices []ebiten.Vertex, indeces []uint16, x, y, screenWidth, screenHeight int, trigger ebiten.MouseButton) *View {
	return &View{spriter: NewMesh(ebitenImage, vertices, indeces, x, y, screenWidth, screenHeight), trigger: trigger}
}

// Compute implements ebiten Update func before draw skipping for main loop
func (v *View) Compute(scr *ebiten.Image) error {
	if inpututil.IsMouseButtonJustPressed(v.trigger) {
		stk := NewStroke(&MouseStrokeSource{btn: v.trigger})
		stk.SetDraggingObject(v.spriter)
		v.stk = stk
	}

	if v.stk != nil {
		v.updateStroke(v.stk)
		if v.stk.IsReleased() {
			v.stk = nil
		}
	}
	return nil
}

// Render implements ebiten Update func after draw skipping for main loop
func (v *View) Render(scr *ebiten.Image) error {
	if v.stk != nil {
		if spr := v.stk.DraggingObject().(Spriter); spr != nil {
			dx, dy := v.stk.PositionDiff()
			spr.Draw(scr, dx, dy, 0.5)
			v.spriter = spr
		}
		return nil
	}
	v.spriter.Draw(scr, 0, 0, 1)
	return nil
}

func (v *View) updateStroke(stroke *Stroke) {
	stroke.Update()
	if !stroke.IsReleased() {
		return
	}

	s := stroke.DraggingObject().(Spriter)
	if s == nil {
		return
	}

	x, y := stroke.PositionDiff()
	s.MoveBy(x, y)

	v.spriter = s

	stroke.SetDraggingObject(nil)
}
