package main

import (
	"fmt"
	"log"
	"maps"
	"saneri/aoc/utils"
	"strconv"
	"strings"
)

func createPartMap(part string) map[string]int {
	trimmed := strings.Trim(part, "{}")
	ratings := strings.Split(trimmed, ",")
	partMap := make(map[string]int)
	for _, rating := range ratings {
		split := strings.Split(rating, "=")
		value, error := strconv.Atoi(split[1])
		if error != nil {
			panic(error)
		}
		partMap[split[0]] = value
	}
	return partMap
}

func getNextStep(current string, partMap map[string]int) string {
	rules := strings.Split(current, ",")
	for _, rule := range rules[:len(rules)-1] {
		split := strings.Split(rule, ":")

		// had to try an inline function at some point :D
		check := strings.FieldsFunc(split[0], func(r rune) bool {
			return r == '<' || r == '>'
		})
		value, error := strconv.Atoi(check[1])
		if error != nil {
			panic(error)
		}
		rating := partMap[check[0]]
		if strings.Contains(rule, "<") && rating < value {
			return split[1]
		} else if strings.Contains(rule, ">") && rating > value {
			return split[1]
		}
	}
	otherwise := rules[len(rules)-1]
	return otherwise
}

func addXmasRating(sum int, xmasRating map[string]int) int {
	for _, rating := range xmasRating {
		sum += rating
	}
	return sum
}

type Interval [2]int

func createIntervalPartMap() map[string]Interval {
	initValue := Interval{1, 4000}
	return map[string]Interval{"x": initValue, "m": initValue, "a": initValue, "s": initValue}
}

func main() {
	data := utils.ReadInputString("input.txt")
	split := strings.Split(data, "\n\n")
	workflowLines := strings.Split(split[0], "\n")
	workflows := make(map[string]string)
	for _, line := range workflowLines {
		parts := strings.Split(line, "{")
		workflows[parts[0]] = parts[1][:len(parts[1])-1]
	}

	parts := strings.Split(split[1], "\n")
	sum := 0
	for part := range parts {
		partMap := createPartMap(parts[part])
		current := workflows["in"]
		found := false
		for !found {
			next := getNextStep(current, partMap)
			if next == "A" {
				sum = addXmasRating(sum, partMap)
				found = true
			} else if next == "R" {
				found = true
			} else {
				current = workflows[next]
			}
		}
	}
	fmt.Println("a:", sum)

	partMap := createIntervalPartMap()
	acceptedPartMaps := findAcceptedParts(workflows["in"], partMap, workflows)

	sumB := 0
	for _, acceptedPartMap := range acceptedPartMaps {
		x := acceptedPartMap["x"]
		m := acceptedPartMap["m"]
		a := acceptedPartMap["a"]
		s := acceptedPartMap["s"]
		sumB += (x[1] - x[0]+1) * (m[1] - m[0]+1) * (a[1] - a[0]+1) * (s[1] - s[0]+1)
	}
	fmt.Println("b:", sumB)
}

func findAcceptedParts(current string, partMap map[string]Interval, workflows map[string]string) []map[string]Interval {
	if strings.Contains(current, ":") {
		split := strings.SplitN(current, ",", 2)
		otherwise := split[1]
		split2 := strings.Split(split[0], ":")
		rule := split2[0]
		next := split2[1]
		check := strings.FieldsFunc(rule, func(r rune) bool {
			return r == '<' || r == '>'
		})
		value, error := strconv.Atoi(check[1])
		if error != nil {
			panic(error)
		}
		nextWorkflow, ok := workflows[next]
		if !ok {
			if next == "A" || next == "R" {
				nextWorkflow = next
			} else {
				log.Fatal("no workflow found for ", next)
			}
		}
		low, high := partMap[check[0]][0], partMap[check[0]][1]
		if low > high {
			log.Fatal("low > high")
		}

		if (strings.Contains(rule, "<") && (high < value || low >= value)) || (strings.Contains(rule, ">") && (low > value || high <= value)){
			if strings.Contains(rule, "<") {
				return findAcceptedParts(nextWorkflow, partMap, workflows)
			} else if strings.Contains(rule, ">") {
				return findAcceptedParts(otherwise, partMap, workflows)
			} else {
				log.Fatal("no < or > in rule")
			}
		}

		highMap := make(map[string]Interval)
		maps.Copy(highMap, partMap)
		lowMap := make(map[string]Interval)
		maps.Copy(lowMap, partMap)

		if strings.Contains(rule, "<") {
			highMap[check[0]] = Interval{value, high}
			lowMap[check[0]] = Interval{low, value - 1}
			return append(findAcceptedParts(nextWorkflow, lowMap, workflows), findAcceptedParts(otherwise, highMap, workflows)...)
		} else if strings.Contains(rule, ">") {
			highMap[check[0]] = Interval{value + 1, high}
			lowMap[check[0]] = Interval{low, value}
			return append(findAcceptedParts(nextWorkflow, highMap, workflows), findAcceptedParts(otherwise, lowMap, workflows)...)
		}
		log.Fatal("no < or > in rule")
	}

	switch current {
	case "A":
		return []map[string]Interval{partMap}
	case "R":
		return nil
	default:
		nextWorkflow, ok := workflows[current]
		if !ok {
			log.Fatal("no workflow found for ", current)
		}
		return findAcceptedParts(nextWorkflow, partMap, workflows)
	}
}
