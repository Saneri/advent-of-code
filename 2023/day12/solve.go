package main

import (
	"fmt"
	"saneri/aoc/utils"
	"strings"
)

func arraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func intArrToStr(arr []int) string {
	ret := ""
	for _, number := range(arr) {
		str := fmt.Sprint(number)
		ret += str
	}
	return ret
}

var cache = make(map[string]int)

func findArrangements(record string, springCounts []int) int {
	key := record + intArrToStr(springCounts)
	value, ok := cache[key] 
	if ok {
		return value
	}
	
	// not enough space for the springs
	if arraySum(springCounts) > len(record) {
		return 0
	}
	if (len(springCounts) == 0 ) {
		if (strings.Contains(record, "#")) {
			return 0
		}
		return 1
	}
	if len(record) == 0 {
		return 1
	}
	switch record[0] {
	case '.':
		return findArrangements(record[1:], springCounts)
	case '#':
		count := springCounts[0]
		if strings.Contains(record[:count], ".") || (len(record) > count && record[count] == '#') {
			return 0
		}
		if len(record) == count {
			return 1
		}
		// take one extra because # areas can't be next to each other
		return findArrangements(record[springCounts[0]+1:], springCounts[1:])
	default:
		test := findArrangements("."+record[1:], springCounts) + findArrangements("#"+record[1:], springCounts)
		cache[key] = test
		return test
	}
}

func repeatArray(arr []int, count int) []int {
	result := make([]int, 0)
	for i := 0; i < count; i++ {
		result = append(result, arr...)
	}
	return result
}

func main() {
	data := utils.ReadInput("input.txt")

	repeatRecords := func(record string, count int) string {
		result := []string{}
		for i := 0; i < count; i++ {
			result = append(result, record)
		}
		return strings.Join(result, "?")

	}

	sumA := 0
	sumB := 0
	for _, line := range data {
		split := strings.Split(line, " ")
		record := split[0]
		sprintCounts := utils.FindNumbers(split[1])
		sumA += findArrangements(repeatRecords(record, 1), repeatArray(sprintCounts, 1))
		sumB += findArrangements(repeatRecords(record, 5), repeatArray(sprintCounts, 5))
	}
	fmt.Println("a:", sumA)
	fmt.Println("b:", sumB)
}
