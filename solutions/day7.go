package main

import (
	"fmt"
	"os"
)

const (
	INPUT_FILE_7 = "inputs/day7"
)

func day7() {
	fmt.Println("Running Day 7")

	prog, err := readProgram(INPUT_FILE_7)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result := maxThrustSignal(prog, 0, 4)
	fmt.Println(result)

	result = maxThrustSignal(prog, 5, 9)
	fmt.Println(result)
}

func maxThrustSignal(prog program, start, end int64) int64 {
	var maxThrustSignal int64 = 0
	for a := start; a <= end; a++ {
		for b := start; b <= end; b++ {
			for c := start; c <= end; c++ {
				for d := start; d <= end; d++ {
					for e := start; e <= end; e++ {
						currentSignal := thrustSignal(prog, a, b, c, d, e)
						if currentSignal > maxThrustSignal && !isSettingsDuplicated([]int64{a, b, c, d, e}) {
							maxThrustSignal = currentSignal
						}
					}
				}
			}
		}
	}

	return maxThrustSignal
}

func thrustSignal(prog program, a, b, c, d, e int64) int64 {
	amplifiers := getAmplifiers(prog)
	var inputs = []int64{a, b, c, d, e}
	for idx, amplifier := range amplifiers {
		amplifier.run(input{noun: amplifier.intCodes[1], verb: amplifier.intCodes[2],
			prompt: []int64{inputs[idx],}, auto: true})
	}
	for !allAmplifiersComplete(amplifiers) {
		for idx, amplifier := range amplifiers {
			var op int
			if idx == 0 {
				op = 4
			} else {
				op = idx - 1
			}
			amplifier.run(input{noun: amplifier.intCodes[1], verb: amplifier.intCodes[2],
				prompt: []int64{amplifiers[op].output}, auto: true})
		}
	}
	return amplifiers[4].output

}

func allAmplifiersComplete(amplifiers []*program) bool {
	for _, a := range amplifiers {
		if !a.complete {
			return false
		}
	}
	return true
}

func getAmplifiers(p program) []*program {
	var amplifiers []*program
	var intCodesOriginal []int64
	intCodesOriginal = p.intCodesOriginal
	for i := 0; i < 5; i++ {
		prog := new(program)
		prog.intCodes = append(prog.intCodes, intCodesOriginal...)
		prog.intCodesOriginal = intCodesOriginal
		amplifiers = append(amplifiers, prog)
	}
	return amplifiers
}

func isSettingsDuplicated(settings []int64) bool {
	cache := make(map[int64]bool)
	for _, d := range settings {
		_, ok := cache[d]
		if ok {
			return true
		}
		cache[d] = true
	}
	return false
}

func toPromptString(seq, op int) []string {
	return []string{fmt.Sprintf("%d", seq), fmt.Sprintf("%d", op)}
}
