package utils

import (
	"os"
	"log"
	"path/filepath"
	"runtime"
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