package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	result := 0

	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	re := regexp.MustCompile(`(do|don't)\(\)|mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(string(content), -1)

	inst := "do"

	for _, m := range matches {
		if strings.TrimSpace(m[1]) == "do" || strings.TrimSpace(m[1]) == "don't" {
			inst = m[1]
			continue
		}

		if inst == "do" {
			a, _ := strconv.Atoi(m[2])
			b, _ := strconv.Atoi(m[3])
			result += a * b
		}
	}
	fmt.Println(result)
}
