package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ADD           int = 1
	MUL           int = 2
	INPUT         int = 3
	OUTPUT        int = 4
	JUMP_IF_TRUE  int = 5
	JUMP_IF_FALSE int = 6
	LESS_THAN     int = 7
	EQUALS        int = 8
	RELATIVE_BASE int = 9
	HALT          int = 99
)

type program struct {
	intCodes          []int64
	intCodesOriginal  []int64
	pos               int
	currentOpCode     int
	currentParamTypes [3]int
	complete          bool
	output            int64
	outputCalculated  bool
	relativeBase      int64
	allOutputs        []int64
	lastRunOuputs     []int64
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			fmt.Println(len(p.intCodes))
		}
	}()
	var (
		inputCount int
	)

	p.intCodes[1] = input.noun
	p.intCodes[2] = input.verb
	p.lastRunOuputs = []int64{}

	for {

		p.parseOpcode()

		if p.currentOpCode == ADD {
			result := *p.getFirstParameter() + *p.getSecondParameter()
			p.updateResult(result)
		}

		if p.currentOpCode == MUL {
			result := *p.getFirstParameter() * *p.getSecondParameter()
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
			if *p.getFirstParameter() < *p.getSecondParameter() {
				result = 1
			}
			p.updateResult(result)
		}

		if p.currentOpCode == EQUALS {
			var result int64
			if *p.getFirstParameter() == *p.getSecondParameter() {
				result = 1
			}
			p.updateResult(result)
		}

		if p.currentOpCode == RELATIVE_BASE {
			p.handleRelativeBaseOpCode()
		}

		if p.currentOpCode == HALT {
			p.complete = true
			return
		}

		p.advance()

	}
}

func (p *program) extendMemory(targetAddr int64) {
	requiredSize := int(targetAddr) + 1
	if requiredSize > len(p.intCodes) {
		newMem := make([]int64, requiredSize-len(p.intCodes))
		p.intCodes = append(p.intCodes, newMem...)
	}
}

func (p *program) handleRelativeBaseOpCode() {
	p.relativeBase = p.relativeBase + *p.getFirstParameter()
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
		inputText = strings.Trim(inputText, "\n")
		code, err := strconv.Atoi(inputText)
		if err != nil {
			panic(err)
		}
		return int64(code), nil
	} else {
		return prompt[index], nil
	}

}

func (p *program) updateResult(result int64) {
	var resultPtr *int64
	if p.currentOpCode == ADD || p.currentOpCode == MUL || p.currentOpCode == LESS_THAN || p.currentOpCode == EQUALS {
		resultPtr = p.getThirdParameter()
		*resultPtr = result
	}

	if p.currentOpCode == INPUT {
		resultPtr = p.getFirstParameter()
		*resultPtr = result
	}

	if p.currentOpCode == OUTPUT {
		resultPtr = p.getFirstParameter()
		p.output = *resultPtr
		p.allOutputs = append(p.allOutputs, p.output)
		p.lastRunOuputs = append(p.lastRunOuputs, p.output)
	}
}

func (p *program) getFirstParameter() *int64 {
	return p.getParameter(p.currentParamTypes[0], 1)
}

func (p *program) getSecondParameter() *int64 {
	return p.getParameter(p.currentParamTypes[1], 2)
}

func (p *program) getThirdParameter() *int64 {
	return p.getParameter(p.currentParamTypes[2], 3)
}

func (p *program) getParameter(parameterType, idx int) *int64 {
	if parameterType == 0 {
		pos := p.intCodes[p.pos+idx]
		p.extendMemory(pos)
		return &p.intCodes[pos]
	}
	if parameterType == 1 {
		pos := int64(p.pos + idx)
		p.extendMemory(pos)
		return &p.intCodes[pos]
	}
	if parameterType == 2 {
		pos := p.intCodes[p.pos+idx] + p.relativeBase
		p.extendMemory(pos)
		return &p.intCodes[pos]
	}
	return &p.intCodes[p.pos+idx]

}

func (p *program) parseOpcode() {
	intCode := p.intCodes[p.pos]
	paddedIntCode := fmt.Sprintf("0000%d", intCode)
	p.currentOpCode, _ = strconv.Atoi(paddedIntCode[len(paddedIntCode)-2:])
	p.currentParamTypes[0], _ = strconv.Atoi(paddedIntCode[len(paddedIntCode)-3 : len(paddedIntCode)-2])
	p.currentParamTypes[1], _ = strconv.Atoi(paddedIntCode[len(paddedIntCode)-4 : len(paddedIntCode)-3])
	p.currentParamTypes[2], _ = strconv.Atoi(paddedIntCode[len(paddedIntCode)-5 : len(paddedIntCode)-4])
}

func (p *program) advance() {
	if p.currentOpCode == ADD || p.currentOpCode == MUL {
		p.pos = p.pos + 4
	}

	if p.currentOpCode == INPUT || p.currentOpCode == OUTPUT {
		p.pos = p.pos + 2
	}

	if p.currentOpCode == JUMP_IF_TRUE {
		leftOperand := *p.getFirstParameter()
		rightOperand := *p.getSecondParameter()
		if leftOperand != 0 {
			p.pos = int(rightOperand)
		} else {
			p.pos = p.pos + 3
		}
	}

	if p.currentOpCode == JUMP_IF_FALSE {
		leftOperand := *p.getFirstParameter()
		rightOperand := *p.getSecondParameter()
		if leftOperand == 0 {
			p.pos = int(rightOperand)
		} else {
			p.pos = p.pos + 3
		}
	}

	if p.currentOpCode == LESS_THAN || p.currentOpCode == EQUALS {
		p.pos = p.pos + 4
	}

	if p.currentOpCode == RELATIVE_BASE {
		p.pos = p.pos + 2
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
