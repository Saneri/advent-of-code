package main

import (
	"fmt"
	"saneri/aoc/utils"
	"strings"
)

func isHorizontalReflection(pattern []string, index int) bool {
	width := len(pattern[0])
	length := len(pattern)
	smudgeUsed := false
	for start, end := index, index+1; start >= 0 && end < length; start, end = start-1, end+1 {
		for i := 0; i < width; i++ {
			if pattern[start][i] != pattern[end][i] {
				if smudgeUsed {
					return false
				}
				smudgeUsed = true
			}
		}
	}
	return smudgeUsed
}

func isVerticalReflection(pattern []string, index int) bool {
	width := len(pattern[0])
	length := len(pattern)
	smudgeUsed := false
	for start, end := index, index+1; start >= 0 && end < width; start, end = start-1, end+1 {
		for i := 0; i < length; i++ {
			if pattern[i][start] != pattern[i][end] {
				if smudgeUsed {
					return false
				}
				smudgeUsed = true
			}
		}
	}
	return smudgeUsed
}

func findReflection(pattern []string) int {
	for index := range pattern[1:] {
		if isHorizontalReflection(pattern, index) {
			fmt.Println("horizontal", index+1)
			return (index + 1) * 100
		}
	}
	for index := range pattern[0][1:] {
		if isVerticalReflection(pattern, index) {
			fmt.Println("vertical", index+1)
			return index + 1
		}
	}
	return 0
}

func main() {
	data := utils.ReadInputString("input.txt")
	patterns := strings.Split(data, "\n\n")

	sum := 0
	for _, pattern := range patterns {
		sum += findReflection(strings.Split(pattern, "\n"))
	}
	fmt.Println(sum)
}
