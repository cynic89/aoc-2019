package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readRange(input string) (uint64, uint64, error) {
	nos := strings.Split(input, "-")
	low, err := strconv.Atoi(nos[0])
	if err != nil {
		return 0, 0, err
	}

	high, _ := strconv.Atoi(nos[1])
	if err != nil {
		return 0, 0, err
	}
	return uint64(low), uint64(high), nil
}

func possiblePasswordCount(low, high uint64) int {
	var count int
	for password := low; password <= high; password++ {
		if isSixDigits(password) && repeatedAdjacentDigits(password) && increasingDigits(password) {
			count++
		}
	}
	return count
}

func possiblePasswordCountPart2(low, high uint64) int {
	var count int
	for password := low; password <= high; password++ {
		if isSixDigits(password) && repeatedOnlyTwiceAdjacentDigits(password) && increasingDigits(password) {
			count++
		}
	}
	return count
}

func isSixDigits(n uint64) bool {
	return n >= 100000 && n <= 999999
}

func repeatedAdjacentDigits(n uint64) bool {
	var prevRem uint64 = 11
	for n > 0 {
		rem := n % 10
		if rem == prevRem {
			return true
		}
		prevRem = rem
		n = n / 10
	}
	return false
}

func repeatedOnlyTwiceAdjacentDigits(n uint64) bool {
	var (
		prevRem     uint64 = 11
		repeatCount uint64 = 1
	)
	for n > 0 {
		rem := n % 10
		if rem == prevRem {
			repeatCount++
		} else {
			if repeatCount == 2 {
				return true
			}
			repeatCount = 1
		}
		prevRem = rem
		n = n / 10
	}
	if repeatCount == 2 {
		return true
	}
	return false
}

func increasingDigits(n uint64) bool {
	var prevRem uint64 = 11
	for n > 0 {
		rem := n % 10
		if prevRem < rem {
			return false
		}
		prevRem = rem
		n = n / 10
	}
	return true
}

func main() {
	low, high, err := readRange("353096-843212")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	countPart1 := possiblePasswordCount(low, high)
	fmt.Println(countPart1)

	countPart2 := possiblePasswordCountPart2(low, high)
	fmt.Println(countPart2)
}
