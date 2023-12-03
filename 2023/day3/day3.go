package main

import (
	"2023/utils"
	"fmt"
	"log"
	"strconv"
)

func main() {
	path := "inputs/input3.txt"
	lines, err := utils.ReadLines(path)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("Result of part one: %v\n", PartOne(lines))
	fmt.Printf("Result of part two: %v\n", PartTwo(lines))
}

// Part one of day 3
func PartOne(lines []string) (sum int) {
	var vals []int
	for i := 0; i < len(lines); i++ {
		if i == 0 {
			vals = CheckTwoLines(lines[i], lines[i+1])
		} else if i == (len(lines) - 1) {
			vals = CheckTwoLines(lines[i], lines[i-1])
		} else {
			vals = CheckThreeLines(lines[i-1], lines[i], lines[i+1])
		}

		for _, num := range vals {
			sum += num
		}
	}

	return
}

// Check if number is adjacant to symbol with two lines
func CheckTwoLines(s string, checked string) (adjacent []int) {
	foundSymbol := false
	num := ""
	for i := 0; i < len(s); i++ {
		if !IsDigit(s[i]) {
			if num != "" && foundSymbol {
				foundNum, _ := strconv.Atoi(num)
				adjacent = append(adjacent, foundNum)
			}
			num = ""
			foundSymbol = false
			continue
		}

		num = fmt.Sprintf("%s%c", num, s[i])

		if i == 0 {
			if !IsDigitOrDot(checked[i]) || !IsDigitOrDot(checked[i+1]) || !IsDigitOrDot(s[i+1]) {
				foundSymbol = true
			}
		} else if i == (len(s) - 1) {
			if !IsDigitOrDot(checked[i]) || !IsDigitOrDot(checked[i-1]) || !IsDigitOrDot(s[i-1]) {
				foundSymbol = true
			}
		} else {
			if !IsDigitOrDot(checked[i]) || !IsDigitOrDot(checked[i-1]) || !IsDigitOrDot(checked[i+1]) || !IsDigitOrDot(s[i+1]) || !IsDigitOrDot(s[i-1]) {
				foundSymbol = true
			}
		}
	}

	if num != "" && foundSymbol {
		foundNum, _ := strconv.Atoi(num)
		adjacent = append(adjacent, foundNum)
	}

	return
}

// Check if number is adjacant to symbol with three lines
func CheckThreeLines(before string, s string, after string) (adjacent []int) {
	foundSymbol := false
	num := ""
	for i := 0; i < len(s); i++ {
		if !IsDigit(s[i]) {
			if num != "" && foundSymbol {
				foundNum, _ := strconv.Atoi(num)
				adjacent = append(adjacent, foundNum)
			}
			num = ""
			foundSymbol = false
			continue
		}

		num = fmt.Sprintf("%s%c", num, s[i])

		if i == 0 {
			if !IsDigitOrDot(after[i]) || !IsDigitOrDot(after[i+1]) ||
				!IsDigitOrDot(before[i]) || !IsDigitOrDot(before[i+1]) ||
				!IsDigitOrDot(s[i+1]) {
				foundSymbol = true
			}
		} else if i == (len(s) - 1) {
			if !IsDigitOrDot(after[i]) || !IsDigitOrDot(after[i-1]) ||
				!IsDigitOrDot(before[i]) || !IsDigitOrDot(before[i-1]) ||
				!IsDigitOrDot(s[i-1]) {
				foundSymbol = true
			}
		} else {
			if !IsDigitOrDot(after[i]) || !IsDigitOrDot(after[i-1]) || !IsDigitOrDot(after[i+1]) ||
				!IsDigitOrDot(before[i]) || !IsDigitOrDot(before[i-1]) || !IsDigitOrDot(before[i+1]) ||
				!IsDigitOrDot(s[i+1]) || !IsDigitOrDot(s[i-1]) {
				foundSymbol = true
			}
		}
	}

	if num != "" && foundSymbol {
		foundNum, _ := strconv.Atoi(num)
		adjacent = append(adjacent, foundNum)
	}

	return
}

