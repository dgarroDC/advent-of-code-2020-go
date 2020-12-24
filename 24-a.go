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

func main() {
	fileBytes, _ := ioutil.ReadFile("24.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines)-1]

	tiles := make(map[pos]bool)
	for _, line := range lines {
		x := 0
		y := 0
		for i := 0; i < len(line); i++ {
			dir := string(line[i])
			if dir == "n" || dir == "s" {
				i++
				dir += string(line[i])
			}
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
		}
		tiles[pos{x, y}] = !tiles[pos{x, y}]
	}

	blackCount := 0
	for _, color := range tiles {
		if color {
			blackCount++
		}
	}

	fmt.Println(blackCount)
}