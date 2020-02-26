package viewdrag

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// View controls the state of the game
type View struct {
	spr *Sprite
	stk *Stroke
}

// NewView gives custom default values
func NewView(ebitenImage *ebiten.Image, x, y, screenWidth, screenHeight int) *View {
	return &View{spr: &Sprite{
		image:        ebitenImage,
		x:            x,
		y:            y,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}}
}

// Update implements ebiten Update func for main
func (v *View) Update(scr *ebiten.Image) error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		stk := NewStroke(&MouseStrokeSource{})
		stk.SetDraggingObject(v.spr)
		v.stk = stk
	}

	if v.stk != nil {
		v.updateStroke(v.stk)
		if v.stk.IsReleased() {
			v.stk = nil
		}
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if v.stk != nil {
		if spr := v.stk.DraggingObject().(*Sprite); spr != nil {
			dx, dy := v.stk.PositionDiff()
			spr.Draw(scr, dx, dy, 0.5)
			v.spr = spr
		}
		return nil
	}

	v.spr.Draw(scr, 0, 0, 1)
	return nil
}

func (v *View) updateStroke(stroke *Stroke) {
	stroke.Update()
	if !stroke.IsReleased() {
		return
	}

	s := stroke.DraggingObject().(*Sprite)
	if s == nil {
		return
	}

	x, y := stroke.PositionDiff()
	s.MoveBy(x, y)

	v.spr = s

	stroke.SetDraggingObject(nil)
}
