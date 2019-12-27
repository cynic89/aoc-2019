package main

import (
	"fmt"
	"github.com/cynic89/aoc-2019/util"
)

const (
	INPUT_FILE_DAY_15 = "inputs/day15"
	WALL              = 0
	FREE              = 1
	OXYGEN            = 2
	NORTH             = 1
	SOUTH             = 2
	WEST              = 3
	EAST              = 4
	WALL_ID           = 10
)

type directedCoord struct {
	util.Coordinates
	dir int
}

var canvas = make(map[util.Coordinates]int)

type route []directedCoord

func NewDirCoord(x, y, dir int) directedCoord {
	return directedCoord{Coordinates: util.Coordinates{x, y}, dir: dir,}
}

func traverse(p *program) {
	visited := make(map[util.Coordinates]bool)
	var queue []route
	initial := NewDirCoord(0, 0, NORTH)
	queue = append(queue, []directedCoord{initial})
	count := 0
	//oxygen := false
	for len(queue) > 0 {
		curr := queue[0]

		lastPos := curr[len(curr)-1]
		if visited[lastPos.Coordinates] {
			count++
			queue = queue[1:]
			continue
		}
		visited[lastPos.Coordinates] = true

		p.reset()
		simulateMove(p, curr[1:])

		queue = tryMove(p, queue, curr, lastPos, NORTH)

		queue = tryMove(p, queue, curr, lastPos, SOUTH)

		queue = tryMove(p, queue, curr, lastPos, WEST)

		queue = tryMove(p, queue, curr, lastPos, EAST)

		queue = queue[1:]
		count++
	}
}

func simulateMove(p *program, curr route) {
	for _, pos := range curr {
		runCode(p, pos.dir)
	}

}

func tryMove(p *program, queue []route, curr route, lastPos directedCoord, dir int64) []route {
	op := runCode(p, int(dir))
	var newPos directedCoord
	switch dir {
	case NORTH:
		newPos = NewDirCoord(lastPos.X, lastPos.Y+1, NORTH)
	case SOUTH:
		newPos = NewDirCoord(lastPos.X, lastPos.Y-1, SOUTH)
	case EAST:
		newPos = NewDirCoord(lastPos.X+1, lastPos.Y, EAST)
	case WEST:
		newPos = NewDirCoord(lastPos.X-1, lastPos.Y, WEST)

	}
	if op == FREE {
		var newRoute = make([]directedCoord, len(curr), len(curr))
		copy(newRoute, curr)
		newRoute = append(newRoute, newPos)
		queue = append(queue, newRoute)
		runCode(p, opposite(int(dir)))
	}

	if op == WALL {
		canvas[newPos.Coordinates] = WALL_ID
	}

	if op == OXYGEN {
		fmt.Printf("Found Oxygen at %v\nNumber of steps = %d\n", newPos, len(curr))
		canvas[newPos.Coordinates] = OXYGEN
	}

	return queue
}

func opposite(dir int) int {
	if dir == NORTH {
		return SOUTH
	}
	if dir == SOUTH {
		return NORTH
	}
	if dir == EAST {
		return WEST
	}
	if dir == WEST {
		return EAST
	}

	panic(100)
}

func (lastCoord directedCoord) getNewDirDirect(toMove int) int {
	return toMove
}

func (lastCoord directedCoord) getNewDir(toMove int) int {
	lastDir := lastCoord.dir

	if (lastDir == NORTH && toMove == NORTH) || (lastDir == WEST && toMove == EAST) ||
		(lastDir == EAST && toMove == WEST) || (lastDir == SOUTH && toMove == SOUTH) {
		return NORTH
	}

	if (lastDir == SOUTH && toMove == NORTH) || (lastDir == WEST && toMove == WEST) ||
		(lastDir == EAST && toMove == EAST) || (lastDir == NORTH && toMove == SOUTH) {
		return SOUTH
	}

	if (lastDir == SOUTH && toMove == WEST) || (lastDir == NORTH && toMove == EAST) ||
		(lastDir == EAST && toMove == NORTH) || (lastDir == WEST && toMove == SOUTH) {
		return EAST
	}

	if (lastDir == NORTH && toMove == WEST) || (lastDir == SOUTH && toMove == EAST) ||
		(lastDir == WEST && toMove == NORTH) || (lastDir == EAST && toMove == SOUTH) {
		return WEST
	}
	panic(100)
}

func runCode(p *program, dir int) int64 {
	p.run(input{noun: p.intCodes[1], verb: p.intCodes[2], prompt: []int64{int64(dir)}, auto: true})
	return p.output
}

func spreadOxygen() int {
	min := 0
	var oxygenCoordinates []util.Coordinates
	oxygenCoordinates = append(oxygenCoordinates, util.Coordinates{-16, -14})
	for {
		var newOxygenCoordinates []util.Coordinates
		for _, v := range oxygenCoordinates {
			if tryFill(v.X+1, v.Y) {
				newOxygenCoordinates = append(newOxygenCoordinates, util.Coordinates{v.X + 1, v.Y})
			}
			if tryFill(v.X-1, v.Y) {
				newOxygenCoordinates = append(newOxygenCoordinates, util.Coordinates{v.X - 1, v.Y})
			}
			if tryFill(v.X, v.Y+1) {
				newOxygenCoordinates = append(newOxygenCoordinates, util.Coordinates{v.X, v.Y + 1})
			}
			if tryFill(v.X, v.Y-1) {
				newOxygenCoordinates = append(newOxygenCoordinates, util.Coordinates{v.X, v.Y - 1})
			}

		}

		if len(newOxygenCoordinates) == 0 {
			return min
		}

		oxygenCoordinates = newOxygenCoordinates

		min++
	}
}

func tryFill(x, y int) bool {
	e := canvas[util.Coordinates{x, y}]
	if e != WALL_ID && e != OXYGEN {
		canvas[util.Coordinates{x, y}] = OXYGEN
		return true
	}
	return false
}

func fillCanvas() {
	var result string
	for y := 25; y >= -20; y-- {
		for x := -25; x <= 20; x++ {
			if x == 0 && y == 0 {
				result += "+"
				continue
			}
			e := canvas[util.Coordinates{x, y}]
			if e == WALL_ID {
				result = result + "#"
			} else if e == OXYGEN {
				result = result + "O"
			} else {
				result = result + "."
			}
		}
		result = result + "\n"
	}

	fmt.Print(result)
}

func day15() {
	fmt.Println("Running Day 15")
	prog, _ := readProgram(INPUT_FILE_DAY_15)
	traverse(&prog)
	fmt.Println(spreadOxygen())
	fillCanvas()
}
