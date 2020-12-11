package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func countAdjacentsOccupied(i int, j int, layout [][]int32) int {
	count := 0
	deltas := [3]int{-1, 0, 1}
	for _, di := range deltas {
		i2 := i + di
		if 0 <= i2 && i2 < len(layout) {
			for _, dj := range deltas {
				if di != 0 || dj != 0 {
					j2 := j + dj
					if 0 <= j2 && j2 < len(layout[i2]) && layout[i2][j2] == '#' {
						count++
					}
				}
			}
		}
	}
	return count
}

func main() {
	fileBytes, _ := ioutil.ReadFile("11.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	layout := make([][]int32, len(lines))
	for i, line := range lines {
		layout[i] = make([]int32, len(line))
		for j, c := range line {
			layout[i][j] = c
		}
	}

	for {
		newLayout := make([][]int32, len(lines))
		occupied := 0
		changed := false
		for i, row := range layout {
			newLayout[i] = make([]int32, len(row))
			for j, seat := range row {
				if seat == 'L' && countAdjacentsOccupied(i, j, layout) == 0 {
					newLayout[i][j] = '#'
					changed = true
				} else if seat == '#' && countAdjacentsOccupied(i, j, layout) >= 4 {
					newLayout[i][j] = 'L'
					changed = true
				} else {
					newLayout[i][j] = seat
				}
				if newLayout[i][j] == '#' {
					occupied++
				}
			}
		}
		if !changed {
			fmt.Println(occupied)
			return
		}
		layout = newLayout
	}
}