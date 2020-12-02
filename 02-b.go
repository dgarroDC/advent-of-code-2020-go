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
		firstPos, _ := strconv.Atoi(match[1])
		secondPos, _ := strconv.Atoi(match[2])
		letter := match[3]
		password := match[4]
		firstMatch := password[firstPos-1] == letter[0]
		secondMatch := password[secondPos-1] == letter[0]
		if firstMatch != secondMatch {
			validCount++
		}
	}
	fmt.Println(validCount)
}
