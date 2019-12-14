package main

import (
	"fmt"
	"os"
)

const (
	INPUT_FILE_DAY_2 = "inputs/day2"
)

func day2() {
	fmt.Println("Running Day 2 ")
	prog, err := readProgram(INPUT_FILE_DAY_2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	prog.run(input{noun: 12, verb: 2})
	resultPart1 := prog.intCodes[0]
	fmt.Println(resultPart1)

	resultPart2 := tryUntil(prog, 19690720)
	fmt.Println(resultPart2)
}
