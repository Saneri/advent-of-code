package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func readInput() []string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}

func updateMaxValue(count int, color string, maxValues map[string]int) {
	switch (color) {
	case "red":
		if (count > maxValues["red"]) {
			maxValues["red"] = count
		}
		break
	case "green":
		if (count > maxValues["green"]) {
			maxValues["green"] = count
		}
		break
	case "blue":
		if (count > maxValues["blue"]) {
			maxValues["blue"] = count
		}
		break
	default: 
		log.Fatal("unknown color: ", color)
	} 
}


func main() {
	data := readInput()
	sum := 0
	for _, line := range data {
		splitByColon := strings.Split(line, ": ")
		sets := strings.Split(splitByColon[1], "; ")
		maxValues := map[string]int{
			"red": 0,
			"green": 0,
			"blue": 0,
		}
		for _, set := range sets {
			balls := strings.Split(set, ", ")
			for _, ball := range balls {
				splitBySpace := strings.Split(ball, " ")
				count, err := strconv.Atoi(splitBySpace[0])
				if (err != nil) {
					log.Fatal(err)
				}
				color := splitBySpace[1]
				updateMaxValue(count, color, maxValues)
			}
			
		}
		fmt.Println(maxValues)
		sum += maxValues["red"] * maxValues["green"] * maxValues["blue"]
	}
	fmt.Println(sum)
}