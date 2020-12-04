package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("04.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")

	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	validCount := 0
	currentFields := make(map[string]bool)
	for _, line := range lines {
		if len(line) == 0 {
			valid := true
			for _, requiredField := range  requiredFields {
				if !currentFields[requiredField] {
					valid = false
					break
				}
			}
			if valid {
				validCount++
			}
			currentFields = make(map[string]bool)
		} else {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				keyValue := strings.Split(field, ":")
				currentFields[keyValue[0]] = true
			}
		}
	}
	fmt.Println(validCount)
}
