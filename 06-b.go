package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("06.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")

	yesCount := 0
	currentYes := make(map[int32]bool)
	firstInGroup := true
	for _, line := range lines {
		if len(line) == 0 {
			yesCount += len(currentYes)
			currentYes = make(map[int32]bool)
			firstInGroup = true
		} else {
			newCurrentYes :=  make(map[int32]bool)
			for _, char := range line {
				if firstInGroup || currentYes[char] {
					newCurrentYes[char] = true
				}
			}
			currentYes = newCurrentYes
			firstInGroup = false
		}
	}
	fmt.Println(yesCount)
}
