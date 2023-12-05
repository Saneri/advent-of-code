package utils

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func ReadInput(filename string) []string {
	_, file, _, ok := runtime.Caller(1)
	if (!ok) {
		panic("Error with getting caller")
	}
	path := filepath.Dir(file)
	data, err := os.ReadFile(path + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}

func ReadInputString(filename string) string {
	_, file, _, ok := runtime.Caller(1)
	if (!ok) {
		panic("Error with getting caller")
	}
	path := filepath.Dir(file)
	data, err := os.ReadFile(path + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func FindNumbers(text string) []int {
	re := regexp.MustCompile(`\d+`)
	stringArray := re.FindAllString(text, -1)
	return ArrayStringToInt(stringArray)
}

func ArrayStringToInt(strArray []string) []int {
	intArray := []int{}
	for _, element := range strArray {
		number, err := strconv.Atoi(element)
		if err != nil {
			log.Fatal(err)
		}
		intArray = append(intArray, number)
	}
	return intArray
}
