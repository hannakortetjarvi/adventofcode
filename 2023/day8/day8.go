package main

import (
	"2023/utils"
	"fmt"
	"log"
	"slices"
	"strings"
)

type Node struct {
	left  string
	right string
}

func main() {
	path := "inputs/input8.txt"
	lines, err := utils.ReadLines(path)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	lines = utils.RemoveEmptyLine(lines)

	//fmt.Printf("Result of part one: %v\n", PartOne(lines))
	fmt.Printf("Result of part two: %v\n", PartTwo(lines))
}

// Part one of day 8
func PartOne(lines []string) (sum int) {
	rules := lines[0]
	nodes := make(map[string]Node)
	current := "AAA"
	final := "ZZZ"
	for i := 1; i < len(lines); i++ {
		line := filterLine(lines[i])
		name, directions, _ := strings.Cut(line, "=")
		left, right, _ := strings.Cut(directions, ",")
		nodes[name] = Node{left, right}
	}

	index := 0
	for current != final {
		sum++
		if index == len(rules) {
			index = 0
		}

		if string(rules[index]) == "L" {
			current = nodes[current].left
		} else {
			current = nodes[current].right
		}
		index++
	}

	return
}

func filterLine(line string) string {
	filtered := strings.ReplaceAll(line, " ", "")
	filtered = strings.ReplaceAll(filtered, "(", "")
	filtered = strings.ReplaceAll(filtered, ")", "")
	return filtered
}

// Part two of day 8
func PartTwo(lines []string) (sum int) {
	rules := lines[0]
	nodes := make(map[string]Node)
	var current []string
	for i := 1; i < len(lines); i++ {
		line := filterLine(lines[i])
		name, directions, _ := strings.Cut(line, "=")
		left, right, _ := strings.Cut(directions, ",")
		nodes[name] = Node{left, right}
		if string(name[2]) == "A" {
			current = append(current, name)
		}
	}

	index := 0
	num := 0
	steps := make([]int, len(current))
	numOfSteps := 0
	for numOfSteps != len(current) {
		num++
		if index == len(rules) {
			index = 0
		}
		for i := 0; i < len(current); i++ {
			new := ""
			if string(rules[index]) == "L" {
				new = nodes[current[i]].left
			} else {
				new = nodes[current[i]].right
			}
			current[i] = new
			if string(current[i][2]) == "Z" {
				steps[i] = num
				numOfSteps++
			}
		}
		index++
	}

	ok := false
	max := slices.Max(steps)
	i := 0
	for !ok {
		i++
		ok = true
		for _, step := range steps {
			if (i*max)%step != 0 {
				ok = false
			}
		}
	}
	sum = i * max
	return
}
