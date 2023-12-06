package main

import (
	"2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	path := "inputs/input6.txt"
	lines, err := utils.ReadLines(path)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("Result of part one: %v\n", PartOne(lines))
	fmt.Printf("Result of part two: %v\n", PartTwo(lines))
}

// Part one of day 6
func PartOne(lines []string) (sum int) {
	_, afterTime, _ := strings.Cut(lines[0], ":")
	times := utils.RemoveWhiteLine(strings.Split(afterTime, " "))
	_, afterDistance, _ := strings.Cut(lines[1], ":")
	distances := utils.RemoveWhiteLine(strings.Split(afterDistance, " "))

	var waysToWin []int
	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		best, _ := strconv.Atoi(distances[i])
		var winning []int
		for j := 0; j <= time; j++ {
			newDistance := j * (time - j)
			if newDistance > best {
				winning = append(winning, newDistance)
			}
		}
		waysToWin = append(waysToWin, len(winning))
	}

	for _, win := range waysToWin {
		if sum == 0 {
			sum += win
		} else {
			sum = sum * win
		}
	}

	return
}

// --------------------------------------------------------------------------------

// Part two of day 6
func PartTwo(lines []string) (sum int) {
	_, afterTime, _ := strings.Cut(lines[0], ":")
	times := strings.ReplaceAll(afterTime, " ", "")
	_, afterDistance, _ := strings.Cut(lines[1], ":")
	distances := strings.ReplaceAll(afterDistance, " ", "")

	var waysToWin []int
	time, _ := strconv.Atoi(times)
	best, _ := strconv.Atoi(distances)
	for j := 0; j <= time; j++ {
		newDistance := j * (time - j)
		if newDistance > best {
			waysToWin = append(waysToWin, newDistance)
		}
	}
	sum = len(waysToWin)

	return
}
