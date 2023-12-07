package main

import (
	"fmt"
	"saneri/aoc/utils"
	"strings"
)

func convertWithMap(entry int, mapping[][]int) int {
	for _, mapEntry := range mapping {
		destination := mapEntry[0] 
		source := mapEntry[1]
		rangeLength := mapEntry[2] 
		if source <= entry && entry < source + rangeLength {
			return entry - source + destination
		}
	}
	return entry
}

func main() {
	data := utils.ReadInputString("input.txt")
	sections := strings.Split(data, "\n\n")
	seeds := utils.FindNumbers(sections[0])
	for _, section := range sections[1:] {
		mappings := [][]int{}
		lines := strings.Split(section, "\n")[1:]
		for _, line := range lines {
			mapping := utils.FindNumbers((line))
			mappings = append(mappings, mapping)
		}
		nextGenSeeds := []int{}
		for _, seed := range seeds {
			nextGenSeeds = append(nextGenSeeds, convertWithMap(seed, mappings))
		}
		seeds = nextGenSeeds
	}
	min := seeds[0]
	for _, value := range seeds {
		if value < min {
			min = value
		}
	}
	fmt.Println(min)
}