package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	INPUT_FILE_2      = "inputs/day2"
	ADD           int = 1
	MUL           int = 2
	INPUT         int = 3
	OUTPUT        int = 4
	JUMP_IF_TRUE  int = 5
	JUMP_IF_FALSE int = 6
	LESS_THAN     int = 7
	EQUALS        int = 8
	HALT          int = 99
)

type program struct {
	intCodes          []int
	intCodesOriginal  []int
	pos               int
	currentOpCode     int
	currentParamTypes [2]int
}

func readProgram(file string) (program, error) {
	f, err := os.Open(file)
	if err != nil {
		return program{}, err
	}

	defer f.Close()

	var (
		intCodes         []int
		intcodesOriginal []int
	)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		vals := strings.Split(text, ",")
		for _, val := range vals {
			valInt, err := strconv.Atoi(val)
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

func (p program) run(noun, verb int) int {

	defer func() { p.reset() }()

	p.intCodes[1] = noun
	p.intCodes[2] = verb
	for {

		p.parseOpcode()

		if p.currentOpCode == ADD {
			result := p.getLeftOperand(p.currentParamTypes[0]) + p.getRightOperand(p.currentParamTypes[1])
			p.updateResult(result)
		}

		if p.currentOpCode == MUL {
			result := p.getLeftOperand(p.currentParamTypes[0]) * p.getRightOperand(p.currentParamTypes[1])
			p.updateResult(result)
		}

		if p.currentOpCode == INPUT {
			result := p.handleInputOpcode()
			p.updateResult(result)
		}

		if p.currentOpCode == OUTPUT {
			p.updateResult(-1)
		}

		if p.currentOpCode == LESS_THAN {
			var result int
			if p.getLeftOperand(p.currentParamTypes[0]) < p.getRightOperand(p.currentParamTypes[1]) {
				result = 1
			}
			p.updateResult(result)
		}

		if p.currentOpCode == EQUALS {
			var result int
			if p.getLeftOperand(p.currentParamTypes[0]) == p.getRightOperand(p.currentParamTypes[1]) {
				result = 1
			}
			p.updateResult(result)
		}

		if p.currentOpCode == HALT {
			return p.intCodes[0]
		}

		p.advance()

	}
}

func (p *program) handleInputOpcode() int {
	fmt.Println("Enter Code")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	code, _ := strconv.Atoi(strings.Trim(input, "\n"))
	return code

}

func (p program) tryUntil(match int) int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			result := p.run(int(noun), int(verb))
			if result == match {
				return 100*noun + verb
			}
		}
	}

	return -1
}

func (p *program) updateResult(result int) {
	if p.currentOpCode == ADD || p.currentOpCode == MUL || p.currentOpCode == LESS_THAN || p.currentOpCode == EQUALS {
		pos := p.intCodes[p.pos+3]
		p.intCodes[pos] = result
	}

	if p.currentOpCode == INPUT {
		pos := p.intCodes[p.pos+1]
		p.intCodes[pos] = result
	}

	if p.currentOpCode == OUTPUT {
		pos := p.intCodes[p.pos+1]
		fmt.Println(p.intCodes[pos])
	}

}

func (p *program) getLeftOperand(parameterType int) int {
	if parameterType == 0 {
		pos := p.intCodes[p.pos+1]
		return p.intCodes[pos]
	}
	return p.intCodes[p.pos+1]
}

func (p *program) getRightOperand(parameterType int) int {
	if parameterType == 0 {
		pos := p.intCodes[p.pos+2]
		return p.intCodes[pos]
	}
	return p.intCodes[p.pos+2]
}

func (p *program) parseOpcode() {
	intCode := p.intCodes[p.pos]
	paddedIntCode := fmt.Sprintf("000%d", intCode)
	p.currentOpCode, _ = strconv.Atoi(paddedIntCode[len(paddedIntCode)-2:])
	p.currentParamTypes[0], _ = strconv.Atoi(paddedIntCode[len(paddedIntCode)-3 : len(paddedIntCode)-2])
	p.currentParamTypes[1], _ = strconv.Atoi(paddedIntCode[len(paddedIntCode)-4 : len(paddedIntCode)-3])
}

func (p *program) advance() {
	if p.currentOpCode == ADD || p.currentOpCode == MUL {
		p.pos = p.pos + 4
	}

	if p.currentOpCode == INPUT || p.currentOpCode == OUTPUT {
		p.pos = p.pos + 2
	}

	if p.currentOpCode == JUMP_IF_TRUE {
		leftOperand := p.getLeftOperand(p.currentParamTypes[0])
		rightOperand := p.getRightOperand(p.currentParamTypes[1])
		if leftOperand != 0 {
			p.pos = rightOperand
		} else {
			p.pos = p.pos + 3
		}
	}

	if p.currentOpCode == JUMP_IF_FALSE {
		leftOperand := p.getLeftOperand(p.currentParamTypes[0])
		rightOperand := p.getRightOperand(p.currentParamTypes[1])
		if leftOperand == 0 {
			p.pos = rightOperand
		} else {
			p.pos = p.pos + 3
		}
	}

	if p.currentOpCode == LESS_THAN || p.currentOpCode == EQUALS {
		p.pos = p.pos + 4
	}

}

func (p *program) reset() {
	copy(p.intCodes, p.intCodesOriginal)
}

func day2() {
	fmt.Println("Running Day 1 Solution")
	prog, err := readProgram(INPUT_FILE_2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	resultPart1 := prog.run(12, 2)
	fmt.Println(resultPart1)

	resultPart2 := prog.tryUntil(19690720)
	fmt.Println(resultPart2)
}
