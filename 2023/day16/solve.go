package main

import (
	"fmt"
	"saneri/aoc/utils"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Beam struct {
	coordinates [2]int
	direction 	Direction
}

func convertSlash(direction Direction) Direction {
	convert := map[Direction]Direction{
		Up:    Right,
		Right: Up,
		Down:  Left,
		Left:  Down,
	}
	return convert[direction]
}

func convertBakslash(direction Direction) Direction {
	convert := map[Direction]Direction{
		Up:    Left,
		Right: Down,
		Down:  Right,
		Left:  Up,
	}
	return convert[direction]
}

func calculateBeams(grid []string, startingBeam Beam) int {
	width := len(grid[0])
	height := len(grid)
	visited := make(map[Beam]bool)
	var getNextBeams func(beam Beam)
	getNextBeams = func(beam Beam) {
		if visited[beam] {
			return
		}
		visited[beam] = true

		x := beam.coordinates[0]
		y := beam.coordinates[1]

		switch beam.direction {
		case Up:
			y--
		case Right:
			x++
		case Down:
			y++
		case Left:
			x--
		}
		if x < 0 || y < 0 || x >= width || y >= height {
			return
		}
		char := grid[y][x]
		coordinates := [2]int{x, y}
		switch char {
		case '.':
			getNextBeams(Beam{coordinates, beam.direction})
		case '/':
			getNextBeams(Beam{coordinates, convertSlash(beam.direction)})
		case '\\':
			getNextBeams(Beam{coordinates, convertBakslash(beam.direction)})
		case '-':
			if beam.direction == Right || beam.direction == Left {
				getNextBeams(Beam{coordinates, beam.direction})
			} else {
				getNextBeams(Beam{coordinates, Right})
				getNextBeams(Beam{coordinates, Left})
			}
		case '|':
			if beam.direction == Up || beam.direction == Down {
				getNextBeams(Beam{coordinates, beam.direction})
			} else {
				getNextBeams(Beam{coordinates, Up})
				getNextBeams(Beam{coordinates, Down})
			}
		}
	}
	getNextBeams(startingBeam)

	sum := 0
	for y, line := range grid {
		for x := range line {
			if visited[Beam{[2]int{x, y}, Up}] || visited[Beam{[2]int{x, y}, Right}] || visited[Beam{[2]int{x, y}, Down}] || visited[Beam{[2]int{x, y}, Left}] {
				sum++
			}
		}
	}
	return sum
}

func main() {
	data := utils.ReadInput("input.txt")
	width := len(data[0])
	height := len(data)
	currentBeam := Beam{[2]int{-1, 0}, Right}
	maxEnergy := 0
	for y := range data {
		maxEnergy = max(calculateBeams(data, Beam{[2]int{-1, y}, Right}), maxEnergy)
		maxEnergy = max(calculateBeams(data, Beam{[2]int{width, y}, Left}), maxEnergy)
	}
	for x := range data[0] {
		maxEnergy = max(calculateBeams(data, Beam{[2]int{x, -1}, Down}), maxEnergy)
		maxEnergy = max(calculateBeams(data, Beam{[2]int{x, height}, Up}), maxEnergy)
	}
	fmt.Println("a:", calculateBeams(data, currentBeam))
	fmt.Println("b:", maxEnergy)
}
