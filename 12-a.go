package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func mod(a, b int) int {
	return (a % b + b) % b
}

func main() {
	fileBytes, _ := ioutil.ReadFile("12.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	pos := [2]int{0, 0}
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dir := 0

	for _, line := range lines {
		actionValue := strings.SplitN(line, "", 2)
		action := actionValue[0]
		value, _ := strconv.Atoi(actionValue[1])
		if action == "L" || action == "R" {
			value /= 90
			if action == "L" {
				value *= -1
			}
			dir += value
			dir = mod(dir, 4)
		} else if action == "F" {
			pos[0] += dirs[dir][0]*value
			pos[1] += dirs[dir][1]*value
		} else if action == "N" {
			pos[0] -= value
		} else if action == "S" {
			pos[0] += value
		} else if action == "E" {
			pos[1] += value
		} else if action == "W" {
			pos[1] -= value
		}
	}

	fmt.Println(math.Abs(float64(pos[0])) + math.Abs(float64(pos[1])))
}
