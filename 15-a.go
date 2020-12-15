package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("15.txt")
	fileContent := strings.TrimSpace(string(fileBytes))

	startingNumbers := strings.Split(fileContent, ",")
	lastTimeSpoken := make(map[int]int)
	var lastSpoken int
	var newSpoken int
	for i := 0; i < 2020; i++ {
		lastSpoken = newSpoken
		if i < len(startingNumbers) {
			newSpoken, _ = strconv.Atoi(startingNumbers[i])
		} else {
			_, wasSpoken := lastTimeSpoken[lastSpoken]
			if wasSpoken {
				newSpoken = i - lastTimeSpoken[lastSpoken]
			} else {
				newSpoken = 0
			}
		}
		if i > 0 {
			lastTimeSpoken[lastSpoken] = i
		}
	}

	fmt.Println(newSpoken)
}
