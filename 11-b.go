package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func countOccupied(i int, j int, layout [][]int32) int {
	count := 0
	deltas := [3]int{-1, 0, 1}
	for _, di := range deltas {
		for _, dj := range deltas {
			if di != 0 || dj != 0 {
				i2 := i + di
				j2 := j + dj
				if inRange(i2, j2, layout) {
					for inRange(i2, j2, layout) && layout[i2][j2] == '.' {
						i2 += di
						j2 += dj
					}
					if inRange(i2, j2, layout) && layout[i2][j2] == '#' {
						count++
					}
				}
			}
		}
	}
	return count
}

func inRange(i int, j int, layout [][]int32) bool {
	return 0 <= i && i < len(layout) && 0 <= j && j < len(layout[i])
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
				if seat == 'L' && countOccupied(i, j, layout) == 0 {
					newLayout[i][j] = '#'
					changed = true
				} else if seat == '#' && countOccupied(i, j, layout) >= 5 {
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