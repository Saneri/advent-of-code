package main

import (
	"fmt"
	"math"
	"saneri/aoc/utils"
	"slices"
	"strings"
)

func main() {
	data := utils.ReadInput("input.txt")

	var sum float64 = 0
	for _, line := range data {
		groups := strings.Split(strings.Split(line, ":")[1], "|")
		winningNumbers := utils.FindNumbers(groups[0])
		ownNumbers := utils.FindNumbers(groups[1])

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
