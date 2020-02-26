package viewdrag

import "github.com/hajimehoshi/ebiten"

// Mesh represents set of triangles.
type Mesh struct {
	image    *ebiten.Image
	verteces []ebiten.Vertex
	indices  []uint16

	x, y                      int
	screenWidth, screenHeight int
}

// MoveBy asks for increments of the movements
func (m *Mesh) MoveBy(x, y int) {
	w, h := m.image.Size()

	m.x += x
	m.y += y

	if m.x < 0 {
		m.x = 0
	}
	if m.x > m.screenWidth-w {
		m.x = m.screenWidth - w
	}
	if m.y < 0 {
		m.y = 0
	}
	if m.y > m.screenHeight-h {
		m.y = m.screenHeight - h
	}
}

// GetPosition gives the current position of the mesh
func (m *Mesh) GetPosition() (int, int) {
	return m.x, m.y
}

// Draw draws the mesh.
func (m *Mesh) Draw(screen *ebiten.Image, dx, dy int, alpha float64) {
	op := &ebiten.DrawTrianglesOptions{}

	// op.GeoM.Translate(float64(m.x+dx), float64(m.y+dy))
	for i := range m.verteces {
		m.verteces[i].DstX += float32(dx)
		m.verteces[i].DstY += float32(dy)
	}

	op.ColorM.Scale(1, 1, 1, alpha)
	screen.DrawTriangles(m.verteces, m.indices, m.image, op)
}
