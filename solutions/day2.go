package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	INPUT_FILE_2        = "inputs/day2"
	ADD          uint64 = 1
	MUL          uint64 = 2
	HALT         uint64 = 99
)

type program struct {
	intCodes         []uint64
	intCodesOriginal []uint64
	pos              int
}

func readProgram() (program, error) {
	f, err := os.Open(INPUT_FILE_2)
	if err != nil {
		return program{}, err
	}

	defer f.Close()

	var (
		intCodes         []uint64
		intcodesOriginal []uint64
	)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		vals := strings.Split(text, ",")
		for _, val := range vals {
			valInt, err := strconv.ParseUint(val, 10, 32)
			if err != nil {
				return program{}, err
			}
			intCodes = append(intCodes, valInt)
			intcodesOriginal = append(intcodesOriginal, valInt)
		}
		break
	}
	return program{intCodes: intCodes, intCodesOriginal: intcodesOriginal}, nil

}

func (p program) run(noun, verb uint64) uint64 {

	defer func() { p.reset() }()

	p.intCodes[1] = noun
	p.intCodes[2] = verb
	for {
		if p.getCurrentOpcode() == HALT {
			return p.intCodes[0]
		}

		if p.getCurrentOpcode() == ADD {
			result := p.getLeftOperand() + p.getRightOperand()
			p.updateResult(result)
		}

		if p.getCurrentOpcode() == MUL {
			result := p.getLeftOperand() * p.getRightOperand()
			p.updateResult(result)
		}

		p.advance()
	}
}

func (p program) tryUntil(match uint64) int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			result := p.run(uint64(noun), uint64(verb))
			if result == match {
				return 100*noun + verb
			}
		}
	}

	return -1
}

func (p *program) updateResult(r uint64) {
	pos := p.intCodes[p.pos+3]
	p.intCodes[pos] = r
}

func (p *program) getLeftOperand() uint64 {
	pos := p.intCodes[p.pos+1]
	return p.intCodes[pos]
}

func (p *program) getRightOperand() uint64 {
	pos := p.intCodes[p.pos+2]
	return p.intCodes[pos]
}

func (p *program) getCurrentOpcode() uint64 {
	return p.intCodes[p.pos]
}

func (p *program) advance() {
	p.pos = p.pos + 4
}

func (p *program) reset() {
	copy(p.intCodes, p.intCodesOriginal)
}

func main() {
	prog, err := readProgram()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	resultPart1 := prog.run(12, 2)
	fmt.Println(resultPart1)

	resultPart2 := prog.tryUntil(19690720)
	fmt.Println(resultPart2)
}
