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
	waypoint := [2]int{-1, 10}

	for _, line := range lines {
		actionValue := strings.SplitN(line, "", 2)
		action := actionValue[0]
		value, _ := strconv.Atoi(actionValue[1])
		if action == "L" || action == "R" {
			value /= 90
			if action == "L" {
				value *= -1
			}
			value = mod(value, 4)
			for i := 0; i < value; i++ {
				waypoint = [2]int{waypoint[1], -waypoint[0]}
			}
		} else if action == "F" {
			pos[0] += waypoint[0]*value
			pos[1] += waypoint[1]*value
		} else if action == "N" {
			waypoint[0] -= value
		} else if action == "S" {
			waypoint[0] += value
		} else if action == "E" {
			waypoint[1] += value
		} else if action == "W" {
			waypoint[1] -= value
		}
	}

	fmt.Println(math.Abs(float64(pos[0])) + math.Abs(float64(pos[1])))
}
