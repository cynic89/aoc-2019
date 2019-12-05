package main

import (
	"fmt"
	"os"
)

const (
	INPUT_FILE_5 = "inputs/day5"
)

func day5() {
	prog, err := readProgram(INPUT_FILE_5)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	prog.run(prog.intCodes[1], prog.intCodes[2])
}
