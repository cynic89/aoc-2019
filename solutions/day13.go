package main

import (
	"fmt"
	"github.com/cynic89/aoc-2019/util"
	"os"
)

const (
	INPUT_FILE_DAY_13       = "inputs/day13"
	INPUT_FILE_DAY_13_PART2 = "inputs/day13_part2"
	EMPTY_TILE              = 0
	WALL_TILE               = 1
	BLOCK_TILE              = 2
	PADDLE                  = 3
	BALL                    = 4
)

func arcadeBoard(outputs []int64) (board map[util.Coordinates]int64, boundary util.Coordinates, ballPos util.Coordinates) {
	var i int
	board = make(map[util.Coordinates]int64)
	for {
		x := int(outputs[i])
		y := int(outputs[i+1])
		tile := outputs[i+2]
		board[util.Coordinates{x, y}] = tile
		i = i + 3
		if x > boundary.X {
			boundary.X = x
		}

		if y > boundary.Y {
			boundary.Y = y
		}

		if tile == BALL {
			ballPos = util.Coordinates{x, y}
		}

		if i == len(outputs) {
			break
		}

	}
	return board, boundary, ballPos
}

func totalBlocksCount(board map[util.Coordinates]int64) (blocks int) {
	for _, v := range board {
		if v == BLOCK_TILE {
			blocks++
		}
	}
	return blocks
}

func draw(board map[util.Coordinates]int64, bounds util.Coordinates) {
	for i := 0; i <= bounds.X; i++ {
		for j := 0; j <= bounds.Y; j++ {
			if board[util.Coordinates{i, j}] == WALL_TILE {
				fmt.Print("W")
			}

			if board[util.Coordinates{i, j}] == BLOCK_TILE {
				fmt.Print("B")
			}

			if board[util.Coordinates{i, j}] == PADDLE {
				fmt.Print("P")
			}

			if board[util.Coordinates{i, j}] == BALL {
				fmt.Print("O")
			}

			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func play(board map[util.Coordinates]int64, bounds, start util.Coordinates) (blocks int) {
	currentPos := util.Coordinates{start.X - 1, start.Y - 1}
	dir := "upleft"
	for {

		if currentPos.X > bounds.X {
			break
		}

		tile := board[util.Coordinates{currentPos.X, currentPos.Y}]
		switch tile {
		case EMPTY_TILE:
		case WALL_TILE:
			if currentPos.Y == 0 || currentPos.Y == bounds.Y {
				dir = boundReflect(dir)
			} else {
				dir = sideReflect(dir)
			}
		case BLOCK_TILE:
			blocks++
		case PADDLE:
			if currentPos.Y == 0 || currentPos.Y == bounds.Y {
				dir = boundReflect(dir)
			} else {
				dir = sideReflect(dir)
			}
		}
		currentPos = nextLoc(currentPos, dir)
	}
	return blocks
}

func sideReflect(dir string) string {
	switch dir {
	case "downright":
		return "downleft"
	case "downleft":
		return "downright"
	case "upleft":
		return "upright"
	case "upright":
		return "upleft"
	}

	return dir

}

func boundReflect(dir string) string {
	switch dir {
	case "downright":
		return "upleft"
	case "downleft":
		return "upright"
	case "upleft":
		return "downright"
	case "upright":
		return "downleft"
	}

	return dir

}

func nextLoc(currentLoc util.Coordinates, dir string) (nextLoc util.Coordinates) {

	if dir == "downright" {
		nextLoc = util.Coordinates{currentLoc.X + 1, currentLoc.Y + 1}

	}
	if dir == "downleft" {
		nextLoc = util.Coordinates{currentLoc.X - 1, currentLoc.Y + 1}
	}

	if dir == "upleft" {
		nextLoc = util.Coordinates{currentLoc.X - 1, currentLoc.Y - 1}
	}

	if dir == "upright" {
		nextLoc = util.Coordinates{currentLoc.X + 1, currentLoc.Y - 1}
	}
	return nextLoc
}

func playAI(p *program) {
	var nextInput int64
	var ballLoc, paddleLoc util.Coordinates
	p.run(input{noun: p.intCodes[1], verb: p.intCodes[2], auto: true})
	//_, _, _ := arcadeBoard(p.allOutputs)
	for !p.complete {
		ballLoc = getTileLoc(p.lastRunOuputs, BALL, ballLoc)
		paddleLoc = getTileLoc(p.lastRunOuputs, PADDLE, paddleLoc)
		if paddleLoc.X > ballLoc.X {
			nextInput = -1
		}

		if paddleLoc.X < ballLoc.X {
			nextInput = 1
		}

		if paddleLoc.X == ballLoc.X {
			nextInput = 0
		}
		p.run(input{noun: p.intCodes[1], verb: p.intCodes[2], prompt: []int64{nextInput}, auto: true})
	}
}

func getTileLoc(outputs []int64, tileType int, defaultLoc util.Coordinates) (tileLoc util.Coordinates) {
	var i int
	for i < len(outputs) {
		x := int(outputs[i])
		y := int(outputs[i+1])
		tile := outputs[i+2]
		if int(tile) == tileType {
			tileLoc = util.Coordinates{x, y}
			return tileLoc
		}
		i = i + 3
	}
	tileLoc = defaultLoc
	return defaultLoc
}

func getPaddleCoordinates(outputs []int64, c util.Coordinates) util.Coordinates {
	fmt.Println(outputs)
	var i int
	for {
		if i == len(outputs) {
			break
		}
		x := int(outputs[i])
		y := int(outputs[i+1])
		tile := outputs[i+2]
		if tile == PADDLE {
			return util.Coordinates{x, y}
		}
		i = i + 3
	}
	return c
}

func day13() {
	fmt.Println("Running Day 13")

	prog, err := readProgram(INPUT_FILE_DAY_13)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	prog.run(input{noun: prog.intCodes[1], verb: prog.intCodes[2]})

	board, _, _ := arcadeBoard(prog.allOutputs)
	//fmt.Println(bounds)
	//draw(board, bounds)

	blocks := totalBlocksCount(board)
	fmt.Println(blocks)

	prog2, err := readProgram(INPUT_FILE_DAY_13_PART2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	playAI(&prog2)
	fmt.Println(prog2.output)

}
