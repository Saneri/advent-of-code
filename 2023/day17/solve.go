package main

import (
	"container/heap"
	"fmt"
	"log"
	"saneri/aoc/utils"
)

type PriorityQueue []*Node

func (piq PriorityQueue) Len() int {
	return len(piq)
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
	heap.Fix(pq, pq.Len()-1)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heat < pq[j].heat
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

type Node struct {
	point     Point
	prevPoint Point
	heat      int
}

type Point struct {
	x, y int
}

type Grid struct {
	values [][]int
	height int
	width  int
}

type Direction int


func getAdjecentNodes(grid Grid, current Node, least int, most int) []Node {
	ret := []Node{}
	x, y := current.point.x, current.point.y
	directions := []Point{}
	// add the points that move vertically
	if current.prevPoint.x != x {
		directions = append(directions, Point{0, -1}, Point{0, 1})
	}
	// add the points that move horizontally
	if current.prevPoint.y != y {
		directions = append(directions, Point{-1, 0}, Point{1, 0})
	}

	for _, dir := range directions {
		totalHeat := current.heat
		for i := 1; i <= most; i++ {
			newX, newY := x+dir.x*i, y+dir.y*i
			if newX >= 0 && newX < grid.width && newY >= 0 && newY < grid.height {
				totalHeat += grid.values[newY][newX]
				if i >= least {
					ret = append(ret, Node{Point{newX, newY}, Point{x, y}, totalHeat})
				}
			}
		}
	}
	return ret
}

func dijkstras(grid Grid, start Point, end Point, minStep int, maxStep int) int {
	queue := make(PriorityQueue, 0)
	node := Node{start, Point{start.x - 1, start.y - 1}, 0}
	queue.Push(&node)
	visitedGraphs := make(map[Node]bool)
	for len(queue) > 0 {
		current := heap.Pop(&queue).(*Node)
		if current.point == end {
			return current.heat
		}
		if _, ok := visitedGraphs[Node{current.point, current.prevPoint, 0}]; ok {
			continue
		}
		visitedGraphs[Node{current.point, current.prevPoint, 0}] = true
		for _, node := range getAdjecentNodes(grid, *current, minStep, maxStep) {
			nodeCopy := node
			queue.Push(&nodeCopy)
		}
	}
	log.Fatal("No path found")
	return -1
}

func gridToInts(grid []string) [][]int {
	ret := [][]int{}
	for _, line := range grid {
		chars := []string{}
		for _, char := range line {
			chars = append(chars, string(char))
		}
		ret = append(ret, utils.ArrayStringToInt(chars))
	}
	return ret
}

func main() {
	data := utils.ReadInput("input.txt")
	grid := Grid{gridToInts(data), len(data), len(data[0])}
	fmt.Println("a", dijkstras(grid, Point{0, 0}, Point{grid.width - 1, grid.height - 1}, 1, 3))
	fmt.Println("b", dijkstras(grid, Point{0, 0}, Point{grid.width - 1, grid.height - 1}, 4, 10))
}
