package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("10.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	joltsBag := make(map[int]bool, len(lines))
	for _, line := range lines {
		jolts, _ := strconv.Atoi(line)
		joltsBag[jolts] = true
	}

	jolts := 0
	joltsUsed := make([]int, 3)
	for {
		used := false
		for i := 1; i <= 3; i++ {
			if joltsBag[jolts+i] {
				jolts += i
				joltsUsed[i-1]++
				used = true
				break
			}
		}
		if !used {
			fmt.Println(joltsUsed[0] * (joltsUsed[2] + 1))
			return
		}
	}
}
