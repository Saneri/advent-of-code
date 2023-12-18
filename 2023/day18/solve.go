package main

import (
	"fmt"
	"log"
	"math"
	"saneri/aoc/utils"
	"strconv"
	"strings"
)

type Border struct {
	x, y int
}

func getNextBorder(current Border, direction string, amount int) Border {
	switch direction {
	case "U":
		return Border{current.x, current.y + amount}
	case "R":
		return Border{current.x + amount, current.y}
	case "D":
		return Border{current.x, current.y - amount}
	case "L":
		return Border{current.x - amount, current.y}
	}
	log.Fatal("Invalid direction:", direction)
	return Border{0, 0}
}

func shoelace(borders []Border) int {
	sum1 := 0
	sum2 := 0
	for i := range borders[:len(borders)-1] {
		sum1 += borders[i].x * borders[i+1].y
		sum2 += borders[i].y * borders[i+1].x
	}
	return int(math.Abs(float64(sum1-sum2))) / 2
}

func main() {
	data := utils.ReadInput("input.txt")
	current := Border{0, 0}
	borders := []Border{current}
	borderLength := 0
	for _, line := range data {
		split := strings.Split(line, " ")
		direction := split[0]
		amount, _ := strconv.Atoi(split[1])
		borderLength += amount
		current = getNextBorder(current, direction, amount)
		borders = append(borders, current)
	}
	fmt.Println("a:", shoelace(borders)+borderLength/2+1)

	current = Border{0, 0}
	borders = []Border{current}
	borderLength = 0
	dir := map[string]string{
		"0": "U",
		"1": "R",
		"2": "D",
		"3": "L",
	}
	for _, line := range data {
		split := strings.Split(line, " ")[2]
		split = split[2 : len(split)-1]
		direction := string(split[len(split)-1])
		amount, error := strconv.ParseInt(split[:len(split)-1], 16, 64)
		if error != nil {
			log.Fatal(error)
		}
		borderLength += int(amount)
		current = getNextBorder(current, dir[direction], int(amount))
		borders = append(borders, current)
	}
	fmt.Println("b:", shoelace(borders)+borderLength/2+1)
}
