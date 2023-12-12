package main

import (
	"2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	path := "inputs/input9.txt"
	lines, err := utils.ReadLines(path)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("Result of part one: %v\n", PartOne(lines))
	fmt.Printf("Result of part two: %v\n", PartTwo(lines))
}

// Part one of day 9
func PartOne(lines []string) (sum int) {
	var arr [][]int
	for i := 0; i < len(lines); i++ {
		var numArr []int
		vals := strings.Split(lines[i], " ")
		for _, val := range vals {
			num, _ := strconv.Atoi(val)
			numArr = append(numArr, num)
		}
		arr = append(arr, numArr)
	}

	for _, nums := range arr {
		var pyramid [][]int
		i := 0
		pyramid = append(pyramid, nums)
		for !allZero(pyramid[len(pyramid)-1]) {
			var newArr []int
			for j := 1; j < len(pyramid[i]); j++ {
				next := pyramid[i][j] - pyramid[i][j-1]
				newArr = append(newArr, next)
			}
			pyramid = append(pyramid, newArr)
			i++
		}

		for j := len(pyramid) - 2; j >= 0; j-- {
			final := pyramid[j+1][len(pyramid[j+1])-1]
			pyramid[j] = append(pyramid[j], final+pyramid[j][len(pyramid[j])-1])
		}
		sum += pyramid[0][len(pyramid[0])-1]
	}

	return
}

func allZero(arr []int) bool {
	for _, num := range arr {
		if num != 0 {
			return false
		}
	}
	return true
}

// Part two of day 9
func PartTwo(lines []string) (sum int) {
	var arr [][]int
	for i := 0; i < len(lines); i++ {
		var numArr []int
		vals := strings.Split(lines[i], " ")
		for _, val := range vals {
			num, _ := strconv.Atoi(val)
			numArr = append(numArr, num)
		}
		arr = append(arr, numArr)
	}

	for _, nums := range arr {
		var pyramid [][]int
		i := 0
		pyramid = append(pyramid, nums)
		for !allZero(pyramid[len(pyramid)-1]) {
			var newArr []int
			for j := 1; j < len(pyramid[i]); j++ {
				next := pyramid[i][j] - pyramid[i][j-1]
				newArr = append(newArr, next)
			}
			pyramid = append(pyramid, newArr)
			i++
		}

		ex := 0
		for j := len(pyramid) - 2; j >= 0; j-- {
			first := pyramid[j][0]
			ex = first - ex
		}
		sum += ex
	}

	return
}
