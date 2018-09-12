package game

import (
	"bufio"
	"io"
)

var g *game

type game struct {
	*robot
}

// StartAndReadFrom starts the game and reads the input from the io.reader provided.
func StartAndReadFrom(r io.Reader) {

	s := bufio.NewScanner(r)

	g = &game{
		robot: &robot{
			table: &table{
				columns: 5,
				rows:    5,
			},
			isOnTable: false,
		},
	}

	for s.Scan() {
		input, err := newInput(s.Text())
		if err == nil {
			g.execCommand(input)
		}
	}
}

func (g *game) execCommand(i input) {

	var cmd command

	commands := map[string]command{
		"LEFT": &leftCommand{
			robot: g.robot,
		},
		"MOVE": &moveCommand{
			robot: g.robot,
		},
		"PLACE": &placeCommand{
			robot: g.robot,
			args:  i.args,
		},
		"REPORT": &reportCommand{
			robot: g.robot,
		},
		"RIGHT": &rightCommand{
			robot: g.robot,
		},
	}

	if cmd = commands[i.command]; cmd != nil {
		cmd.execute()
	}
}
