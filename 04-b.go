package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func match(pattern string, value string) bool {
	match, _ := regexp.MatchString("^(" + pattern + ")$", value)
	return match
}

func numberInRange(value string, low int, high int) bool {
	number, _ := strconv.Atoi(value)
	return low <= number && number <= high
}

func yearInRange(value string, low int, high int) bool {
	return match("[0-9]{4}", value) && numberInRange(value, low, high)
}

func main() {
	fileBytes, _ := ioutil.ReadFile("04.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")

	requiredFields := map[string]func(string2 string) bool{
		"byr" : func(value string) bool {
			return yearInRange(value, 1920, 2002)
		},
		"iyr" : func(value string) bool {
			return yearInRange(value, 2010, 2020)
		},
		"eyr" : func(value string) bool {
			return yearInRange(value, 2020, 2030)
		},
		"hgt" : func(value string) bool {
			r, _ := regexp.Compile("^([0-9]+)(cm|in)$")
			match := r.FindStringSubmatch(value)
			if match == nil {
				return false
			}
			if match[2] == "cm" {
				return numberInRange(match[1], 150, 193)
			} else {
				return numberInRange(match[1], 59, 76)
			}
		},
		"hcl" : func(value string) bool {
			return match("#[0-9a-f]{6}", value)
		},
		"ecl": func(value string) bool {
			return match("amb|blu|brn|gry|grn|hzl|oth", value)
		},
		"pid": func(value string) bool {
			return match("^[0-9]{9}$", value)
		},
	}

	validCount := 0
	currentFields := make(map[string]string)
	for _, line := range lines {
		if len(line) == 0 {
			valid := true
			for requiredField, requirement := range requiredFields {
				if currentFields[requiredField] == "" || !requirement(currentFields[requiredField]) {
					valid = false
					break
				}
			}
			if valid {
				validCount++
			}
			currentFields = make(map[string]string)
		} else {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				keyValue := strings.Split(field, ":")
				currentFields[keyValue[0]] = keyValue[1]
			}
		}
	}
	fmt.Println(validCount)
}
