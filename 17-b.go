package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func countAdjacentActive(x int, y int, z int, w int, dimension map[int]map[int]map[int]map[int]int32) int {
	count := 0
	deltas := [3]int{-1, 0, 1}
	for _, dx := range deltas {
		for _, dy := range deltas {
			for _, dz := range deltas {
				for _, dw := range deltas {
					if dx != 0 || dy != 0 || dz != 0 || dw != 0 {
						if dimension[x+dx][y+dy][z+dz][w+dw] == '#' {
							count++
						}
					}
				}
			}
		}
	}
	return count
}

func main() {
	fileBytes, _ := ioutil.ReadFile("17.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	dimension := make(map[int]map[int]map[int]map[int]int32)
	for i, line := range lines {
		dimension[i] = make(map[int]map[int]map[int]int32)
		for j, c := range line {
			dimension[i][j] = map[int]map[int]int32{0 : {0: c}}
		}
	}

	minX := -1
	maxX := len(lines) + 1
	minY := -1
	maxY := len(lines[0]) + 1
	minZ := -1
	maxZ := 1
	minW := -1
	maxW := 1

	var activeCount int
	for cycle := 0; cycle < 6; cycle++ {
		newDimension := make(map[int]map[int]map[int]map[int]int32)
		activeCount = 0
		for x := minX; x <= maxX; x++ {
			newDimension[x] = make(map[int]map[int]map[int]int32)
			for y := minY; y <= maxY; y++ {
				newDimension[x][y] = make(map[int]map[int]int32)
				for z := minZ; z <= maxZ; z++ {
					newDimension[x][y][z] = make(map[int]int32)
					for w := minW; w <= maxW; w++ {
						active := false
						adjacentActive := countAdjacentActive(x, y, z, w, dimension)
						if dimension[x][y][z][w] == '#' {
							if adjacentActive == 2 || adjacentActive == 3 {
								active = true
							}
						} else if adjacentActive == 3 {
							active = true
						}
						if active {
							newDimension[x][y][z][w] = '#'
							activeCount++
							if x - 1 < minX {
								minX = x - 1
							}
							if x + 1 > maxX {
								maxX = x + 1
							}
							if y - 1 < minY {
								minY = y - 1
							}
							if y + 1 > maxY {
								maxY = y + 1
							}
							if z - 1 < minZ {
								minZ = z - 1
							}
							if z + 1 > maxZ {
								maxZ = z + 1
							}
							if w - 1 < minW {
								minW = w - 1
							}
							if w + 1 > maxW {
								maxW = w + 1
							}
						} else {
							newDimension[x][y][z][w] = '.'
						}
					}
				}
			}
		}
		dimension = newDimension
	}

	fmt.Println(activeCount)
}