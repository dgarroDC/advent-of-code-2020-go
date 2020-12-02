package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("02.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	validCount := 0
	r, _ := regexp.Compile("([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)")
	for _, line := range lines {
		match := r.FindStringSubmatch(line)
		lowest, _ := strconv.Atoi(match[1])
		highest, _ := strconv.Atoi(match[2])
		letter := match[3]
		password := match[4]
		count := strings.Count(password, letter)
		if lowest <= count && count <= highest {
			validCount++
		}
	}
	fmt.Println(validCount)
}
