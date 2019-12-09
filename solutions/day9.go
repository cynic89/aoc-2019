package main

import (
	"fmt"
	"os"
)

const (
	INPUT_FILE_DAY_9 = "inputs/day9"
)

func day9() {
	fmt.Println("Running Day 9")

	prog, err := readProgram(INPUT_FILE_DAY_9)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	prog.run(input{noun: prog.intCodes[1], verb: prog.intCodes[2]})
	printOuputs(prog)

}

func printOuputs(p program) {
	for i, op := range p.allOutputs {
		fmt.Printf("%d", op)
		if i < len(p.allOutputs)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
