package main

import (
	"2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var redAmount = 12
var greenAmount = 13
var blueAmount = 14

func main() {
	path := "inputs/input2.txt"
	lines, err := utils.ReadLines(path)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	PartOne(lines)
	PartTwo(lines)
}

// Part one of day 2
func PartOne(lines []string) {
	var sum = 0
	for _, line := range lines {
		replaced := strings.ReplaceAll(line, " ", "")
		before, after, _ := strings.Cut(replaced, ":")
		if checkAmounts(after) {
			_, gameNum, _ := strings.Cut(before, "Game")
			num, _ := strconv.Atoi(gameNum)
			sum += num
		}
	}

	fmt.Printf("Result of part one: %v\n", sum)
}

// Check the rolls of each color and find if the amount of cubes is lower than globally defined
func checkAmounts(value string) bool {
	rounds := strings.Split(value, ";")
	for _, round := range rounds {
		colors := strings.Split(round, ",")
		for _, color := range colors {
			// Check red
			beforeRed, _, foundRed := strings.Cut(color, "red")
			if foundRed {
				num, _ := strconv.Atoi(beforeRed)
				if num > redAmount {
					return false
				}
			}
			// Check green
			beforeGreen, _, foundGreen := strings.Cut(color, "green")
			if foundGreen {
				num, _ := strconv.Atoi(beforeGreen)
				if num > greenAmount {
					return false
				}
			}
			// Check blue
			beforeBlue, _, foundBlue := strings.Cut(color, "blue")
			if foundBlue {
				num, _ := strconv.Atoi(beforeBlue)
				if num > blueAmount {
					return false
				}
			}
		}
	}
	return true
}

// --------------------------------------------------------------------------------

// Part two of day 2
func PartTwo(lines []string) {
	sum := 0
	for _, line := range lines {
		replaced := strings.ReplaceAll(line, " ", "")
		_, after, _ := strings.Cut(replaced, ":")
		red, green, blue := checkAmountsOfColor(after)
		sum += (red * green * blue)
	}

	fmt.Printf("Result of part two: %v\n", sum)
}

// Check the rolls of each color and find biggest number for the colors
func checkAmountsOfColor(value string) (int, int, int) {
	blue := 0
	green := 0
	red := 0
	rounds := strings.Split(value, ";")
	for _, round := range rounds {
		colors := strings.Split(round, ",")
		for _, color := range colors {
			// Check red
			beforeRed, _, foundRed := strings.Cut(color, "red")
			if foundRed {
				num, _ := strconv.Atoi(beforeRed)
				if num > red {
					red = num
				}
				continue
			}
			// Check green
			beforeGreen, _, foundGreen := strings.Cut(color, "green")
			if foundGreen {
				num, _ := strconv.Atoi(beforeGreen)
				if num > green {
					green = num
				}
				continue
			}
			// Check blue
			beforeBlue, _, foundBlue := strings.Cut(color, "blue")
			if foundBlue {
				num, _ := strconv.Atoi(beforeBlue)
				if num > blue {
					blue = num
				}
				continue
			}
		}
	}
	return red, green, blue
}
