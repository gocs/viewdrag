package viewdrag

import "github.com/hajimehoshi/ebiten"

// Sprite represents an image.
type Sprite struct {
	image *ebiten.Image

	x, y                      int
	screenWidth, screenHeight int
}

// MoveBy asks for increments of the movements
func (s *Sprite) MoveBy(x, y int) {
	w, h := s.image.Size()

	s.x += x
	s.y += y

	if s.x < 0 {
		s.x = 0
	}
	if s.x > s.screenWidth-w {
		s.x = s.screenWidth - w
	}
	if s.y < 0 {
		s.y = 0
	}
	if s.y > s.screenHeight-h {
		s.y = s.screenHeight - h
	}
}

// GetPosition gives the current position of the sprite
func (s *Sprite) GetPosition() (int, int) {
	return s.x, s.y
}

// Draw draws the sprite.
func (s *Sprite) Draw(screen *ebiten.Image, dx, dy int, alpha float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.x+dx), float64(s.y+dy))
	op.ColorM.Scale(1, 1, 1, alpha)
	screen.DrawImage(s.image, op)
}
