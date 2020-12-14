package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("13.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")

	earliest, _ := strconv.Atoi(lines[0])
	ids := strings.Split(lines[1], ",")

	min := earliest
	var idMin int
	for _, idString := range ids {
		if idString != "x" {
			id, _ := strconv.Atoi(idString)
			wait := id - (earliest % id)
			if wait == id {
				wait = 0
			}
			if wait < min {
				min = wait
				idMin = id
			}
		}
	}

	fmt.Println(idMin * min)
}
