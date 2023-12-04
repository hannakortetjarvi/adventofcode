package main

import (
	"2023/utils"
	"fmt"
	"log"
	"slices"
	"strings"
)

var redAmount = 12
var greenAmount = 13
var blueAmount = 14

func main() {
	path := "inputs/input4.txt"
	lines, err := utils.ReadLines(path)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("Result of part one: %v\n", PartOne(lines))
	fmt.Printf("Result of part two: %v\n", PartTwo(lines))
}

// Part one of day 4
func PartOne(lines []string) (sum int) {
	for _, line := range lines {
		_, after, _ := strings.Cut(line, ": ")
		winning, pulled, _ := strings.Cut(after, "|")
		winningList := utils.RemoveWhiteLine(strings.Split(winning, " "))
		pulledList := utils.RemoveWhiteLine(strings.Split(pulled, " "))
		sum += checkCards(winningList, pulledList)
	}

	return
}

func checkCards(winning []string, pulled []string) (sum int) {
	for _, card := range pulled {
		if slices.Contains(winning, card) {
			if sum == 0 {
				sum += 1
			} else {
				sum = sum * 2
			}
		}
	}
	return
}

// --------------------------------------------------------------------------------

// Part two of day 4
func PartTwo(lines []string) (sum int) {
	cards := initList(lines)

	for i := 0; i < len(lines); i++ {
		_, after, _ := strings.Cut(lines[i], ": ")
		winning, pulled, _ := strings.Cut(after, "|")
		winningList := utils.RemoveWhiteLine(strings.Split(winning, " "))
		pulledList := utils.RemoveWhiteLine(strings.Split(pulled, " "))

		found := checkWinningCards(winningList, pulledList, i+1)
		for _, f := range found {
			cards[f] = cards[f] + cards[i+1]
		}
	}

	for _, nums := range cards {
		sum += nums
	}

	return
}

func initList(array []string) map[int]int {
	m := make(map[int]int)
	for i, _ := range array {
		index := i + 1
		m[index] = 1
	}
	return m
}

func checkWinningCards(winning []string, pulled []string, cardNum int) (cards []int) {
	for _, card := range pulled {
		if slices.Contains(winning, card) {
			newCard := cardNum + len(cards) + 1
			cards = append(cards, newCard)
		}
	}
	return
}
