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

func validate(count int, color string) bool {
	switch (color) {
	case "red":
		if (count > 12) {
			return false
		}
	case "green":
		if (count > 13) {
			return false
		}
	case "blue":
		if (count > 14) {
			return false
		}
	default: 
		log.Fatal("unknown color: ", color)
	} 
	return true
}


func main() {
	data := readInput()
	sum := 0
	for _, line := range data {
		splitByColon := strings.Split(line, ": ")
		gameId := strings.Split(splitByColon[0], " ")[1]
		sets := strings.Split(splitByColon[1], "; ")
		allValid := true
		for _, set := range sets {
			balls := strings.Split(set, ", ")
			for _, ball := range balls {
				splitBySpace := strings.Split(ball, " ")
				count, err := strconv.Atoi(splitBySpace[0])
				if (err != nil) {
					log.Fatal(err)
				}
				color := splitBySpace[1]
				if (!validate(count, color)) {
					allValid = false
					break
				}
			}
		}
		if (allValid) {
			id, err := strconv.Atoi(gameId)
			if (err != nil) {
				log.Fatal(err)
			}
			sum += id
		}
	}
	fmt.Println(sum)
}