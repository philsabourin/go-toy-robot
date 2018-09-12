package game

import "fmt"

// robot executes the given commands.
type robot struct {
	position
	*table
	isOnTable bool
}

func (r *robot) place(p position) {
	r.goTo(p)
}

func (r *robot) left() {
	r.goTo(r.position.left())
}

func (r *robot) right() {
	r.goTo(r.position.right())
}

func (r *robot) move() {
	r.goTo(r.position.advance())
}

func (r *robot) report() {
	if r.isOnTable == false {
		fmt.Printf("not in place\n")
	} else {
		fmt.Printf("%d,%d,%s\n", r.position.xcoord, r.position.ycoord, r.position.direction.string())
	}
}

func (r *robot) goTo(position position) {
	if r.table.inbounds(position) {
		if !r.isOnTable {
			r.isOnTable = true
		}
		r.position = position
	}
}
