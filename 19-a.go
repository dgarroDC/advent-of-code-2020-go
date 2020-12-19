package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func toposort(rules map[int]string, i int, visted []bool, sorted *[]int, numberRegexp *regexp.Regexp) {
	visted[i] = true

    dependencies := numberRegexp.FindAllString(rules[i], -1)
	for _, dependency := range dependencies {
		j, _ := strconv.Atoi(dependency)
		if !visted[j] {
			toposort(rules, j, visted, sorted, numberRegexp)
		}
	}

	*sorted = append(*sorted, i)
}


func main() {
	fileBytes, _ := ioutil.ReadFile("19.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	rules := make(map[int]string, 0)
	for _, line := range lines {
		if line == "" {
			break
		}
		numberRule := strings.Split(line, ": ")
		number, _ := strconv.Atoi(numberRule[0])
		rules[number] = numberRule[1]
	}

	numberRegexp, _ := regexp.Compile("[0-9]+")
	sorted := make([]int, 0)
	visted := make([]bool, len(rules))
	for i := 0; i < len(rules); i++ {
		if !visted[i] {
			toposort(rules, i, visted, &sorted, numberRegexp)
		}
	}

	regexps := make([]string, len(rules))
	for _, i := range sorted {
		if !numberRegexp.MatchString(rules[i]) {
			regexps[i] = rules[i][1:2]
		} else {
			options := strings.Split(rules[i], " | ")
			regexpOptions := make([]string, len(options))
			for i, option := range options {
				r := ""
				numbers := strings.Split(option, " ")
				for _, number := range numbers {
					numberInt, _ := strconv.Atoi(number)
					r += "(" + regexps[numberInt] + ")"
				}
				regexpOptions[i] = r
			}
			regexps[i] = strings.Join(regexpOptions, "|")
		}
	}

	rule0, _ := regexp.Compile("^" + regexps[0] + "$")
	matches := 0
	for _, msg := range lines[len(rules)+1:] {
		if rule0.MatchString(msg) {
			matches++
		}
	}

	fmt.Println(matches)
}