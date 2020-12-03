package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("03.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	trees := 0
	x := 0
	for _, line := range lines {
		if line[x] == '#' {
			trees++
		}
		x += 3
		x %= len(line)
	}
	fmt.Println(trees)
}
