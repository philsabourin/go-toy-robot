package main

import (
	"regexp"
	"strings"
	"testing"

	capturer "github.com/kami-zh/go-capturer"
	"github.com/philsabourin/go-toy-robot/game"
)

func TestScenario1(t *testing.T) {

	scenario := `
	MOVE
	LEFT
	RIGHT
	REPORT
	`

	want := `
	not in place
	`

	got := capturer.CaptureStdout(func() {
		game.StartAndReadFrom(strings.NewReader(scenario))
	})

	if isOutputEqual(got, want) == false {
		t.Error("outputs are not equal")
	}
}

func TestScenario2(t *testing.T) {

	scenario := `
	PLACE 1,2,NORTH
	REPORT
	`

	want := `
	1,2,NORTH
	`

	got := capturer.CaptureStdout(func() {
		game.StartAndReadFrom(strings.NewReader(scenario))
	})

	if isOutputEqual(got, want) == false {
		t.Error("outputs are not equal")
	}
}

func TestScenario3(t *testing.T) {

	scenario := `
	PLACE 9,9,NORTH
	REPORT
	`

	want := `
	not in place
	`

	got := capturer.CaptureStdout(func() {
		game.StartAndReadFrom(strings.NewReader(scenario))
	})

	if isOutputEqual(got, want) == false {
		t.Error("outputs are not equal")
	}
}

func TestScenario4(t *testing.T) {

	scenario := `
	PLACE 1,1,NORTH
	RIGHT
	REPORT
	RIGHT
	REPORT
	RIGHT
	REPORT
	RIGHT
	REPORT
	`

	want := `
	1,1,EAST
	1,1,SOUTH
	1,1,WEST
	1,1,NORTH
	`

	got := capturer.CaptureStdout(func() {
		game.StartAndReadFrom(strings.NewReader(scenario))
	})

	if isOutputEqual(got, want) == false {
		t.Error("outputs are not equal")
	}
}

func TestScenario5(t *testing.T) {

	scenario := `
	PLACE 1,1,NORTH
	LEFT
	REPORT
	LEFT
	REPORT
	LEFT
	REPORT
	LEFT
	REPORT
	`

	want := `
	1,1,WEST
	1,1,SOUTH
	1,1,EAST
	1,1,NORTH
	`

	got := capturer.CaptureStdout(func() {
		game.StartAndReadFrom(strings.NewReader(scenario))
	})

	if isOutputEqual(got, want) == false {
		t.Error("outputs are not equal")
	}
}

func TestScenario6(t *testing.T) {

	scenario := `
	PLACE 1,1,NORTH
	MOVE
	RIGHT
	REPORT
	MOVE
	RIGHT
	REPORT
	MOVE
	RIGHT
	REPORT
	MOVE
	RIGHT
	REPORT
	`

	want := `
	1,2,EAST
	2,2,SOUTH
	2,1,WEST
	1,1,NORTH
	`

	got := capturer.CaptureStdout(func() {
		game.StartAndReadFrom(strings.NewReader(scenario))
	})

	if isOutputEqual(got, want) == false {
		t.Error("outputs are not equal")
	}
}

func TestScenario7(t *testing.T) {

	scenario := `
	PLACE 4,4,NORTH
	MOVE
	MOVE
	MOVE
	MOVE
	MOVE
	REPORT
	`

	want := `
	4,4,NORTH
	`

	got := capturer.CaptureStdout(func() {
		game.StartAndReadFrom(strings.NewReader(scenario))
	})

	if isOutputEqual(got, want) == false {
		t.Error("outputs are not equal")
	}
}

func TestScenario8(t *testing.T) {

	scenario := `
	AUTODESTRUCT
	TAKEOFF
	KILL
	REPORT
	`

	want := `
	not in place
	`

	got := capturer.CaptureStdout(func() {
		game.StartAndReadFrom(strings.NewReader(scenario))
	})

	if isOutputEqual(got, want) == false {
		t.Error("outputs are not equal")
	}
}

func isOutputEqual(s1 string, s2 string) bool {
	space := regexp.MustCompile(`\s+`)
	o1 := space.ReplaceAllString(s1, "")
	o2 := space.ReplaceAllString(s2, "")

	if o1 != o2 {
		return false
	}

	return true
}
