package main

import (
	"fmt"
	"log"
	"saneri/aoc/utils"
	"slices"
	"strconv"
	"unicode"
)

func isNumber(number rune) bool {
	if unicode.IsDigit(number) {
		return true
	}
	return false
}

func isGear(symbol rune) bool {
	if symbol == '*' {
		return true
	}
	return false
}

func addGear(x int, y int, gearSymbols map[string][]int, numberIndex int) {
	coordinateStr := fmt.Sprintf("%d,%d", x, y)
	arr := gearSymbols[coordinateStr]
	if !slices.Contains(arr, numberIndex) {
		gearSymbols[coordinateStr] = append(arr, numberIndex)
	}
}

func hasSymbolsNearby(x int, y int, data []string, gearSymbols map[string][]int, numberIndex int) {
	height := len(data) - 1
	width := len(data[0]) - 1
	if x > 0 && isGear(rune(data[y][x-1])) {
		addGear(x-1, y, gearSymbols, numberIndex)
	}
	if x < width && isGear(rune(data[y][x+1])) {
		addGear(x+1, y, gearSymbols, numberIndex)
	}
	if y > 0 && isGear(rune(data[y-1][x])) {
		addGear(x, y-1, gearSymbols, numberIndex)
	}
	if y < height && isGear(rune(data[y+1][x])) {
		addGear(x, y+1, gearSymbols, numberIndex)
	}
	if x > 0 && y > 0 && isGear(rune(data[y-1][x-1])) {
		addGear(x-1, y-1, gearSymbols, numberIndex)
	}
	if x < width && y > 0 && isGear(rune(data[y-1][x+1])) {
		addGear(x+1, y-1, gearSymbols, numberIndex)

	}
	if x > 0 && y < height && isGear(rune(data[y+1][x-1])) {
		addGear(x-1, y+1, gearSymbols, numberIndex)

	}
	if x < width && y < height && isGear(rune(data[y+1][x+1])) {
		addGear(x+1, y+1, gearSymbols, numberIndex)

	}
}

func main() {
	data := utils.ReadInput("input.txt")

	// each gears coordinates and the amount of adjecent items
	// coordinate (x=1, y=2) would be simply "1,2" mapped to numberIndeces that has seen the gear
	gearSymbols := map[string][]int{}
	values := map[int]int{}
	numberIndex := 0
	for y, line := range data {
		number := ""
		for x, char := range line {
			if isNumber(char) {
				number += string(char)
				hasSymbolsNearby(x, y, data, gearSymbols, numberIndex)
			}
			if !isNumber(char) || x == len(line)-1 {
				if len(number) > 0 {
					value, err := strconv.Atoi(number)
					if err != nil {
						log.Fatal(err)
					}
					number = ""
					values[numberIndex] = value
					numberIndex++
				}
			}
		}
	}
	sum := 0
	for _, value := range gearSymbols {
		if len(value) == 2 {
			sum += values[value[0]] * values[value[1]]
		}
	}
	fmt.Println(sum)
}
