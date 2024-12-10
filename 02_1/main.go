package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// vars
	result := 0
	// open file
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	// read file and write left and right list
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		var lineint []int

		for _, v := range line {
			vi, _ := strconv.Atoi(v)
			lineint = append(lineint, vi)
		}

		if safe(lineint) {
			result++
		}

	}
	fmt.Println(result)
}

func safe(nums []int) bool {
	state := "inc"
	if nums[0] > nums[1] {
		state = "dec"
	}

	prev := nums[0]
	for _, num := range nums[1:] {
		if num == prev ||
			(state == "inc" && (num < prev || num > prev+3)) ||
			(state == "dec" && (num > prev || num < prev-3)) {
			return false
		}
		prev = num
	}
	return true
}
