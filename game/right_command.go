package game

// rightCommand turns the robot to its right.
type rightCommand struct {
	*robot
}

func (c *rightCommand) execute() {
	if c.valid() {
		c.robot.right()
	}
}

func (c *rightCommand) valid() bool {
	if c.robot.isOnTable == false {
		return false
	}
	return true
}
