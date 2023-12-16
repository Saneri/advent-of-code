package main

import (
	"fmt"
	"log"
	"maps"
	"saneri/aoc/utils"
	"strconv"
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

func isInsideLoop(grid []string, point [2]int) bool {
	x := point[0]
	y := point[1]
	amountOfWalls := 0
	amountOfF7 := 0
	amountOfLJ := 0
	for i := x; i >= 0; i-- {
		switch grid[y][i] {
		case '|':
			amountOfWalls++
		case '7':
			amountOfF7++
		case 'F':
			amountOfF7++
		case 'L':
			amountOfLJ++
		case 'J':
			amountOfLJ++
		}
	}
	amountOfWalls += min(amountOfF7, amountOfLJ)
	return amountOfWalls % 2 == 1
}

func findEnclosedTiles(tiles map[string]bool, grid []string, current [2]int, visited map[string]bool) (map[string]bool, bool) {
	gridWidth := len(grid[0])
	gridHeight := len(grid)
	x := current[0]
	y := current[1]
	tiles[pipeToStr([2]int{x, y})] = true
	enclosed := true
	if x == 0 || x == gridWidth-1 || y == 0 || y == gridHeight-1 || !isInsideLoop(grid, current) {
		return tiles, false
	}

	for _, dir := range [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
		newX, newY := x+dir[0], y+dir[1]
		if !visited[pipeToStr([2]int{newX, newY})] && !tiles[pipeToStr([2]int{newX, newY})] {
			_, wasEnclosed := findEnclosedTiles(tiles, grid, [2]int{newX, newY}, visited)
			if !wasEnclosed {
				enclosed = false
			}
		}
	}
	return tiles, enclosed
}

func removeJunk(grid []string, visited map[string]bool) []string {
	newGrid := make([]string, len(grid))
    copy(newGrid, grid)

	for y, line := range newGrid {
        for x := range line {
			_, ok := visited[pipeToStr([2]int{x, y})]
            if !ok {
                newGrid[y] = newGrid[y][:x] + "." + newGrid[y][x+1:]
            }
        }
    }
    return newGrid
}

func getMazeLength(grid []string, start [2]int, end [2]int) (int, int) {
	sSpot := findChar(grid, 'S')
	visited := map[string]bool{pipeToStr(sSpot): true}
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
	visited[pipeToStr(end)] = true

	grid = removeJunk(grid, visited)

	gridWidth := len(grid[0])
	gridHeight := len(grid)
	enclosedTiles := make(map[string]bool)
	for tile := range visited {
		split := strings.Split(tile, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		
		for _, dir := range [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			newX, newY := x+dir[0], y+dir[1]
			if newX >= 0 && newY >= 0 && newX < gridWidth && newY < gridHeight && grid[newY][newX] == '.' {
				newTiles, areEnclosed := findEnclosedTiles(maps.Clone(enclosedTiles), grid, [2]int{newX, newY}, visited)
				if areEnclosed {
					maps.Copy(enclosedTiles, newTiles)
				}
			}
		}
	}
	return length, len(enclosedTiles)
}

func main() {
	data := utils.ReadInput("input.txt")
	start, end := findStartAndEnd(data)
	length, tiles := getMazeLength(data, start, end)
	fmt.Println("a:", length / 2)
	fmt.Println("b:", tiles)
}
