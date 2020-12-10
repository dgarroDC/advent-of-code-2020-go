package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("09.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	numbers := make([]int, len(lines))
	for i, line := range lines {
		numbers[i], _ = strconv.Atoi(line)
	}

	target := 50047984
	l := 0
	r := 1
	sum := numbers[l] + numbers[r]
	for {
		if sum == target {
			min := numbers[l]
			max := numbers[l]
			for _, number := range numbers[l+1:r+1] {
				if number < min {
					min = number
				}
				if number > max {
					max = number
				}
			}
			fmt.Println(min + max)
			return
		} else if sum < target {
			r++
			sum += numbers[r]
		} else if sum > target {
			sum -= numbers[l]
			l++
		}
	}
}
