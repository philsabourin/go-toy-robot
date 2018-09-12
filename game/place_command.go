package game

import (
	"errors"
	"strconv"
	"strings"
)

// placeCommand places the robot on the table with the provided coordinates.
type placeCommand struct {
	*robot
	args     string
	position position
}

func (c *placeCommand) execute() {
	if c.valid() {
		c.robot.place(c.position)
	}
}

func (c *placeCommand) valid() bool {
	p, err := newPositionFromArgs(c.args)
	if err != nil {
		return false
	}

	c.position = p

	return true
}

func newPositionFromArgs(args string) (position, error) {
	var p position
	var d direction
	var x int
	var y int
	var err error

	a := strings.Split(args, ",")

	if x, err = strconv.Atoi(a[0]); err != nil {
		return p, errors.New("Invalid x coordinate")
	}
	if y, err = strconv.Atoi(a[1]); err != nil {
		return p, errors.New("Invalid y coordinate")
	}

	switch a[2] {
	case "NORTH":
		d = n
	case "WEST":
		d = w
	case "SOUTH":
		d = s
	case "EAST":
		d = e
	default:
		return p, errors.New("Invalid cardinal")
	}

	p = newPosition(x, y, d)

	return p, nil
}
