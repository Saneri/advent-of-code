package main

import (
	"fmt"
	"saneri/aoc/utils"
	"strings"
)

func getBeatRecordCount(time int, distance int) int {
	count := 0
	for i:=1; i <= time; i++ {
		if (time - i) * i > distance {
			count++
		}
	}
	return count
}

func solve(input string) int {
	data := strings.Split(input, "\n")
	times := utils.FindNumbers(data[0])
	distances := utils.FindNumbers(data[1])

	sum := 1
	for index := range(times) {
		time := times[index]
		distance := distances[index]
		sum *= getBeatRecordCount(time, distance)
	}
	return sum
}
 
func main() {
	fmt.Println("Part 1:", solve(utils.ReadInputString("input.txt")))
	fmt.Println("Part 2:", solve(strings.Replace(utils.ReadInputString("input.txt"), " ", "", -1)))
}