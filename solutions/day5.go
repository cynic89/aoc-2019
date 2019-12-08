package main

import (
	"fmt"
	"os"
)

const (
	INPUT_FILE_5 = "inputs/day5"
)

func day5() {
	fmt.Println("Running Day 5")

	prog, err := readProgram(INPUT_FILE_5)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//part 1
	prog.run(input{noun: prog.intCodes[1], verb: prog.intCodes[2], prompt: []int64{1}})
	fmt.Println(prog.output)
	//part 2
	prog.reset()
	prog.run(input{noun: prog.intCodes[1], verb: prog.intCodes[2], prompt: []int64{5}})
	fmt.Println(prog.output)
}
