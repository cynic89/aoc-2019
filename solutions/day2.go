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
	intCodes          []int64
	intCodesOriginal  []int64
	pos               int
	currentOpCode     int
	currentParamTypes [2]int
	complete          bool
	output            int64
	outputCalculated  bool
}

type input struct {
	noun   int64
	verb   int64
	prompt []int64
	auto   bool
}

func readProgram(file string) (program, error) {
	f, err := os.Open(file)
	if err != nil {
		return program{}, err
	}

	defer f.Close()

	var (
		intCodes         []int64
		intcodesOriginal []int64
	)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		vals := strings.Split(text, ",")
		for _, val := range vals {
			valInt, err := strconv.ParseInt(val, 10, 64)
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

func (p *program) run(input input) {
	var (
		inputCount int
	)

	p.intCodes[1] = input.noun
	p.intCodes[2] = input.verb

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
			result, err := p.handleInputOpcode(input.prompt, inputCount, input.auto)
			if err != nil {
				return
			}
			p.updateResult(result)
			inputCount++
		}

		if p.currentOpCode == OUTPUT {
			p.updateResult(-1)
		}

		if p.currentOpCode == LESS_THAN {
			var result int64
			if p.getLeftOperand(p.currentParamTypes[0]) < p.getRightOperand(p.currentParamTypes[1]) {
				result = 1
			}
			p.updateResult(result)
		}

		if p.currentOpCode == EQUALS {
			var result int64
			if p.getLeftOperand(p.currentParamTypes[0]) == p.getRightOperand(p.currentParamTypes[1]) {
				result = 1
			}
			p.updateResult(result)
		}

		if p.currentOpCode == HALT {
			p.complete = true
			return
		}

		p.advance()

	}
}

func (p *program) handleInputOpcode(prompt [] int64, index int, auto bool) (int64, error) {
	var inputText string
	if ! (len(prompt) > index) {
		if auto {
			return -1, fmt.Errorf("Input should be auto")
		}
		fmt.Println("Enter Code")
		reader := bufio.NewReader(os.Stdin)
		inputText, _ = reader.ReadString('\n')
	} else {
		return prompt[index], nil
	}
	code, err := strconv.ParseInt(strings.Trim(inputText, "\n"), 10, 64)
	panic(err)
	//fmt.Printf("Input = %d ", code)
	return code, nil

}

func (p *program) updateResult(result int64) (pos int64) {
	if p.currentOpCode == ADD || p.currentOpCode == MUL || p.currentOpCode == LESS_THAN || p.currentOpCode == EQUALS {
		pos = p.intCodes[p.pos+3]
		p.intCodes[pos] = result
	}

	if p.currentOpCode == INPUT {
		pos = p.intCodes[p.pos+1]
		p.intCodes[pos] = result
	}

	if p.currentOpCode == OUTPUT {
		pos = p.intCodes[p.pos+1]
		p.output = p.intCodes[pos]
	}
	return pos
}

func (p *program) getLeftOperand(parameterType int) int64 {
	if parameterType == 0 {
		pos := p.intCodes[p.pos+1]
		return p.intCodes[pos]
	}
	return p.intCodes[p.pos+1]
}

func (p *program) getRightOperand(parameterType int) int64 {
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
			p.pos = int(rightOperand)
		} else {
			p.pos = p.pos + 3
		}
	}

	if p.currentOpCode == JUMP_IF_FALSE {
		leftOperand := p.getLeftOperand(p.currentParamTypes[0])
		rightOperand := p.getRightOperand(p.currentParamTypes[1])
		if leftOperand == 0 {
			p.pos = int(rightOperand)
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
	p.output = 0
	p.outputCalculated = false
	p.pos = 0
	p.currentOpCode = 0
	p.complete = false
}

func tryUntil(p program, match int64) int64 {
	var intCodesOriginal []int64
	intCodesOriginal = p.intCodesOriginal
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			prog := new(program)
			prog.intCodes = append(prog.intCodes, intCodesOriginal...)
			prog.intCodesOriginal = intCodesOriginal
			prog.run(input{noun: int64(noun), verb: int64(verb)})
			result := prog.intCodes[0]
			if result == match {
				return int64(100*noun + verb)
			}
		}
	}

	return -1
}

func day2() {
	fmt.Println("Running Day 2 ")
	prog, err := readProgram(INPUT_FILE_2)
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
