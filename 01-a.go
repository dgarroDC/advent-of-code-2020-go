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
		other := 2020 - number
		if numbers[other] {
			fmt.Println(number * other)
			break
		}
		numbers[number] = true
	}
}
