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

func (_ nop) execute(_ int, p *program) {
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

func main() {
	fileBytes, _ := ioutil.ReadFile("08.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	program := newProgram(lines)
	executed := make(map[int]bool)
	for !executed[program.pc] {
		executed[program.pc] = true
		program.execute()
	}

	fmt.Println(program.acc)
}
