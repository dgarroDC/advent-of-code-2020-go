package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("03.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	width := len(lines[0])
	height := len(lines)

	result := 1
	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, slope := range slopes {
		trees := 0
		x := 0
		for y := 0;  y < height; y += slope[1] {
			if lines[y][x] == '#' {
				trees++
			}
			x += slope[0]
			x %= width
		}
		result *= trees
	}
	fmt.Println(result)
}
