package game

// command interface for all of the game commands.
type command interface {
	execute()
	valid() bool
}
