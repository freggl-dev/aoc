package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//result := 0
	puzzle := make([][]string, 140)

	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()
	i := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		puzzle[i] = append(puzzle[i], scanner.Text())
		i++
	}
	fmt.Println(puzzle)
}
