package game

// reportCommand outputs the robot x and y coordinates with the direction its facing.
type reportCommand struct {
	*robot
}

func (c *reportCommand) execute() {
	c.robot.report()
}

func (c *reportCommand) valid() bool { return true }
