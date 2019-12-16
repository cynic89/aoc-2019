package main

import (
	"bufio"
	"fmt"
	"github.com/cynic89/aoc-2019/util"
	"os"
)

const (
	INPUT_FILE_DAY_16 = "inputs/day16"
)

func readSignals() ([]int, error) {
	f, err := os.Open(INPUT_FILE_DAY_16)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var signals []int
	for scanner.Scan() {
		text := scanner.Text()
		for _, c := range text {
			signals = append(signals, int(c)-48)
		}
	}
	return signals, nil

}

func readSignalsPart2() ([]int, error) {
	f, err := os.Open(INPUT_FILE_DAY_16)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var signals []int
	for scanner.Scan() {
		text := scanner.Text()
		for _, c := range text {
			signals = append(signals, int(c)-48)
		}
	}

	for i := 0; i < 10000; i++ {
		signals = append(signals, signals...)
		fmt.Println(i)
	}

	return signals, nil

}

func pattern(digit int, size int) []int {
	var pattern []int
	var initialPattern = []int{0, 1, 0, -1}
	var toAppend []int
	for _, p := range initialPattern {
		for i := 0; i < digit; i++ {
			toAppend = append(toAppend, p)
		}
	}

	for len(pattern) <= size+1 {
		pattern = append(pattern, toAppend...)
	}

	return pattern
}

func output(input []int, noOfDigits int) (results []int) {

	for digit := 1; digit <= noOfDigits; digit++ {
		var result int
		pattern := pattern(digit, len(input))
		for i, ip := range input {
			result += ip * pattern[i+1]
		}
		results = append(results, util.Abs(result%10))
	}

	return results

}

func outputAfterPhases(input []int, noOfDigits, phases int) []int {
	op := input
	for i := 0; i < phases; i++ {
		op = output(op, noOfDigits)
	}

	return op

}

func part1() {
	signals, err := readSignals()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(len(signals))
	op := outputAfterPhases(signals, len(signals), 100)

	for i := 0; i < 8; i++ {
		fmt.Print(op[i])
	}
	fmt.Println()
}

func part2() {
	signals, err := readSignalsPart2()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	op := outputAfterPhases(signals, len(signals), 100)

	for i := 0; i < 8; i++ {
		fmt.Print(op[i])
	}
	fmt.Println()
}

func day16() {
	part2()

}
