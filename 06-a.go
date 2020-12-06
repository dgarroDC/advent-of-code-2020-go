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
	for _, line := range lines {
		if len(line) == 0 {
			yesCount += len(currentYes)
			currentYes = make(map[int32]bool)
		} else {
			for _, char := range line {
				currentYes[char] = true
			}
		}
	}
	fmt.Println(yesCount)
}
