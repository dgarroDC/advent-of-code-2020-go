package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("21.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	ingredientsAppearances := make(map[string]int)
	allergensPossibilities := make(map[string]map[string]bool)
	for _, line := range lines {
		ingrendientsAllergens := strings.SplitN(line, " (", 2)
		ingrendients := strings.Split(ingrendientsAllergens[0], " ")
		allergens := strings.Split(strings.Split(strings.Split(ingrendientsAllergens[1], "contains ")[1], ")")[0], ", ")
		for _, ingrendient := range ingrendients {
			ingredientsAppearances[ingrendient]++
		}
		for _, allergen := range allergens {
			newPossibilities := make(map[string]bool)
			_, notFirstTime := allergensPossibilities[allergen]
			for _, ingrendient := range ingrendients {
				if !notFirstTime || allergensPossibilities[allergen][ingrendient] {
					newPossibilities[ingrendient] = true
				}
			}
			allergensPossibilities[allergen] = newPossibilities
		}
	}

	allAppearances := 0
	for ingredient, appearances := range ingredientsAppearances {
		shouldCount := true
		for _, possibilities := range allergensPossibilities {
			if possibilities[ingredient] {
				shouldCount = false
				break
			}
		}
		if shouldCount {
			allAppearances += appearances
		}
	}

	fmt.Println(allAppearances)
}
