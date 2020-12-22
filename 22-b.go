package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getWinner(decks [][]int) int {
	history := make([][][]int, 0)
	for len(decks[0]) != 0 && len(decks[1]) != 0 {
		for _, pastDecks := range history {
			if len(decks[0]) == len(pastDecks[0]) {
				sameDecks := true
				for i := range decks[0] {
					if decks[0][i] != pastDecks[0][i] {
						sameDecks = false
						break
					}
				}
				if sameDecks {
					for i := range decks[1] {
						if decks[1][i] != pastDecks[1][i] {
							sameDecks = false
							break
						}
					}
					if sameDecks {
						return 0
					}
				}
			}
		}

		copyDecks := make([][]int, 2)
		for player := 0; player <= 1; player++ {
			copyDecks[player] = make([]int, len(decks[player]))
			for i, card := range decks[player] {
				copyDecks[player][i] = card
			}
		}
		history = append(history, copyDecks)

		wins := 0
		if len(decks[0]) > decks[0][0] && len(decks[1]) > decks[1][0] {
			subDecks := make([][]int, 2)
			for player := 0; player <= 1; player++ {
				subDecks[player] = make([]int, decks[player][0])
				for i, card := range decks[player][1:decks[player][0]+1] {
					subDecks[player][i] = card
				}
			}
			wins = getWinner(subDecks)
		} else if decks[0][0] < decks[1][0] {
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

	return winner
}

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

	winner := getWinner(decks)
	score := 0
	cards := len(decks[winner])
	for i, card := range decks[winner] {
		score += (cards - i) * card
	}


	fmt.Println(score)
}
