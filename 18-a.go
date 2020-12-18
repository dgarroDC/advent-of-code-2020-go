package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("18.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	sum := 0
	for _, line := range lines {
		line := strings.Replace(line, " ", "", -1)
		i := 0
		sum += eval(line, &i)
	}

	fmt.Println(sum)
}

func eval(line string, i *int) int {
	res := 0
	nextOp := "+"
	for *i < len(line) {
		c := line[*i:*i+1]
		*i++
		if c == ")" {
			break
		} else if c == "+" || c == "*" {
			nextOp = c
		} else {
			var num int
			if c == "(" {
				num = eval(line, i)
			} else {
				num, _ = strconv.Atoi(c)
			}
			if nextOp == "+" {
				res += num
			} else  {
				res *= num
			}
		}
	}
	return res
}