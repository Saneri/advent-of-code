package main

import (
	"fmt"
	"log"
	"saneri/aoc/utils"
	"strings"
)

func extrapolate(history []int, backwards bool) int {
	if (len(history) == 1) {
		log.Fatal("history is empty")
	}
	allZeros := true
	diffs := []int{}
	for i := range(history[1:]) {
		diff := history[i+1] - history[i]
		diffs = append(diffs, diff)
		if diff != 0 {
			allZeros = false
		}
	}
	if !allZeros {
		diffToNext := extrapolate(diffs, backwards)
		if (backwards) {
			return history[0] - diffToNext
		}
		return history[len(history)-1] + diffToNext
	}
	return history[0]
}

func main() {
	data := utils.ReadInput("input.txt")
	aSum := 0
	bSum := 0 
	for _, line := range data {
		history := utils.ArrayStringToInt(strings.Split(line, " "))
		aSum += extrapolate(history, false)
		bSum += extrapolate(history, true)
	}
	fmt.Println("a:", aSum)
	fmt.Println("b:", bSum)
}