package util

import "fmt"

type nomad struct {
	initial    Coordinates
	facing     Dir
	size       Coordinates
	currentPos Coordinates
}

type Dir string

var (
	Up    Dir = "Up"
	Down  Dir = "Down"
	Right Dir = "Right"
	Left  Dir = "Left"
)

func NewNomad(size Coordinates, opts ...func(*nomad)) *nomad {
	n := new(nomad)
	n.size = size
	n.facing = Up

	for _, o := range opts {
		o(n)
	}
	n.currentPos = n.initial
	return n
}

func Initial(p Coordinates) func(*nomad) {
	return func(n *nomad) { n.initial = p }
}

func Facing(f Dir) func(*nomad) {
	return func(n *nomad) { n.facing = f }
}

func (n *nomad) Advance() (Coordinates, error) {
	return n.Move(1, Up, false)
}

func (n *nomad) AdvancePeek() (Coordinates, error) {
	return n.Move(1, Up, true)
}

func (n *nomad) Retreat() (Coordinates, error) {
	return n.Move(1, Down, false)
}

func (n *nomad) Move(steps int, dirToMove Dir, peek bool) (Coordinates, error) {
	var newPos Coordinates
	var newDir Dir
	cDir := n.facing

	if (cDir == Up && dirToMove == Up) || (cDir == Left && dirToMove == Right) ||
		(cDir == Right && dirToMove == Left) || (cDir == Down && dirToMove == Down) {
		newDir = Up
		newPos = Coordinates{n.currentPos.X, n.currentPos.Y - steps}
	}

	if (cDir == Down && dirToMove == Up) || (cDir == Left && dirToMove == Left) ||
		(cDir == Right && dirToMove == Right) || (cDir == Up && dirToMove == Down) {
		newDir = Down
		newPos = Coordinates{n.currentPos.X, n.currentPos.Y + steps}
	}

	if (cDir == Down && dirToMove == Left) || (cDir == Up && dirToMove == Right) ||
		(cDir == Right && dirToMove == Up) || (cDir == Left && dirToMove == Down) {
		newDir = Right
		newPos = Coordinates{n.currentPos.X + steps, n.currentPos.Y}
	}

	if (cDir == Up && dirToMove == Left) || (cDir == Down && dirToMove == Right) ||
		(cDir == Left && dirToMove == Up) || (cDir == Right && dirToMove == Down) {
		newDir = Left
		newPos = Coordinates{n.currentPos.X - steps, n.currentPos.Y}
	}

	if !(newPos.X < n.size.X) || !(newPos.Y < n.size.Y) {
		return Coordinates{}, fmt.Errorf("Out of Bounds")
	}
	if !peek {
		n.currentPos.X = newPos.X
		n.currentPos.Y = newPos.Y
		n.facing = newDir
	}

	return newPos, nil
}

func (n *nomad) GetCurrentPos() (Coordinates, Dir) {
	return n.currentPos, n.facing
}
