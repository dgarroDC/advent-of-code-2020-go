package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("14.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	mem := make(map[int]int)
	sum := 0
	var mask string
	for _, line := range lines {
		assignment := strings.Split(line, " = ")
		assigned := assignment[0]
		if assigned == "mask" {
			mask = assignment[1]
		} else {
			address, _ := strconv.Atoi(strings.Split(strings.Split(assigned, "[")[1], "]")[0])
			intValue, _ := strconv.Atoi(assignment[1])
			binaryValue := strconv.FormatInt(int64(intValue), 2)
			for len(binaryValue) < 36 {
				binaryValue = "0" + binaryValue
			}
			for i, c := range mask {
				if c != 'X' {
					binaryValue = binaryValue[:i] + string(c) + binaryValue[i+1:]
				}
			}
			int64Value, _ := strconv.ParseInt(binaryValue, 2, 64)
			intValue = int(int64Value)
			sum = sum  + intValue - mem[address]
			mem[address] = intValue
		}
	}

	fmt.Println(sum)
}
