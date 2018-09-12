package game

// position holds the current robot position or the position the robot its trying to advance/place itself to.
type position struct {
	xcoord int
	ycoord int
	direction
}

func (p *position) left() position {
	return newPosition(p.xcoord, p.ycoord, p.direction.left())
}

func (p *position) right() position {
	return newPosition(p.xcoord, p.ycoord, p.direction.right())
}

func (p *position) advance() position {
	return newPosition(p.xcoord+p.direction.xdisplacement(), p.ycoord+p.direction.ydisplacement(), p.direction)
}

// newPosition creates a new position where the robot will try to advance/place itself to.
func newPosition(xcoord int, ycoord int, direction direction) position {
	return position{
		xcoord:    xcoord,
		ycoord:    ycoord,
		direction: direction,
	}
}
