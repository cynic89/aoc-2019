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
	var text string
	for scanner.Scan() {
		text = scanner.Text()
		break
	}
	fft := ""
	for i := 0; i < 2000; i++ {
		fft = fft+text
	}
	for i, c := range fft {
		signals = append(signals, int(c)-48)
		if i / 100000 >=1 && i%100000==0 {
			fmt.Println(i)
		}
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
	var initialPattern = []int{0, 1, 0, -1}
	for digit := 1; digit <= noOfDigits; digit++ {
		var result int
		for i, ip := range input {
			idx := ((i+1)/digit)%4
			result += ip * initialPattern[idx]
		}
		results = append(results, util.Abs(result%10))
	}

	return results

}

func outputAfterPhases(input []int, noOfDigits, phases int) []int {
	op := input
	for i := 0; i < phases; i++ {
		fmt.Println(i)
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

	fmt.Println()
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

	offset := len(op) / 2
	toCheck := 5972877
	var skip int
	fmt.Println(toCheck % len(op))
	c := 1
	for {
		end := len(op) * c
		start := end - offset
		if toCheck >= start && toCheck < end {
			fmt.Println("found")
			fmt.Println(start)
			skip = toCheck - start
			break
		}

		if start > toCheck {
			break
		}

		c++
	}

	fmt.Println(skip)

	for i := len(op)/2 + skip; i < len(op)/2+skip+8; i++ {
		fmt.Print(op[i])
	}

	fmt.Println()
}

func day16() {
	part1()

}
