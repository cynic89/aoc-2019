package main

import (
	"fmt"
	"github.com/cynic89/aoc-2019/util"
	"os"
)

const (
	INPUT_FILE_DAY_19 = "inputs/day19"
)

func affectedPoints(p *program, size util.Coordinates) (count int) {
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			status := getDroneStatus(p, x, y)
			if status == 1 {
				count++
			}
		}
	}
	return count
}

func squareOrigin(p *program, size util.Coordinates) util.Coordinates {
	var x, y int
	for {
		if getDroneStatus(p, x, y) == 0 {
			x++
		} else {
			if getDroneStatus(p, x+size.X-1, y) == 1 {
				if getDroneStatus(p, x, y+size.Y-1) == 1 {
					return util.Coordinates{x, y}
				} else {
					x++
					continue
				}
			} else {
				x = 0
				y++
			}
		}
	}
}

func getDroneStatus(p *program, x, y int) int {
	p.reset()
	p.run(input{noun: p.intCodes[1], verb: p.intCodes[2], auto: true, prompt: []int64{int64(x), int64(y)}})
	return int(p.output)
}

func day19() {
	fmt.Println("Running Day 19")
	p, err := readProgram(INPUT_FILE_DAY_19)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c := affectedPoints(&p, util.Coordinates{50, 50})
	fmt.Println(c)

	o := squareOrigin(&p, util.Coordinates{100, 100})
	fmt.Println(o)

}
