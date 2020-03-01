package viewdrag

import "github.com/hajimehoshi/ebiten"

// Mesh represents set of triangles.
type Mesh struct {
	// image basis for the mesh
	image *ebiten.Image

	// the mesh themselves
	vertices []ebiten.Vertex
	indices  []uint16

	// actually the bounds of the mesh
	// base vector of the whole mesh
	x, y int
	// farthest vector in all of the meshes
	width, height int

	// dimensions of the viewport
	scrWidth, scrHeight int
}

// NewMesh generates new mesh with farthest vector
func NewMesh(ebitenImage *ebiten.Image, vertices []ebiten.Vertex, indeces []uint16, x, y, screenWidth, screenHeight int) *Mesh {
	w, h := Size(vertices)

	return &Mesh{
		image:     ebitenImage,
		vertices:  vertices,
		indices:   indeces,
		x:         x,
		y:         y,
		width:     int(w),
		height:    int(h),
		scrWidth:  screenWidth,
		scrHeight: screenHeight,
	}
}

// Size returns the width and height of the mesh
func Size(vertices []ebiten.Vertex) (w, h float32) {
	for _, v := range vertices {
		if w < v.DstX {
			w = v.DstX
		}
		if h < v.DstY {
			h = v.DstY
		}
	}
	return
}

// MoveBy asks for increments of the movements
func (m *Mesh) MoveBy(x, y int) {
	m.x += x
	m.y += y

	// if exceeds bounds return back
	// if m.x, m.y is <1/2 of m.width, m.height { set to 1/2 m.width, m.height}
	// m.width favorably gives buffer
	if m.x < m.scrWidth/2-m.width {
		m.x = m.scrWidth/2 - m.width
	}
	if m.x > m.scrWidth/2 {
		m.x = m.scrWidth / 2
	}
	if m.y < m.scrHeight/2-m.height {
		m.y = m.scrHeight/2-m.height
	}
	if m.y > m.scrHeight/2 {
		m.y = m.scrHeight/2
	}
}

// GetPosition gives the displacement position of the of the whole mesh
func (m *Mesh) GetPosition() (int, int) {
	return m.x, m.y
}

// Draw draws the mesh.
func (m *Mesh) Draw(screen *ebiten.Image, dx, dy int, alpha float64) {
	op := &ebiten.DrawTrianglesOptions{}

	// op.GeoM.Translate(float64(m.x+dx), float64(m.y+dy))
	vx := []ebiten.Vertex{}
	for _, v := range m.vertices {
		vx = append(vx, ebiten.Vertex{
			DstX:   v.DstX + float32(m.x+dx),
			DstY:   v.DstY + float32(m.y+dy),
			ColorA: v.ColorA,
			ColorB: v.ColorB,
			ColorG: v.ColorG,
			ColorR: v.ColorR,
			SrcX:   v.SrcX,
			SrcY:   v.SrcY,
		})
	}

	op.ColorM.Scale(1, 1, 1, alpha)
	screen.DrawTriangles(vx, m.indices, m.image, op)
}
