package viewdrag

import "github.com/hajimehoshi/ebiten"

// Sprite represents an image.
type Sprite struct {
	image *ebiten.Image

	// actually the bounds of the mesh
	// base vector of the whole mesh
	x, y int
	// farthest vector in all of the meshes
	width, height int

	// dimensions of the viewport
	scrWidth, scrHeight int
}

// NewSprite generates new Sprite
func NewSprite(ebitenImage *ebiten.Image, x, y, screenWidth, screenHeight int) *Sprite {
	w, h := ebitenImage.Size()
	return &Sprite{
		image:     ebitenImage,
		x:         x,
		y:         y,
		width:     int(w),
		height:    int(h),
		scrWidth:  screenWidth,
		scrHeight: screenHeight,
	}
}

// MoveBy asks for increments of the movements
func (s *Sprite) MoveBy(x, y int) {
	s.x += x
	s.y += y

	s.x, s.y = keepSpriteInsideView(s.x, s.y, s.scrWidth, s.scrHeight, s.width, s.height)
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

// control displacement of the sprite
// whether the sprite is big or small, keep the sprite inside view
func keepSpriteInsideView(x, y, scrWidth, scrHeight, width, height int) (int, int) {
	if scrWidth < width {
		if x < scrWidth-width {
			x = scrWidth - width
		}
		if x > 0 {
			x = 0
		}
	} else {
		if x < 0 {
			x = 0
		}
		if x > scrWidth-width {
			x = scrWidth - width
		}
	}
	if scrHeight < height {
		if y < scrHeight-height {
			y = scrHeight - height
		}
		if y > 0 {
			y = 0
		}
	} else {
		if y < 0 {
			y = 0
		}
		if y > scrHeight-height {
			y = scrHeight - height
		}
	}
	return x, y
}
