package game

// leftCommand turns the robot to its left.
type leftCommand struct {
	*robot
}

func (c *leftCommand) execute() {
	if c.valid() {
		c.robot.left()
	}
}

func (c *leftCommand) valid() bool {
	if c.robot.isOnTable == false {
		return false
	}
	return true
}
