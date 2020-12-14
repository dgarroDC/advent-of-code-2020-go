package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func writeMem(mem *map[int64]int, binaryAddress string, floatingBits []int, value int, sum *int) {
	if len(floatingBits) == 0 {
		address, _ := strconv.ParseInt(binaryAddress, 2, 64)
		*sum = *sum  + value - (*mem)[address]
		(*mem)[address] = value
	} else {
		nextBit := floatingBits[0]
		writeMem(mem, binaryAddress[:nextBit] + "0" + binaryAddress[nextBit+1:], floatingBits[1:], value, sum)
		writeMem(mem, binaryAddress[:nextBit] + "1" + binaryAddress[nextBit+1:], floatingBits[1:], value, sum)
	}
}

func main() {
	fileBytes, _ := ioutil.ReadFile("14.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	mem := make(map[int64]int)
	sum := 0
	var mask string
	for _, line := range lines {
		assignment := strings.Split(line, " = ")
		assigned := assignment[0]
		if assigned == "mask" {
			mask = assignment[1]
		} else {
			address, _ := strconv.Atoi(strings.Split(strings.Split(assigned, "[")[1], "]")[0])
			binaryAddress := strconv.FormatInt(int64(address), 2)
			intValue, _ := strconv.Atoi(assignment[1])
			for len(binaryAddress) < 36 {
				binaryAddress = "0" + binaryAddress
			}
			floatingBits := make([]int, 0)
			for i, c := range mask {
				if c == '1' {
					binaryAddress = binaryAddress[:i] + string(c) + binaryAddress[i+1:]
				} else if c == 'X' {
					floatingBits = append(floatingBits, i)
				}
			}
			writeMem(&mem, binaryAddress, floatingBits, intValue, &sum)
		}
	}

	fmt.Println(sum)
}
