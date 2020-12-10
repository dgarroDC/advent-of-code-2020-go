package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("09.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	numbers := make([]int, len(lines))
	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(line)
	}

	preamble := 25
	for i := preamble; i < len(lines); i++ {
		number := numbers[i]
		numbersBefore := make(map[int]bool)
		sumFound := false
		for j := i - preamble; j < i; j++ {
			a := numbers[j]
			b := number - a
			if numbersBefore[b] {
				sumFound = true
			}
			numbersBefore[a] = true
		}
		if !sumFound {
			fmt.Println(number)
			return
		}
	}
}
