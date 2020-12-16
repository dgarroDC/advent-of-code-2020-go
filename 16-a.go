package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("16.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	fields := make([][]int, 0)
	var i int
	for i = 0; lines[i] != ""; i++ {
		rangesS := strings.Split(strings.Split(lines[i], ": ")[1], " or ")
		rangesS = append(strings.Split(rangesS[0], "-"), strings.Split(rangesS[1], "-")...)
		ranges := make([]int, 4)
		for i, v := range rangesS {
			ranges[i], _ = strconv.Atoi(v)
		}
		fields = append(fields, ranges)
	}

	errorRate := 0
	for _, line := range lines[i+5:] {
		valuesS := strings.Split(line, ",")
		for _, valueS := range valuesS {
			value, _ := strconv.Atoi(valueS)
			someValid := false
			for _, ranges := range fields {
				if (ranges[0] <= value && value <= ranges[1]) || (ranges[2] <= value && value <= ranges[3]) {
					someValid = true
					break
				}
			}
			if !someValid {
				errorRate += value
			}
		}
	}

	fmt.Println(errorRate)
}
