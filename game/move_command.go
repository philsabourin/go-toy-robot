package game

// moveCommand advances the robot by one position forward.
type moveCommand struct {
	*robot
}

func (c *moveCommand) execute() {
	if c.valid() {
		c.robot.move()
	}
}

func (c *moveCommand) valid() bool {
	if c.robot.isOnTable == false {
		return false
	}
	return true
}
