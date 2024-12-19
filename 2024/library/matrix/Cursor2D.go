package matrix

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

type Cursor2D struct {
	Row       int
	Column    int
	Direction int
}

func (t Cursor2D) Forward() Cursor2D {
	switch t.Direction {
	case UP:
		return Cursor2D{t.Row - 1, t.Column, UP}
	case RIGHT:
		return Cursor2D{t.Row, t.Column + 1, RIGHT}
	case DOWN:
		return Cursor2D{t.Row + 1, t.Column, DOWN}
	case LEFT:
		return Cursor2D{t.Row, t.Column - 1, LEFT}
	}
	panic("Unsupported direction")
}

func (t Cursor2D) MoveUp() Cursor2D {
	return Cursor2D{t.Row - 1, t.Column, t.Direction}
}
func (t Cursor2D) MoveRight() Cursor2D {
	return Cursor2D{t.Row, t.Column + 1, t.Direction}
}
func (t Cursor2D) MoveDown() Cursor2D {
	return Cursor2D{t.Row + 1, t.Column, t.Direction}
}
func (t Cursor2D) MoveLeft() Cursor2D {
	return Cursor2D{t.Row, t.Column - 1, t.Direction}
}

func (t Cursor2D) TurnUp() Cursor2D {
	return Cursor2D{t.Row, t.Column, UP}
}
func (t Cursor2D) TurnRight() Cursor2D {
	return Cursor2D{t.Row, t.Column, RIGHT}
}
func (t Cursor2D) TurnDown() Cursor2D {
	return Cursor2D{t.Row, t.Column, DOWN}
}
func (t Cursor2D) TurnLeft() Cursor2D {
	return Cursor2D{t.Row, t.Column, LEFT}
}

func (t Cursor2D) TurnClockwise() Cursor2D {
	switch t.Direction {
	case UP:
		return Cursor2D{t.Row, t.Column, RIGHT}
	case RIGHT:
		return Cursor2D{t.Row, t.Column, DOWN}
	case DOWN:
		return Cursor2D{t.Row, t.Column, LEFT}
	case LEFT:
		return Cursor2D{t.Row, t.Column, UP}
	}
	panic("Unsupported direction")
}

func (t Cursor2D) Clone() Cursor2D {
	return Cursor2D{t.Row, t.Column, t.Direction}
}

func (t Cursor2D) ToVector() Vector[int] {
	return Vector[int]{data: []int{t.Row, t.Column}, size: 2}
}
