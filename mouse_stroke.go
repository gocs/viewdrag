package viewdrag

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// MouseStrokeSource is a StrokeSource implementation of mouse.
type MouseStrokeSource struct {
	btn ebiten.MouseButton
}

// Position gives the current position by a mouse
func (m *MouseStrokeSource) Position() (int, int) {
	return ebiten.CursorPosition()
}

// IsJustReleased asks if the mouse button is up
func (m *MouseStrokeSource) IsJustReleased() bool {
	return inpututil.IsMouseButtonJustReleased(m.btn)
}
