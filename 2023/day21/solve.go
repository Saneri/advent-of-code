package main

import (
	"fmt"
	"log"
	"saneri/aoc/utils"
)

type point struct {
	x, y int
}

func findStart(grid []string) point {
	for y, line := range grid {
		for x, char := range line {
			if char == 'S' {
				grid[y] = grid[y][:x] + "." + grid[y][x+1:]
				return point{x, y}
			}
		}
	}
	log.Fatal("No start found")
	return point{-1, -1}
}

func neighbours(p point, grid []string) []point {
	height := len(grid)
	width := len(grid[0])
	neighbours := []point{}
	if p.x > 0 && grid[p.y][p.x-1] == '.' {
		neighbours = append(neighbours, point{p.x - 1, p.y})
	}
	if p.x < width-1 && grid[p.y][p.x+1] == '.' {
		neighbours = append(neighbours, point{p.x + 1, p.y})
	}
	if p.y > 0 && grid[p.y-1][p.x] == '.' {
		neighbours = append(neighbours, point{p.x, p.y - 1})
	}
	if p.y < height-1 && grid[p.y+1][p.x] == '.' {
		neighbours = append(neighbours, point{p.x, p.y + 1})
	}
	return neighbours
}

func printGrid(grid []string, points map[point]bool) {
	for y, line := range grid {
		for x, char := range line {
			if char == '.' {
				if points[point{x, y}] {
					fmt.Print("O")
				} else {
					fmt.Print(".")
				}
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Println()
	}
	fmt.Println()

}

func main() {
	data := utils.ReadInput("input.txt")
	start := findStart(data)
	rounds := 64
	startPoints := map[point]bool{start: true}
	endPoints := map[point]bool{}
	for i := 0; i < rounds; i++ {
		for point := range startPoints {
			for _, n := range neighbours(point, data) {
				endPoints[n] = true
			}
		}
		// printGrid(data, endPoints)
		startPoints = endPoints
		endPoints = map[point]bool{}
	}
	fmt.Println("a:", len(startPoints))

}
