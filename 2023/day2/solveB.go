package main

import (
	"saneri/aoc/utils"
	"fmt"
	"log"
	"strings"
	"strconv"
)

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
	data := utils.ReadInput("input.txt")
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
		sum += maxValues["red"] * maxValues["green"] * maxValues["blue"]
	}
	fmt.Println(sum)
}