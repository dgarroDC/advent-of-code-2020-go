package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("23.txt")
	fileContent := strings.TrimSpace(string(fileBytes))

	cupsS := strings.Split(fileContent, "")
	total := len(cupsS)
	cups := make([]int, total)
	for i, cupS := range cupsS {
		cups[i], _ = strconv.Atoi(cupS)
	}

	current := 0
	for move := 0; move < 100; move++ {
		pick := []int{cups[(current+1)%total], cups[(current+2)%total], cups[(current+3)%total]}
		current = cups[current]
		dest := current - 1
		if dest == 0 {
			dest = total
		}
		for pick[0] == dest || pick[1] == dest || pick[2] == dest {
			dest--
			if dest == 0 {
				dest = total
			}
		}
		newCups := make([]int, 0)
		var destI int
		for i := 0; i < total; i++ {
			cup := cups[i]
			if pick[0] != cup && pick[1] != cup && pick[2] != cup {
				if cup == dest {
					destI = len(newCups)
				}
				newCups = append(newCups, cup)
			}
		}
		cups = newCups
		newCups = make([]int, 0)
		for i, cup := range cups {
			newCups = append(newCups, cup)
			if i == destI {
				newCups = append(newCups, pick...)
			}
		}
		cups = newCups
		for i, cup := range cups {
			if cup == current {
				current = (i + 1) % total
				break
			}
		}
	}

	found1 := false
	labels := ""
	labelsBefore1 := ""
	for _, cup := range cups {
		if cup == 1 {
			found1 = true
		} else if found1 {
			labels += strconv.Itoa(cup)
		} else {
			labelsBefore1 += strconv.Itoa(cup)
		}
	}
	labels += labelsBefore1

	fmt.Println(labels)
}
