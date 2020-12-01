package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("01.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	numbers := make(map[int]bool)
	for _, line := range lines {
		number, _ := strconv.Atoi(line)
		numbers[number] = true
	}

	for i := 0; i < len(lines); i++ {
		numberA, _ := strconv.Atoi(lines[i])
		for j := i + 1; j < len(lines); j++ {
			numberB, _ := strconv.Atoi(lines[j])
			other := 2020 - numberA - numberB
			if numbers[other] {
				fmt.Println(numberA * numberB * other)
				return
			}
		}
	}
}
