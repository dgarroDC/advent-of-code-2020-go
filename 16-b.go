package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func valid(ranges []int, value int) bool {
	return (ranges[0] <= value && value <= ranges[1]) || (ranges[2] <= value && value <= ranges[3])
}

func testPossibilities(posiblesFieldsByIndex *[]map[string]bool, solution *map[string]int) bool {
	nextIndex := len(*solution)
	if nextIndex == len(*posiblesFieldsByIndex) {
		return true
	}
	for field, _  := range (*posiblesFieldsByIndex)[nextIndex] {
		if _, ok := (*solution)[field]; !ok {
			(*solution)[field] = nextIndex
			if testPossibilities(posiblesFieldsByIndex, solution) {
				return true
			}
			delete(*solution, field)
		}
	}
	return false
}

// Kind of slow but ok, maybe checking posiblesIndexesByField instead?
func main() {
	fileBytes, _ := ioutil.ReadFile("16.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	fields := make(map[string][]int, 0)
	var i int
	for i = 0; lines[i] != ""; i++ {
		fieldRanges := strings.Split(lines[i], ": ")
		field := fieldRanges[0]
		rangesS := strings.Split(fieldRanges[1], " or ")
		rangesS = append(strings.Split(rangesS[0], "-"), strings.Split(rangesS[1], "-")...)
		ranges := make([]int, 5)
		for i, v := range rangesS {
			ranges[i], _ = strconv.Atoi(v)
		}
		fields[field] = ranges
	}

	myTicket := strings.Split(lines[i+2], ",")

	validTickets := make([][]string, 0)
	for _, line := range lines[i+5:] {
		valuesS := strings.Split(line, ",")
		someInvalid := false
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
				someInvalid = true
				break
			}
		}
		if !someInvalid {
			validTickets = append(validTickets, valuesS)
		}
	}

	posiblesFieldsByIndex := make([]map[string]bool, 0)
	for i := 0; i < len(validTickets[0]); i++ {
		posiblesFields := make(map[string]bool)
		for field, _ := range fields{
			posiblesFields[field] = true
		}
		posiblesFieldsByIndex = append(posiblesFieldsByIndex, posiblesFields)
	}
	for _, ticket := range validTickets {
		for i, valueS := range ticket {
			for field, _ := range posiblesFieldsByIndex[i] {
				value, _ := strconv.Atoi(valueS)
				ranges := fields[field]
				if !valid(ranges, value) {
					delete(posiblesFieldsByIndex[i], field)
				}
			}
		}
	}
	notDone := true
	for notDone {
		for _, posiblesFields := range posiblesFieldsByIndex {
			notDone = false
			if len(posiblesFields) == 1 {
				var onlyField string
				for onlyField, _ = range posiblesFields {
				}
				for _, posiblesFields2 := range posiblesFieldsByIndex {
					if len(posiblesFields2) > 1 && posiblesFields2[onlyField] {
						notDone = true
						delete(posiblesFields2, onlyField)
					}
				}
			}
		}
	}

	solution := make(map[string]int, 0)
	testPossibilities(&posiblesFieldsByIndex, &solution)

	multiplication := 1
	for field, index := range solution{
		if strings.Split(field, " ")[0] == "departure"{
			myTicketValue, _ := strconv.Atoi(myTicket[index])
			multiplication *= myTicketValue
		}
	}
	fmt.Println(multiplication)
}
