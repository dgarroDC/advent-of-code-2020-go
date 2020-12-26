package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("25.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	cardPublic, _ := strconv.Atoi(lines[0])
	doorPublic, _ := strconv.Atoi(lines[1])

	cardLoopSize := 0
	value := 1
	subjectNumber := 7
	divisor := 20201227
	for value != cardPublic {
		value *= subjectNumber
		value %= divisor
		cardLoopSize++
	}

	value = 1
	subjectNumber = doorPublic
	for i := 0; i < cardLoopSize; i++ {
		value *= subjectNumber
		value %= divisor
	}

	fmt.Println(value)
}
