package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("23.txt")
	fileContent := strings.TrimSpace(string(fileBytes))

	cupsS := strings.Split(fileContent, "")
	total := 1000000
	initialCups := make([]int, len(cupsS))
	nexts := make([]int, total)
	for i, cupS := range cupsS {
		initialCups[i], _ = strconv.Atoi(cupS)
	}
	for i := 0; i < len(initialCups) - 1; i++ {
		nexts[initialCups[i]-1] = initialCups[i+1]
	}
	nexts[initialCups[len(initialCups)-1]-1] = len(initialCups) + 1
	for cup := len(initialCups) + 1; cup < total; cup++ {
		nexts[cup-1] = cup + 1
	}
	nexts[total-1] = initialCups[0]

	current := initialCups[0]
	for move := 0; move < 10000000; move++ {
		pick0 := nexts[current-1]
		pick1 := nexts[pick0-1]
		pick2 := nexts[pick1-1]
		dest := current - 1
		if dest == 0 {
			dest = total
		}
		for dest == pick0 || dest == pick1 || dest == pick2 {
			dest--
			if dest == 0 {
				dest = total
			}
		}
		nexts[current-1] = nexts[pick2-1]
		nexts[pick2-1] = nexts[dest-1]
		nexts[dest-1] = pick0
		current = nexts[current-1]
	}

	mul := nexts[0] * nexts[nexts[0]-1]

	fmt.Println(mul)
}
