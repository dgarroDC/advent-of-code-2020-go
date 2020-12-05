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

	max := int64(0)
	for _, line := range lines {
		binary := strings.Map(func(r rune) rune {
			if r == 'F' || r == 'L' {
				return '0'
			}
			return '1'
		}, line)
		seatId, _ := strconv.ParseInt(binary, 2, 64)
		if seatId > max {
			max = seatId
		}
	}
	fmt.Println(max)
}
