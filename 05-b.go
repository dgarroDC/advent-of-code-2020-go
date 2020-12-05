package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("05.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	seats := make(map[int64]bool)
	for _, line := range lines {
		binary := strings.Map(func(r rune) rune {
			if r == 'F' || r == 'L' {
				return '0'
			}
			return '1'
		}, line)
		seatId, _ := strconv.ParseInt(binary, 2, 64)
		seats[seatId] = true
	}

	highestBack, _ := strconv.ParseInt("0000000111", 2, 64)
	lowestFront, _ := strconv.ParseInt("1111111000", 2, 64)

	for seatId := highestBack + 1; seatId < lowestFront; seatId++ {
		if !seats[seatId] {
			fmt.Println(seatId)
			return
		}
	}
}
