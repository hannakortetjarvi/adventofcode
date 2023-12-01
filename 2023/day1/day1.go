package main

import (
	"2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var nums = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	path := "../inputs/input1.txt"
	lines, err := utils.ReadLines(path)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	PartOne(lines)
	PartTwo(lines)
}

func PartOne(lines []string) {
	var sum int
	for _, line := range lines {
		first := FirstInt(line)
		last := LastInt(line)
		added := fmt.Sprintf("%s%s", first, last)

		num, _ := strconv.Atoi(added)
		sum += num
	}

	fmt.Printf("Result of part one: %v\n", sum)
}

func FirstInt(value string) string {
	for _, character := range value {
		if _, err := strconv.Atoi(string(character)); err == nil {
			return string(character)
		}
	}
	return ""
}

func LastInt(value string) string {
	for i := len(value) - 1; i >= 0; i-- {
		if _, err := strconv.Atoi(string(value[i])); err == nil {
			return string(value[i])
		}
	}
	return ""
}

// --------------------------------------------------------------------------------

func PartTwo(lines []string) {

	var sum int
	for _, line := range lines {
		first := FirstIntOrString(line)
		last := LastIntOrString(line)
		added := fmt.Sprintf("%s%s", first, last)

		num, _ := strconv.Atoi(added)
		sum += num
	}

	fmt.Printf("Result of part two: %v\n", sum)
}

func FirstIntOrString(value string) string {
	var substr string
	for _, character := range value {
		if _, err := strconv.Atoi(string(character)); err == nil {
			return string(character)
		}
		substr = fmt.Sprintf("%s%s", substr, string(character))

		for i, digit := range digits {
			if strings.Contains(substr, digit) {
				return nums[i]
			}
		}
	}
	return ""
}

func LastIntOrString(value string) string {
	var substr string
	for i := len(value) - 1; i >= 0; i-- {
		if _, err := strconv.Atoi(string(value[i])); err == nil {
			return string(value[i])
		}

		substr = fmt.Sprintf("%s%s", string(value[i]), substr)

		for i, digit := range digits {
			if strings.Contains(substr, digit) {
				return nums[i]
			}
		}
	}
	return ""
}
