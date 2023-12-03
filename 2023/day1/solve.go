package main

import (
	"saneri/aoc/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func stringToNumber(str string) int {
	wordToNumber := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}

	if (len(str) > 1) {
		number := wordToNumber[str]
		if (number == 0){
			log.Fatal("fatal: ", str)
		}
		return number
	}
	number, err := strconv.Atoi(str)
	if (err != nil) {
		log.Fatal(err)
	}
	return number
}

func main() {
	data := utils.ReadInput("input.txt")
	sum := 0
	re := regexp.MustCompile("[1-9]|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)")
	for _, line := range data {
		// had to create this monster for loop because I found no ways of finding overlapping matches in go regex
		digits := []string{}
		for (len(line) > 0) {
			str := re.FindString(line)
			if (len(strings.TrimSpace(str)) > 0) {
				digits = append(digits, str)
			}
			line = line[1:]
		}
		first := stringToNumber(digits[0])
		second := stringToNumber(digits[len(digits)-1])
		sum += first * 10 + second
	}
	fmt.Println(sum)
}