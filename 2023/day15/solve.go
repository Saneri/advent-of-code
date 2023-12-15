package main

import (
	"fmt"
	"saneri/aoc/utils"
	"strconv"
	"strings"
)

func hashAlgorithm(str string) int {
	value := 0
	for _, char := range str {
		value += int(char)
		value *= 17
		value %= 256
	}
	return value
}

func main() {
	data := strings.Split(utils.ReadInputString("input.txt"), ",")
	sumA := 0
	boxes := map[int][][2]string{}
	for _, step := range data {
		parts := strings.FieldsFunc(step, func(r rune) bool {
			return r == '=' || r == '-'
		})
		if len(parts) == 1 { // -
			label := parts[0]
			for val := 0; val <= 256; val++ {
				box, ok := boxes[val]
				if ok {
					for i, b := range box {
						if b[0] == label {
							boxes[val] = append(box[:i], box[i+1:]...)
							break
						}
					}
				}
			}
		} else { // =
			label := parts[0]
			value := hashAlgorithm(parts[0])
			focalLength := parts[1]
			found := false
			for val := 0; val <= 256; val++ {
				box, ok := boxes[val]
				if ok {
					for index, b := range box {
						if b[0] == label {
							box[index][1] = focalLength
							found = true
							break
						}
					}
				}
			}
			if !found {
				boxes[value] = append(boxes[value], [2]string{label, focalLength})
			}
		}
		sumA += hashAlgorithm(step)
	}
	fmt.Println("a:", sumA)

	sumB := 0
	for mapIndex, box := range boxes {
		if len(box) > 0 {
			for index, b := range box {
				focalLength, _ := strconv.Atoi(b[1])
				sumB += (mapIndex + 1) * (index + 1) * focalLength
			}
		}
	}
	fmt.Println("b:", sumB)
}
