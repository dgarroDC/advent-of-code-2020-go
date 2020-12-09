package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type operation interface {
	execute(argument int, p *program)
}

type acc struct {
}

func (_ acc) execute(argument int, p *program) {
	p.acc += argument
	p.pc++
}

type jmp struct {
}

func (_ jmp) execute(argument int, p *program) {
	p.pc += argument
}

type nop struct {
}

func (_ nop) execute(argument int, p *program) {
	p.pc++
}

type instruction struct {
	operation operation
	argument int
}

func (i instruction) execute(p *program) {
	i.operation.execute(i.argument, p)
}

type program struct {
	instructions []instruction
	pc int
	acc int
}

func newProgram(lines []string) *program {
	instructions := make([]instruction, 0, len(lines))
	for _, line := range lines {
		operationArgument := strings.Split(line, " ")
		mnemonic := operationArgument[0]
		argument, _ := strconv.Atoi(operationArgument[1])
		var op operation
		switch mnemonic {
		case "acc":
			op = acc{}
		case "jmp":
			op = jmp{}
		case "nop":
			op = nop{}
		}
		instructions = append(instructions, instruction{op, argument})
	}
	return &program{instructions, 0, 0}
}

func (p *program) execute() {
	instruction := p.instructions[p.pc]
	instruction.execute(p)
}

func (p *program) reset() {
	p.pc = 0
	p.acc = 0
}

func (p *program) halts() bool {
	p.reset()
	executed := make(map[int]bool)
	for !executed[p.pc] {
		executed[p.pc] = true
		p.execute()
		if p.pc == len(p.instructions) {
			return true
		}
	}
	return false
}

func main() {
	fileBytes, _ := ioutil.ReadFile("08.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	p := newProgram(lines)
	for i, _ := range p.instructions {
		oldOp := p.instructions[i].operation
		var newOp operation
		switch oldOp.(type) {
		case jmp:
			newOp = nop{}
		case nop:
			newOp = jmp{}
		}
		if newOp != nil {
			p.instructions[i].operation = newOp
			if p.halts() {
				fmt.Println(p.acc)
				return
			}
			p.instructions[i].operation = oldOp
		}
	}
}
