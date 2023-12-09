package main

import (
	"fmt"
	"log"
	"regexp"
	"saneri/aoc/utils"
	"strings"
)

// greates common factor
func gcd(a int, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

// least common multiple
func lcm(arr []int) int {
	ans := arr[0]
	for _, element := range arr {
		ans = ((element * ans) / (gcd(element, ans)))
	}
	return ans
}

func findWords(text string) []string {
	re := regexp.MustCompile(`\w+`)
	return re.FindAllString(text, -1)
}

func main() {
	data := utils.ReadInputString("input.txt")
	split := strings.Split(data, "\n\n")
	instructions := split[0]
	program := strings.Split(split[1], "\n")
	mapping := map[string][]string{}
	for _, line := range program {
		node := findWords(line)
		mapping[node[0]] = node[1:]
	}
	currentNodes := []string{}
	for key := range mapping {
		if key[2] == 'A' {
			currentNodes = append(currentNodes, key)
		}
	}
	goal := "Z"
	steps := make([]int, len(currentNodes))

	for index, node := range currentNodes {
		found := false
		for !found {
			for _, instruction := range instructions {
				steps[index]++
				if string(instruction) == "L" {
					node = mapping[node][0]
				} else if string(instruction) == "R" {
					node = mapping[node][1]
				} else {
					log.Fatal("Invalid instruction ", string(instruction))
				}
				if string(node[2]) == goal {
					found = true
					break
				}
			}
		}
	}
	fmt.Println(lcm(steps))
}
