package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// vars
	var left []int
	var right []int
	var result int

	// open file
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	// read file and write left and right list
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		l, _ := strconv.Atoi(line[0])
		r, _ := strconv.Atoi(line[1])
		left = append(left, l)
		right = append(right, r)
	}

	// sort slices
	slices.Sort(left)
	slices.Sort(right)

	// check if lists are of same length
	if len(left) != len(right) {
		log.Fatalln("Number of left and right numbers differ")
	}

	for i := 0; i < len(left); i++ {
		ln := left[i]
		rn := right[i]

		if ln > rn {
			result += ln - rn
		} else {
			result += rn - ln
		}
	}

	fmt.Println(result)
}
