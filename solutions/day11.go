package main

import (
	"fmt"
	"github.com/cynic89/aoc-2019/util"
	"os"
	"sort"
)

const (
	INPUT_FILE_DAY_11 = "inputs/day11"
)

func day11() {
	fmt.Println("Running Day 11")

	prog, err := readProgram(INPUT_FILE_DAY_11)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	panels := paintPanels(&prog, 0)
	fmt.Println(len(panels))
	prog.reset()

	panels = paintPanels(&prog, 1)
	print(panels)
}

func print(panelsMap map[util.Coordinates]int64) {
	min, max := getBoundaries(panelsMap)

	fmt.Println()
	for y := max.Y; y >= min.Y; y-- {
		for x := min.X; x <= max.X; x++ {
			if panelsMap[util.Coordinates{x, y}] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func getBoundaries(panelsMap map[util.Coordinates]int64) (min, max util.Coordinates) {
	var locs []util.Coordinates
	for l, _ := range panelsMap {
		locs = append(locs, l)
	}
	sort.SliceStable(locs, func(i, j int) bool {
		return locs[i].X < locs[j].X
	})
	minX := locs[0].X
	maxX := locs[len(locs)-1].X

	sort.SliceStable(locs, func(i, j int) bool {
		return locs[i].Y < locs[j].Y
	})
	minY := locs[0].Y
	maxY := locs[len(locs)-1].Y

	min = util.Coordinates{minX, minY}
	max = util.Coordinates{maxX, maxY}
	return min, max
}

func paintPanels(p *program, start int) (panelMap map[util.Coordinates]int64) {
	point := util.Coordinates{0, 0}
	panelMap = make(map[util.Coordinates]int64)
	panelMap[point] = int64(start)
	currentDir := "up"

	for {
		p.run(input{noun: p.intCodes[1], verb: p.intCodes[2],
			prompt: []int64{panelMap[point]}, auto: true})

		if p.complete {
			break
		}

		color := p.allOutputs[len(p.allOutputs)-2]
		dir := p.allOutputs[len(p.allOutputs)-1]

		panelMap[point] = color
		if dir == 0 {
			currentDir, point = move(currentDir, "left", point)
		} else {
			currentDir, point = move(currentDir, "right", point)
		}
	}
	return panelMap
}

func move(current, towards string, loc util.Coordinates) (next string, newLoc util.Coordinates) {
	if current == "up" {
		if towards == "left" {
			next = "left"
			newLoc = left(loc)
		} else {
			next = "right"
			newLoc = right(loc)
		}
	}

	if current == "left" {
		if towards == "left" {
			next = "down"
			newLoc = down(loc)
		} else {
			next = "up"
			newLoc = up(loc)
		}
	}

	if current == "down" {
		if towards == "left" {
			next = "right"
			newLoc = right(loc)
		} else {
			next = "left"
			newLoc = left(loc)
		}
	}

	if current == "right" {
		if towards == "left" {
			next = "up"
			newLoc = up(loc)
		} else {
			next = "down"
			newLoc = down(loc)
		}
	}
	return next, newLoc
}

func left(l util.Coordinates) util.Coordinates {
	return util.Coordinates{l.X - 1, l.Y}
}

func right(l util.Coordinates) util.Coordinates {
	return util.Coordinates{l.X + 1, l.Y}
}

func up(l util.Coordinates) util.Coordinates {
	return util.Coordinates{l.X, l.Y + 1}
}
func down(l util.Coordinates) util.Coordinates {
	return util.Coordinates{l.X, l.Y - 1}
}
