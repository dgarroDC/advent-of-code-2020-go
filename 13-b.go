package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func gcd(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	fileBytes, _ := ioutil.ReadFile("13.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")

	ids := strings.Split(lines[1], ",")

	l := 1
	t := 0
	for i, idString := range ids {
		if idString != "x" {
			id, _ := strconv.Atoi(idString)
			for (t + i) % id != 0 {
				t += l
			}
			l = lcm(l, id)
		}
	}

	fmt.Println(t)
}
