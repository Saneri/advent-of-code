package main

import (
	"fmt"
	"saneri/aoc/utils"
	"slices"
	"strings"
)

func main() {
	data := utils.ReadInput("input.txt")

	iteratorIndex := 0
	tickets := make([]int, len(data))
	for i := range tickets {
		tickets[i] = 1
	}
	for _, line := range data {
		groups := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumbers := utils.FindNumbers(groups[0])
		ownNumbers := utils.FindNumbers(groups[1])

		winningNumberCount := 0
		for _, ownNumber := range ownNumbers {
			if slices.Contains(winningNumbers, ownNumber) {
				winningNumberCount++
			}
		}

		amountOfTickets := tickets[iteratorIndex]
		iteratorIndex++
		for i := iteratorIndex; i < winningNumberCount+iteratorIndex; i++ {
			tickets[i] += amountOfTickets
		}
	}
	sum := 0
	for _, ticketCount := range tickets {
		sum += ticketCount
	}
	fmt.Println(sum)
}
