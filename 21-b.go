package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

	done := false
	for !done {
		for allergen, possibilities := range allergensPossibilities {
			done = true
			if len(possibilities) > 1 {
				done = false
				for ingredient, _ := range possibilities {
					for otherAllergen, otherPossibilities := range allergensPossibilities {
						if allergen != otherAllergen && otherPossibilities[ingredient] && len(otherPossibilities) == 1 {
							delete(allergensPossibilities[allergen], ingredient)
						}
					}
				}
			}
		}
	}

	allergens := make([]string, 0)
	for allergen := range allergensPossibilities {
		allergens = append(allergens, allergen)
	}
	sort.Strings(allergens)

	dangerousIngredients := make([]string, len(allergens))
	for i, allergen := range allergens {
		for onlyIngredient := range allergensPossibilities[allergen] {
			dangerousIngredients[i] = onlyIngredient
		}
	}
	canonicalList := strings.Join(dangerousIngredients, ",")

	fmt.Println(canonicalList)
}
