package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func reverse(s string) string {
	reversed := ""
	for _, c := range s {
		reversed = string(c) + reversed
	}
	return reversed
}

func normalize(s string) string {
	reversed := reverse(s)
	if strings.Compare(s, reversed) < 0 {
		return s
	} else {
		return reversed
	}
}

func main() {
	fileBytes, _ := ioutil.ReadFile("20.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	bordersCount := make(map[string]int)
	tileBorders := make(map[int][]string)
	id := -1
	borders := make([]string, 4)
	for i, line := range lines {
		if id == -1 {
			id, _ = strconv.Atoi(strings.Split(strings.Split(line, "Tile ")[1], ":")[0])
		} else if line == "" {
			borders[2] = lines[i-1]
			for _, border := range borders {
				bordersCount[normalize(border)]++
			}
			tileBorders[id] = borders
			id = -1
			borders = make([]string, 4)
		} else {
			if borders[0] == "" {
				borders[0] = line
			}
			borders[3] += line[:1]
			borders[1] += line[len(line)-1:]
		}
	}

	// Nice input property: bordersCount[normalize(border)] is 1 or 2
	// Assumption: 2 is from 2 distinct tiles, this fails with the sample input, but works with my input

	cornersMul := 1
	for id, borders := range tileBorders {
		matchCount := 0
		for _, border := range borders {
			if bordersCount[normalize(border)] == 2 {
				matchCount++
			}
		}
		if matchCount == 2 {
			cornersMul *= id
		}
	}

	fmt.Println(cornersMul)
}
