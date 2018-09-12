package game

// direction is an interface the Cardinal points.
type direction interface {
	string() string
	xdisplacement() int
	ydisplacement() int
	left() direction
	right() direction
}

// north North Cardinal.
type north struct{}

func (n north) string() string     { return "NORTH" }
func (n north) left() direction    { return w }
func (n north) right() direction   { return e }
func (n north) xdisplacement() int { return 0 }
func (n north) ydisplacement() int { return 1 }

// west West Cardinal.
type west struct{}

func (w west) string() string     { return "WEST" }
func (w west) left() direction    { return s }
func (w west) right() direction   { return n }
func (w west) xdisplacement() int { return -1 }
func (w west) ydisplacement() int { return 0 }

// south South Cardinal.
type south struct{}

func (s south) string() string     { return "SOUTH" }
func (s south) left() direction    { return e }
func (s south) right() direction   { return w }
func (s south) xdisplacement() int { return 0 }
func (s south) ydisplacement() int { return -1 }

// east East Cardinal.
type east struct{}

func (e east) string() string     { return "EAST" }
func (e east) left() direction    { return n }
func (e east) right() direction   { return s }
func (e east) xdisplacement() int { return 1 }
func (e east) ydisplacement() int { return 0 }

// n North Cardinal.
var n = north{}

// w West Cardinal.
var w = west{}

// s South Cardinal.
var s = south{}

// e East Cardinal.
var e = east{}
