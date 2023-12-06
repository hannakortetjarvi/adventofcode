package main

import (
	"2023/utils"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	path := "inputs/input5.txt"
	lines, err := utils.ReadLines(path)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	lines = utils.RemoveEmptyLine(lines)

	fmt.Printf("Result of part one: %v\n", PartOne(lines))
	fmt.Printf("Result of part two: %v\n", PartTwo(lines))
}

// Part one of day 5
func PartOne(lines []string) int {
	seeds := InitSeeds(lines[0])
	var found []string
	for _, line := range lines {
		if !utils.IsDigit(line[0]) {
			found = nil
			continue
		}
		nums := getNums(line)
		for seed, val := range seeds {
			if !slices.Contains(found, seed) {
				prev := seeds[seed]
				newVal := GetDestination(nums[0], nums[1], nums[2], val)
				if prev != newVal {
					seeds[seed] = newVal
					found = append(found, seed)
				}
			}
		}

	}

	min := math.MaxInt32
	for _, s := range seeds {
		if s < min {
			min = s
		}
	}

	return min
}

func InitSeeds(s string) map[string]int {
	seedsMap := make(map[string]int)
	_, after, _ := strings.Cut(s, ": ")
	seeds := strings.Split(after, " ")
	for _, seed := range seeds {
		num, _ := strconv.Atoi(seed)
		seedsMap[seed] = num
	}
	return seedsMap
}

func getNums(s string) (nums []int) {
	seeds := strings.Split(s, " ")
	for _, seed := range seeds {
		num, _ := strconv.Atoi(seed)
		nums = append(nums, num)
	}
	return
}

func GetDestination(dest int, sour int, rang int, num int) int {
	if sour <= num && num <= sour+rang-1 {
		diff := num - sour
		return dest + diff
	}
	return num
}

// --------------------------------------------------------------------------------

// Part two of day 5
func PartTwo(lines []string) int {
	_, after, _ := strings.Cut(lines[0], ": ")
	seedPairs := strings.Split(after, " ")
	min := LargestSeed(seedPairs)
	fmt.Println(min)

	for i := 0; i < len(seedPairs)/2; i++ {
		seeds := InitRangeSeeds(seedPairs[i*2], seedPairs[i*2+1])
		for j := 0; j < len(seeds); j++ {
			checked := false
			for _, line := range lines {
				if checked && utils.IsDigit(line[0]) {
					continue
				}
				if !utils.IsDigit(line[0]) {
					checked = false
					continue
				}
				nums := getNums(line)
				prev := seeds[j]
				newVal := GetDestination(nums[0], nums[1], nums[2], seeds[j])
				if prev != newVal {
					seeds[j] = newVal
					checked = true
				}
			}
			if seeds[j] < min {
				min = seeds[j]
				fmt.Println(min)
			}
		}
	}

	return min
}

func LargestSeed(seeds []string) (max int) {
	for _, seed := range seeds {
		num, _ := strconv.Atoi(seed)
		if num > max {
			max = num
		}
	}
	return
}

func InitRangeSeeds(start string, end string) []int {
	s, _ := strconv.Atoi(start)
	e, _ := strconv.Atoi(end)

	seedsSlice := make([]int, e)
	for j := 0; j < e; j++ {
		seedsSlice[j] = s + j
	}
	return seedsSlice
}
