package main

import (
	"fmt"
	"saneri/aoc/utils"
	"sort"
	"strconv"
	"strings"
)

var pokerCards = map[string]int{
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"T":  10,
	"J":  1,
	"Q":  12,
	"K":  13,
	"A":  14,
}

func compareHand(hand string) int {
	split := strings.Split(hand, " ")
	cards := split[0]
	cardCounts := make(map[string]int)
	maxValue := 0
	maxKey := ""
	for _, card := range cards {
		cardCounts[string(card)]++
		count := cardCounts[string(card)]
		if (string(card) != "J") { 
			if (count > maxValue) {
				maxValue = count
				maxKey = string(card)
			} else if (count == maxValue) && (pokerCards[string(card)] > pokerCards[maxKey]) {
				maxKey = string(card)
			}
		}
	}
	if (maxValue > 0) {
		cardCounts[maxKey] += cardCounts["J"]
		delete(cardCounts, "J")
	}
	score := 1
	pairCount := 0
	threeCount := 0
	for _, count := range cardCounts {
 		if count == 5 {
			score = max(7, score)
		} else if count == 4 {
			score = max(6, score)
		} else if count == 3 {
			score = max(4, score)
			threeCount++
		} else if count == 2 {
			score = max(2, score)
			pairCount++
		}
	}
	if threeCount > 0 && pairCount > 0 {
		score = max(5, score)
	} else if pairCount > 1 {
		score = max(3, score)
	}
	return score
}

func main() {
	data := utils.ReadInput("input.txt")
	sort.Slice(data, func(i, j int) bool {
		hand1 := compareHand(data[i])
		hand2 := compareHand(data[j]) 
		if hand1 == hand2 {
			cards1 := strings.Split(data[i], " ")[0]
			cards2 := strings.Split(data[j], " ")[0]
			for index := range(cards1) {
				if cards1[index] != cards2[index] {
					return pokerCards[string(cards1[index])] < pokerCards[string(cards2[index])]
				}
			}
		}
		return hand1 < hand2
	})

	sum := 0
	for index, hand := range(data) {
		bet, _ := strconv.Atoi(strings.Split(hand, " ")[1])
		sum += (index+1) * bet
	}
	fmt.Println(sum)
}
