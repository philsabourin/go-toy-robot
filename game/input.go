package game

import (
	"errors"
	"strings"
)

// Input holds the command name and arguments entered by a player in the console.
type input struct {
	command string
	args    string
}

// newInput creates an input to be passed to the commands.
func newInput(ni string) (input, error) {
	var i input
	var cmd []string
	args := ""

	if cmd = strings.Fields(ni); len(cmd) == 0 {
		return i, errors.New("Invalid command")
	}

	if len(cmd) > 1 {
		args = cmd[1]
	}

	return input{
		command: cmd[0],
		args:    args,
	}, nil
}
