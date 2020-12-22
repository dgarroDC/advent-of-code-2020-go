package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("22.txt")
	fileContent := string(fileBytes)
	lines := strings.Split(fileContent, "\n")
	lines = lines[:len(lines) - 1]

	decks := make([][]int, 2)
	readingDeck := 0
	for _, line := range lines[1:] {
		if line == "Player 2:" {
			readingDeck = 1
		} else if line != "" {
			card, _ := strconv.Atoi(line)
			decks[readingDeck] = append(decks[readingDeck], card)
		}
	}

	for len(decks[0]) != 0 && len(decks[1]) != 0 {
		wins := 0
		if decks[0][0] < decks[1][0] {
			wins = 1
		}
		loses := 1 - wins
		decks[wins] = append(append(decks[wins][1:], decks[wins][0]), decks[loses][0])
		decks[loses] = decks[loses][1:]
	}

	winner := 0
	if len(decks[0]) == 0 {
		winner = 1
	}
	score := 0
	cards := len(decks[winner])
	for i, card := range decks[winner] {
		score += (cards - i) * card
	}

	fmt.Println(score)
}
