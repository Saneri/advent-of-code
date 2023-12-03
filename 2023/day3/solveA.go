package main

import (
	"fmt"
	"log"
	"saneri/aoc/utils"
	"strconv"
	"unicode"
)

func isNumber(number rune) bool {
	if unicode.IsDigit(number) {
		return true
	}
	return false
}

func isSymbol(symbol rune) bool {
	if !unicode.IsDigit(symbol) && symbol != '.' {
		return true
	}
	return false
}

func hasSymbolsNearby(x int, y int, data []string) bool {
	height := len(data) - 1
	width := len(data[0]) - 1
	if x > 0 && isSymbol(rune(data[y][x-1])) {
		return true
	}
	if x < width && isSymbol(rune(data[y][x+1])) {
		return true
	}
	if y > 0 && isSymbol(rune(data[y-1][x])) {
		return true
	}
	if y < height && isSymbol(rune(data[y+1][x])) {
		return true
	}
	if x > 0 && y > 0 && isSymbol(rune(data[y-1][x-1])) {
		return true
	}
	if x < width && y > 0 && isSymbol(rune(data[y-1][x+1])) {
		return true
	}
	if x > 0 && y < height && isSymbol(rune(data[y+1][x-1])) {
		return true
	}
	if x < width && y < height && isSymbol(rune(data[y+1][x+1])) {
		return true
	}
	return false
}

func main() {
	data := utils.ReadInput("input.txt")
	sum := 0
	for y, line := range data {
		number := ""
		shouldBeCounted := false
		for x, char := range line {
			if isNumber(char) {
				number += string(char)
				if hasSymbolsNearby(x, y, data) {
					shouldBeCounted = true
				}
			}
			if !isNumber(char) || x == len(line)-1 {
				if shouldBeCounted && len(number) > 0 {
					value, err := strconv.Atoi(number)
					if err != nil {
						log.Fatal(err)
					}
					sum += value
				}
				number = ""
				shouldBeCounted = false
			}
		}
	}
	fmt.Println(sum)
}
