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
		sum += eval(line, &i, false)
	}

	fmt.Println(sum)
}

func eval(line string, i *int, par bool) int {
	res := 0
	for *i < len(line) {
		c := line[*i:*i+1]
		*i++
		if c == ")" {
			if !par {
				*i--
			}
			break
		} else if c == "*" {
			res *= eval(line, i, false)
		} else if c == "+" {
			continue
		} else {
			var num int
			if c == "(" {
				num = eval(line, i, true)
			} else {
				num, _ = strconv.Atoi(c)
			}
			res += num
		}
	}
	return res
}


//func eval(line string, i *int) int {
//	var first int
//	if line[*i] == '(' {
//		*i++
//		first = eval(line, i)
//	} else {
//		first, _ = strconv.Atoi(line[*i:*i+1])
//		*i++
//	}
//	if *i == len(line) {
//		return first
//	} else {
//		c := line[*i]
//		*i++
//		if c == ')' {
//			return first
//		} else {
//			second := eval(line, i)
//			if c == '+' {
//				return first + second
//			} else {
//				return first * second
//			}
//		}
//	}
//
//}

// (1 + 1 * (1 * 1)) + 1