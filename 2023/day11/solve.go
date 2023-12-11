package main

import (
	"fmt"
	"saneri/aoc/utils"
)

func findEmptyLines(grid []string) (map[int]bool, map[int]bool) {
	emptyRows := map[int]bool{}
	emptyCols := map[int]bool{}

	for y := range grid {
		empty := true
		for x := range grid[0] {
			if grid[y][x] != '.' {
				empty = false
				break
			}
		}
		emptyRows[y] = empty
	}
	for x := range grid[0] {
		empty := true
		for y := range grid {
			if grid[y][x] != '.' {
				empty = false
				break
			}
		}
		emptyCols[x] = empty
	}
	return emptyRows, emptyCols
}

func findGalaxies(grid []string) [][2]int {
	galaxies := [][2]int{}
	for y := range grid {
		for x := range grid[0] {
			if grid[y][x] == '#' {
				galaxies = append(galaxies, [2]int{x, y})
			}
		}
	}
	return galaxies
}

func main() {
	data := utils.ReadInput("input.txt")
	emptyRows, emptyCols := findEmptyLines(data)

	manhattanDistance := func(galaxy1, galaxy2 [2]int, expansion int) int {
		xdist := 0
		ydist := 0
		x1, x2 := galaxy1[0], galaxy2[0]
		if galaxy1[0] > galaxy2[0] {
			x1, x2 = galaxy2[0], galaxy1[0]
		}
		for x := x1; x < x2; x++ {
			if emptyCols[x] {
				xdist += expansion
			} else {
				xdist++
			}
		}
		y1, y2 := galaxy1[1], galaxy2[1]
		if galaxy1[1] > galaxy2[1] {
			y1, y2 = galaxy2[1], galaxy1[1]
		}
		for y := y1; y < y2; y++ {
			if emptyRows[y] {
				ydist += expansion
			} else {
				ydist++
			}
		}
		return ydist + xdist
	}

	galaxies := findGalaxies(data)
	countDist := func(expansion int) int {
		totalDist := 0
		visited := map[[2]int]bool{}
		for _, galaxy1 := range galaxies {
			visited[galaxy1] = true
			for _, galaxy2 := range galaxies {
				if visited[galaxy2] {
					continue
				}
				totalDist += manhattanDistance(galaxy1, galaxy2, expansion)
			}
		}
		return totalDist
	}
	fmt.Println("a:", countDist(2))
	fmt.Println("b:", countDist(1000000))
}
