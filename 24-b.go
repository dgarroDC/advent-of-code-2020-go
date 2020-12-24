package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type pos struct {
	x int
	y int
}

func posToDir(x int, y int, dir string) pos {
	switch dir {
	case "e":
		x++
	case "se":
		if y%2 != 0 {
			x++
		}
		y++
	case "sw":
		if y%2 == 0 {
			x--
		}
		y++
	case "w":
		x--
	case "nw":
		if y%2 == 0 {
			x--
		}
		y--
	case "ne":
		if y%2 != 0 {
			x++
		}
		y--
	}
	return pos{x, y}
}

func adjacentPositions(position pos) []pos {
	result := make([]pos, 6)
	dirs := [6]string{"e", "se", "sw", "w", "nw", "ne"}
	for i, dir := range dirs {
		result[i] = posToDir(position.x, position.y, dir)
	}
	return result
}

func countAdjacentBlacks(position pos, blacks map[pos]bool) int {
	count := 0
	for _, adjacentPos := range adjacentPositions(position) {
		if blacks[adjacentPos] {
			count++
		}
	}
	return count
}

func main() {
	fileBytes, _ := ioutil.ReadFile("24.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines)-1]

	blacks := make(map[pos]bool)
	for _, line := range lines {
		currentPos := pos{0, 0}
		for i := 0; i < len(line); i++ {
			dir := string(line[i])
			if dir == "n" || dir == "s" {
				i++
				dir += string(line[i])
			}
			currentPos = posToDir(currentPos.x, currentPos.y, dir)
		}
		if blacks[currentPos] {
			delete(blacks, currentPos)
		} else {
			blacks[currentPos] = true
		}
	}

	for day := 0; day < 100; day++ {
		newBlacks := make(map[pos]bool)
		for blackPos := range blacks {
			adjacentBlacks := countAdjacentBlacks(blackPos, blacks)
			if adjacentBlacks == 1 || adjacentBlacks == 2{
				newBlacks[blackPos] = true
			}
			for _, adjacentPos := range adjacentPositions(blackPos) {
				if !blacks[adjacentPos] && !newBlacks[adjacentPos] && countAdjacentBlacks(adjacentPos, blacks) == 2 {
					newBlacks[adjacentPos] = true
				}
			}
		}
		blacks = newBlacks
	}

	fmt.Println(len(blacks))
}