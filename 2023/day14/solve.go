package main

import (
	"fmt"
	"saneri/aoc/utils"
	"strings"
)

const (
	NORTH int = 0
	WEST  int = 1
	SOUTH int = 2
	EAST  int = 3
)

func moveRocks(grid [][]rune, direction int) ([][]rune, bool) {
	rocksMoved := false
	for y, row := range grid {
		for x, char := range row {
			if char == '.' {
				switch direction {
				case NORTH:
					if y+1 < len(grid) && grid[y+1][x] == 'O' {
						grid[y][x] = 'O'
						grid[y+1][x] = '.'
						rocksMoved = true
					}
				case EAST:
					if x+1 < len(row) && grid[y][x+1] == 'O' {
						grid[y][x] = 'O'
						grid[y][x+1] = '.'
						rocksMoved = true
					}
				case SOUTH:
					if y > 0 && grid[y-1][x] == 'O' {
						grid[y][x] = 'O'
						grid[y-1][x] = '.'
						rocksMoved = true
					}
				case WEST:
					if x > 0 && grid[y][x-1] == 'O' {
						grid[y][x] = 'O'
						grid[y][x-1] = '.'
						rocksMoved = true
					}
				}
			}
		}
	}
	return grid, rocksMoved
}

func tiltPlatform(grid [][]rune, direction int) [][]rune {
	rocksMoved := true
	for rocksMoved {
		grid, rocksMoved = moveRocks(grid, direction)
	}
	return grid
}

func calcTotalLoad(grid [][]rune) int {
	sum := 0
	length := len(grid)
	for i, row := range grid {
		sum += strings.Count(string(row), "O") * (length - i)
	}
	return sum
}

func gridToString(grid [][]rune) string {
	rows := []string{}
	for _, row := range grid {
		rows = append(rows, string(row))
	}
	return strings.Join(rows, "\n")
}

func stringToGrid(str string) [][]rune {
	rows := strings.Split(str, "\n")
	grid := make([][]rune, len(rows))
	for i, row := range rows {
		grid[i] = []rune(row)
	}
	return grid
}

func main() {
	data := utils.ReadInput("input.txt")
	grid := make([][]rune, len(data))
	for i, line := range data {
		grid[i] = []rune(line)
	}
	cache := make(map[string]int)
	cycles := 1000000000
	for i := 0; i < cycles; i++ {
		grid = tiltPlatform(grid, NORTH)
		grid = tiltPlatform(grid, EAST)
		grid = tiltPlatform(grid, SOUTH)
		grid = tiltPlatform(grid, WEST)
		cycle, hit := cache[gridToString(grid)]
		
		if hit {
			remain := cycles - i -1
			modulo := remain % (i - cycle)
			for key, index := range cache {
				if (index == modulo + cycle) {
					fmt.Println(calcTotalLoad(stringToGrid(key)))
					break
				}
			}
			break
		}
		cache[gridToString(grid)] = i		
	}
}
