package main

import (
	"fmt"
	"os"
	"strings"
)

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func main() {
	file, _ := os.ReadFile("input.txt")

	cards := delete_empty(strings.Split(string(file), "\n"))

	total_score := 0
	cards_played := 0
	copies := map[int]int{}

	for i, card := range cards {
		winning_numbers := delete_empty(strings.Split(strings.Split(strings.Split(card, "|")[0], ":")[1], " "))
		card_numbers := delete_empty(strings.Split(strings.Split(card, "|")[1], " "))
		card_score := 0
		numbers_won := 0
		cards_played += 1 + copies[i]

		for _, card_number := range card_numbers {
			for _, winning_number := range winning_numbers {
				if card_number == winning_number {
					numbers_won += 1
					if card_score == 0 {
						card_score = 1
					} else {
						card_score *= 2
					}
				}
			}
		}

		for j := 1; j <= numbers_won; j++ {
			copies[i+j] += 1 + copies[i]
		}

		total_score += card_score
	}

	fmt.Println(cards_played)
	// fmt.Println(total_score)
}
