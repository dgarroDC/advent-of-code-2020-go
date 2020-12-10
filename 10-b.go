package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func arrangements(jolts int, joltsBag *map[int]bool, mem *map[int]int) int {
	if (*mem)[jolts] == 0 {
		res := 0
		for i := 1; i <= 3; i++ {
			if (*joltsBag)[jolts+i] {
				res += arrangements(jolts+i, joltsBag, mem)
			}
		}
		if res == 0 {
			res = 1
		}
		(*mem)[jolts] = res
	}
	return (*mem)[jolts]
}

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

	mem := make(map[int]int)
	fmt.Println(arrangements(0, &joltsBag, &mem))
}
