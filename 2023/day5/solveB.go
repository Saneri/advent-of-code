package main

import (
	"fmt"
	"saneri/aoc/utils"
	"strings"
)

func convertWithMap(entry int, mapping [][]int) int {
	for _, mapEntry := range mapping {
		destination := mapEntry[0]
		source := mapEntry[1]
		rangeLength := mapEntry[2]
		if source <= entry && entry < source+rangeLength {
			return entry - source + destination
		}
	}
	return entry
}

func generateSeeds(values [][]int, mapping [][]int) [][]int {
	seeds := [][]int{}
	for _, seed := range values {
		value := seed[0]
		length := seed[1]
		for length > 0 {
			foundSplit := false
			minLength := length
			for _, mapEntry := range mapping {
				source := mapEntry[1]
				sourceLen := mapEntry[2]
				if source <= value && value < source+sourceLen {
					moveLen := min(source+sourceLen-value, length)
					seeds = append(seeds, []int{value, moveLen})
					length -= moveLen
					value += moveLen
					foundSplit = true
					break;
				} else if source > value {
					minLength = min(source - value, minLength)
				}
			}
			if !foundSplit {
				skipLength := min(length, minLength)
				seeds = append(seeds, []int{value, skipLength})
				value += skipLength
				length -= skipLength
			}
		}
	}
	return seeds
}

func main() {
	data := utils.ReadInputString("input.txt")
	sections := strings.Split(data, "\n\n")
	rawSeedData := utils.FindNumbers(sections[0])
	seeds := [][]int{}
	for len(rawSeedData) > 1 {
		seeds = append(seeds, []int{rawSeedData[0], rawSeedData[1]})
		rawSeedData = rawSeedData[2:]
	}
	for _, section := range sections[1:] {
		// parse mappings
		mappings := [][]int{}
		lines := strings.Split(section, "\n")[1:]
		for _, line := range lines {
			mapping := utils.FindNumbers((line))
			mappings = append(mappings, mapping)
		}

		// optimize seeds
		seeds = generateSeeds(seeds, mappings)

		// convert seeds
		nextGenSeeds := [][]int{}
		for _, seed := range seeds {
			newSeed := []int{convertWithMap(seed[0], mappings), seed[1]}
			nextGenSeeds = append(nextGenSeeds, newSeed)
		}
		seeds = nextGenSeeds

		}
		min := seeds[0][0]
		for _, value := range seeds {
			if value[0] < min {
				min = value[0]
			}
			if value[0] + value[1] -1 < min {
				min = value[0] + value[1] -1
			}
	}
	fmt.Println(min)
}
