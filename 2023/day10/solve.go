package main

import (
	"fmt"
	"log"
	"saneri/aoc/utils"
	"strings"
)

func findChar(grid []string, c rune) [2]int {
	for y, line := range grid {
		for x, char := range line {
			if char == c {
				return [2]int{x, y}
			}
		}
	}
	log.Fatal("No S found")
	return [2]int{}
}

func findStartAndEnd(grid []string) ([2]int, [2]int) {
	s := findChar(grid, 'S')
	visited := make(map[string]bool)
	start := getNextPipe(grid, s, "URDL", visited)
	visited[pipeToStr(start)] = true
	end := getNextPipe(grid, s, "URDL", visited)
	return start, end
}

func getNextPipe(grid []string, current [2]int, check string, visited map[string]bool) [2]int {
	allowed := map[string]string{
		"U": "F|7",
		"R": "-7J",
		"D": "|JL",
		"L": "-FL",
	}
	gridWidth := len(grid[0])
	gridHeight := len(grid)
	if strings.Contains(check, "U") && current[1] > 0 {
		nextCoord := [2]int{current[0], current[1] - 1}
		nextChar := grid[current[1]-1][current[0]]
		if !visited[pipeToStr(nextCoord)] && strings.Contains(allowed["U"], string(nextChar)) {
			return nextCoord
		}
	}
	if strings.Contains(check, "R") && current[0] < gridWidth {
		nextCoord := [2]int{current[0] + 1, current[1]}
		nextChar := grid[current[1]][current[0]+1]
		if !visited[pipeToStr(nextCoord)] && strings.Contains(allowed["R"], string(nextChar)) {
			return nextCoord
		}
	}
	if strings.Contains(check, "D") && current[1] < gridHeight {
		nextCoord := [2]int{current[0], current[1] + 1}
		nextChar := grid[current[1]+1][current[0]]
		if !visited[pipeToStr(nextCoord)] && strings.Contains(allowed["D"], string(nextChar)) {
			return nextCoord
		}
	}
	if strings.Contains(check, "L") && current[0] > 0 {
		nextCoord := [2]int{current[0] - 1, current[1]}
		nextChar := grid[current[1]][current[0]-1]
		if !visited[pipeToStr(nextCoord)] && strings.Contains(allowed["L"], string(nextChar)) {
			return nextCoord
		}
	}
	log.Fatal("No next pipe found from", current)
	return [2]int{}
}

func pipeToStr(coord [2]int) string {
	return fmt.Sprint(coord[0]) + "," + fmt.Sprint(coord[1])
}

func getMazeLength(grid []string, start [2]int, end [2]int) int {
	visited := make(map[string]bool)
	current := start
	length := 2
	for current != end {
		visited[pipeToStr(current)] = true
		char := grid[current[1]][current[0]]
		switch char {
		case 'F':
			current = getNextPipe(grid, current, "RD", visited)
		case '-':
			current = getNextPipe(grid, current, "LR", visited)
		case '7':
			current = getNextPipe(grid, current, "LD", visited)
		case '|':
			current = getNextPipe(grid, current, "UD", visited)
		case 'J':
			current = getNextPipe(grid, current, "LU", visited)
		case 'L':
			current = getNextPipe(grid, current, "UR", visited)
		default:
			log.Fatal("Unknown char", string(char))
		}
		length++
	}
	return length
}

func main() {
	data := utils.ReadInput("input.txt")
	start, end := findStartAndEnd(data)
	length := getMazeLength(data, start, end)
	fmt.Println(length / 2)
}
