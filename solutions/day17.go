package main

import (
	"fmt"
	"github.com/cynic89/aoc-2019/util"
	"math"
	"os"
)

const (
	INPUT_FILE_DAY_17        = "inputs/day17"
	INPUT_FILE_DAY_17_PART_2 = "inputs/day17_part2"
)

func mapLayout(p *program) map[util.Coordinates]int {
	p.run(input{noun: p.intCodes[1], verb: p.intCodes[2], prompt: nil, auto: false,})

	var x, y int
	l := make(map[util.Coordinates]int)
	for _, o := range p.allOutputs {
		if o == '\n' {
			x = 0
			y++
		} else {
			if o == '^' {
				fmt.Println(x, y)
			}
			l[util.Coordinates{x, y}] = int(o)
			x++
		}
	}
	return l
}

func alignmentParamsSum(l map[util.Coordinates]int) int {
	var sum int
	for k, _ := range l {
		if isIntersection(l, k) {
			sum += k.X * k.Y
		}
	}

	return sum
}

func isIntersection(l map[util.Coordinates]int, pos util.Coordinates) bool {
	return l[pos] == '#' && l[util.Coordinates{pos.X + 1, pos.Y}] == '#' &&
		l[util.Coordinates{pos.X - 1, pos.Y}] == '#' && l[util.Coordinates{pos.X, pos.Y + 1}] == '#' &&
		l[util.Coordinates{pos.X, pos.Y - 1}] == '#'
}

func drawLayout(l map[util.Coordinates]int) {
	size := math.Sqrt(float64(len(l)))

	for y := 0; y < int(size); y++ {
		for x := 0; x < int(size); x++ {
			fmt.Printf("%c", l[util.Coordinates{x, y}])
		}
		fmt.Println()
	}
}

type movement struct {
	dir   string
	steps int
}

func visit(l map[util.Coordinates]int) []*movement {
	n := util.NewNomad(util.Coordinates{51, 51}, util.Initial(util.Coordinates{46, 50}), util.Facing(util.Left))
	var allMoves []*movement
	allMoves = append(allMoves, &movement{"L", 0})

	for {
		currentMovement := allMoves[len(allMoves)-1]
		c, _ := n.AdvancePeek()
		if l[c] == '#' {
			currentMovement.steps = currentMovement.steps + 1
			n.Advance()
		} else {
			c, _ := n.Move(1, util.Left, true)
			if l[c] == '#' {
				allMoves = append(allMoves, &movement{"L", 1})
				n.Move(1, util.Left, false)
				continue
			}
			c, _ = n.Move(1, util.Right, true)
			if l[c] == '#' {
				allMoves = append(allMoves, &movement{"R", 1})
				n.Move(1, util.Right, false)
				continue
			}
			return allMoves

		}
	}

}

func day17() {
	fmt.Println("Running Day 17")
	p, err := readProgram(INPUT_FILE_DAY_17)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	l := mapLayout(&p)
	s := alignmentParamsSum(l)
	fmt.Println(s)
	drawLayout(l)
	visit(l)
	fmt.Println()
	mainRoutines := []int64{'A', ',', 'B', ',', 'A', ',', 'C', ',', 'C', ',', 'A', ',', 'B', ',', 'C', ',', 'B', ',', 'B', '\n'}
	a := []int64{'L', ',', 56, ',', 'R', ',', 49, 48, ',', 'L', ',', 56, ',', 'R', ',', 56, '\n'}
	b := []int64{'L', ',', 49, 50, ',', 'R', ',', 56, ',', 'R', ',', 56, '\n'}
	c := []int64{'L', ',', 56, ',', 'R', ',', 54, ',', 'R', ',', 54, ',', 'R', ',', 49, 48, ',', 'L', ',', 56, '\n'}
	vf := []int64{'n', '\n'}

	p2, err := readProgram(INPUT_FILE_DAY_17_PART_2)
	p2.run(input{noun: p.intCodes[1], verb: p.intCodes[2], auto: true, prompt: mainRoutines})
	p2.run(input{noun: p.intCodes[1], verb: p.intCodes[2], auto: true, prompt: a})
	p2.run(input{noun: p.intCodes[1], verb: p.intCodes[2], auto: true, prompt: b})
	p2.run(input{noun: p.intCodes[1], verb: p.intCodes[2], auto: true, prompt: c})
	p2.run(input{noun: p.intCodes[1], verb: p.intCodes[2], auto: true, prompt: vf})

	fmt.Println(p2.output)

}
