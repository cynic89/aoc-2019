package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	INPUT_FILE = "inputs/day1"
)

func readInputs() ([]uint64, error) {
	f, err := os.Open(INPUT_FILE)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var inputs []uint64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		numText := scanner.Text()
		num, err := strconv.ParseUint(numText, 10, 32)
		if err != nil {
			return nil, err
		}

		inputs = append(inputs, num)

	}

	return inputs, nil

}

func fuelPart1(masses []uint64) (fuel uint64) {
	for _, mass := range masses {
		fuel = fuel + (mass/3 - 2)
	}
	return fuel
}

func fuelPart2(masses []uint64) (fuel uint64) {
	for _, mass := range masses {
		fuel = fuel + calcFuel(mass)
	}
	return fuel
}

func calcFuel(mass uint64) (totalFuel uint64) {
	var fuel int
	for {
		fuel = int(mass/3 - 2)
		if fuel <= 0 {
			break
		} else {
			totalFuel = totalFuel + uint64(fuel)
		}
		mass = uint64(fuel)
	}
	return totalFuel
}

func main() {
	masses, err := readInputs()
	if err != nil {
		fmt.Println(err)
	}

	fuelPart1 := fuelPart1(masses)
	fmt.Println(fuelPart1)

	fuelPart2 := fuelPart2(masses)
	fmt.Println(fuelPart2)
}
