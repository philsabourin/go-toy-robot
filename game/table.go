package game

// table is the area for the robot to move on to.
type table struct {
	columns int
	rows    int
}

func (t *table) inbounds(p position) bool {
	if p.xcoord < 0 || p.xcoord > t.columns-1 || p.ycoord < 0 || p.ycoord > t.rows-1 {
		return false
	}

	return true
}