// Check if byte is number
func IsDigit(value byte) bool {
	if _, err := strconv.Atoi(string(value)); err == nil {
		return true
	}
	return false
}

// Check if byte is digit or a dot
func IsDigitOrDot(value byte) bool {
	return (IsDigit(value) || value == '.')
}

// Part two of day 3
func PartTwo(lines []string) (sum int) {
	var vals []int
	for i := 0; i < len(lines); i++ {
		if i == 0 {
			vals = CheckTwoLinesGear(lines[i], lines[i+1])
		} else if i == (len(lines) - 1) {
			vals = CheckTwoLinesGear(lines[i], lines[i-1])
		} else {
			vals = CheckThreeLinesGear(lines[i-1], lines[i], lines[i+1])
		}

		for _, num := range vals {
			sum += num
		}
	}

	return
}

// Check if gear has two numbers near it with two lines
func CheckTwoLinesGear(s string, checked string) (gear []int) {
	num := ""
	nextNum := ""
	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			if IsDigit(s[i]) {
				num = fmt.Sprintf("%s%c", num, s[i])
			} else {
				num = ""
			}
			continue
		}
		nextNum = CheckNextNum(s[i+1:])
		checkedLineNums := FindAdjacantNums(checked, i)

		if num != "" {
			checkedLineNums = append(checkedLineNums, num)
		}
		if nextNum != "" {
			checkedLineNums = append(checkedLineNums, nextNum)
		}

		fmt.Printf("Found nums: %v\n", checkedLineNums)

		if len(checkedLineNums) == 2 {
			firstNum, _ := strconv.Atoi(checkedLineNums[0])
			secondNum, _ := strconv.Atoi(checkedLineNums[1])
			gear = append(gear, firstNum*secondNum)
		}
	}

	return
}

// Check if the string starts with number and return that number
func CheckNextNum(s string) (found string) {
	for i := 0; i < len(s); i++ {
		if IsDigit(s[i]) {
			found = fmt.Sprintf("%s%c", found, s[i])
		} else {
			break
		}
	}
	return
}

// Check if there are numbers near gearIndex
func FindAdjacantNums(s string, gearIndex int) (found []string) {
	num := ""
	startIndex := 0

	for i := 0; i < len(s); i++ {
		if IsDigit(s[i]) {
			if num == "" {
				startIndex = i
			}
			num = fmt.Sprintf("%s%c", num, s[i])
		} else {
			if num != "" {
				if startIndex == gearIndex+1 {
					found = append(found, num)
				} else if i-1 == gearIndex-1 {
					found = append(found, num)
				} else if startIndex <= gearIndex && gearIndex <= i-1 {
					found = append(found, num)
				}
			}
			num = ""
		}
	}

	return
}

// Check if gear has two numbers near it with three lines
func CheckThreeLinesGear(before string, s string, after string) (gear []int) {
	num := ""
	nextNum := ""
	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			if IsDigit(s[i]) {
				num = fmt.Sprintf("%s%c", num, s[i])
			} else {
				num = ""
			}
			continue
		}

		nextNum = CheckNextNum(s[i+1:])
		numsBefore := FindAdjacantNums(before, i)
		numsAfter := FindAdjacantNums(after, i)
		foundNums := append(numsBefore, numsAfter...)

		if num != "" {
			foundNums = append(foundNums, num)
		}
		if nextNum != "" {
			foundNums = append(foundNums, nextNum)
		}

		fmt.Printf("Found nums: %v\n", foundNums)

		if len(foundNums) == 2 {
			firstNum, _ := strconv.Atoi(foundNums[0])
			secondNum, _ := strconv.Atoi(foundNums[1])
			gear = append(gear, firstNum*secondNum)
		}
	}

	return
}
