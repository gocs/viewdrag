package viewdrag

// StrokeSource represents a input device to provide strokes.
type StrokeSource interface {
	Position() (int, int)
	IsJustReleased() bool
}

// Stroke manages the current drag state by mouse.
type Stroke struct {
	source StrokeSource

	// initX and initY represents the position when dragging starts.
	initX int
	initY int

	// currentX and currentY represents the current position
	currentX int
	currentY int

	released bool

	// draggingObject represents a Spriter
	// that is being dragged.
	draggingSpriter Spriter
}

// NewStroke asks for devices' custom controls and provides custom default values
func NewStroke(source StrokeSource) *Stroke {
	cx, cy := source.Position()
	return &Stroke{
		source:   source,
		initX:    cx,
		initY:    cy,
		currentX: cx,
		currentY: cy,
	}
}

// Update sets the current stroke's position
func (s *Stroke) Update() {
	if s.released {
		return
	}
	if s.source.IsJustReleased() {
		s.released = true
		return
	}
	x, y := s.source.Position()
	s.currentX = x
	s.currentY = y
}

// IsReleased asks if the mouse button is up
func (s *Stroke) IsReleased() bool {
	return s.released
}

// Position gives the current stroke position
func (s *Stroke) Position() (int, int) {
	return s.currentX, s.currentY
}

// PositionDiff gives the position after the change from initial to current
func (s *Stroke) PositionDiff() (int, int) {
	dx := s.currentX - s.initX
	dy := s.currentY - s.initY
	return dx, dy
}

// DraggingSprite gives the Spriter being dragged
func (s *Stroke) DraggingSprite() Spriter {
	return s.draggingSpriter
}

// SetDraggingSpriter sets the Spriter to be dragged
func (s *Stroke) SetDraggingSpriter(sprite Spriter) {
	s.draggingSpriter = sprite
}
