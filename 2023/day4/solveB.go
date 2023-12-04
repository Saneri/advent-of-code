package main

import (
	"fmt"
	"log"
	"regexp"
	"saneri/aoc/utils"
	"slices"
	"strconv"
	"strings"
)

func arrayStringToInt(strArray []string) []int {
	intArray := []int{}
	for _, element := range strArray {
		number, err := strconv.Atoi(element)
		if err != nil {
			log.Fatal(err)
		}
		intArray = append(intArray, number)
	}
	return intArray
}

func main() {
	data := utils.ReadInput("input.txt")

	iteratorIndex := 0
	tickets := make([]int, len(data))
	for i := range tickets {
		tickets[i] = 1
	}
	re := regexp.MustCompile(`\d+`)
	for _, line := range data {
		groups := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumbers := arrayStringToInt(re.FindAllString(groups[0], -1))
		ownNumbers := arrayStringToInt(re.FindAllString(groups[1], -1))

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
