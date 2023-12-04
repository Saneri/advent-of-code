package main

import (
	"fmt"
	"log"
	"math"
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

	re := regexp.MustCompile(`\d+`)
	var sum float64 = 0
	for _, line := range data {
		groups := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumbers := arrayStringToInt(re.FindAllString(groups[0], -1))
		ownNumbers := arrayStringToInt(re.FindAllString(groups[1], -1))

		var winningNumberCount float64 = 0
		for _, ownNumber := range ownNumbers {
			if slices.Contains(winningNumbers, ownNumber) {
				winningNumberCount++
			}
		}
		if winningNumberCount > 0 {
			sum += (math.Pow(2, winningNumberCount-1))
		}
	}
	fmt.Printf("%.0f\n", sum)
}
