package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func containsShinyGold(bag string, contents map[string][]string) bool {
	for _, bagInside := range contents[bag] {
		if bagInside == "shiny gold" || containsShinyGold(bagInside, contents) {
			return true
		}
	}
	return false
}

func main() {
	fileBytes, _ := ioutil.ReadFile("07.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines)-1]

	contents := make(map[string][]string)
	for _, line := range lines {
		rule := strings.Split(line, "contain ")
		container := strings.Split(rule[0], " bags")[0]
		content := make([]string, 0)
		rule[1] = rule[1][:len(rule[1])-1]
		if rule[1] != "no other bags" {
			contentBags := strings.Split(rule[1], ", ")
			for _, contentBag := range contentBags {
				contentBag = strings.Split(contentBag, " bag")[0]
				bag := strings.SplitN(contentBag, " ", 2)[1]
				content = append(content, bag)
			}
		}
		contents[container] = content
	}

	bagsWithShinyGold := 0
	for bag, _ := range contents {
		if containsShinyGold(bag, contents) {
			bagsWithShinyGold++
		}
	}
	fmt.Println(bagsWithShinyGold)
}