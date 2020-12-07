package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type contentDescription struct {
	quantity int
	bag string
}

func bagsInsideCount(bag string, contents map[string][]contentDescription) int {
	count := 0
	for _, bagInside := range contents[bag] {
		count += bagInside.quantity * (1 + bagsInsideCount(bagInside.bag, contents))
	}
	return count
}


func main() {
	fileBytes, _ := ioutil.ReadFile("07.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines)-1]

	contents := make(map[string][]contentDescription)
	for _, line := range lines {
		rule := strings.Split(line, "contain ")
		container := strings.Split(rule[0], " bags")[0]
		content := make([]contentDescription, 0)
		rule[1] = rule[1][:len(rule[1])-1]
		if rule[1] != "no other bags" {
			contentBags := strings.Split(rule[1], ", ")
			for _, contentBag := range contentBags {
				contentBag = strings.Split(contentBag, " bag")[0]
				quantityBag := strings.SplitN(contentBag, " ", 2)
				quantity, _ := strconv.Atoi(quantityBag[0])
				content = append(content, contentDescription{quantity, quantityBag[1]})
			}
		}
		contents[container] = content
	}

	fmt.Println(bagsInsideCount("shiny gold", contents))
}