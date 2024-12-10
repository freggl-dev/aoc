package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	result := 0

	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(string(content), -1)

	for _, m := range matches {
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])
		result += a * b
	}
	fmt.Println(result)
}
