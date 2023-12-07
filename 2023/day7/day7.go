package main

import (
	"2023/utils"
	"fmt"
	"log"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	rank  int
	bid   int
	cards string
}

func main() {
	path := "inputs/input7.txt"
	lines, err := utils.ReadLines(path)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("Result of part one: %v\n", PartOne(lines))
	fmt.Printf("Result of part two: %v\n", PartTwo(lines))
}

// Part one of day 7
func PartOne(lines []string) (sum int) {
	handsList := make([]Hand, len(lines))
	for i := 0; i < len(lines); i++ {
		vals := strings.Split(lines[i], " ")
		rank, _ := strconv.Atoi(vals[1])
		hand := Hand{i + 1, rank, vals[0]}
		handsList[i] = hand
	}

	for j := 0; j < len(handsList); j++ {
		for i := j + 1; i < len(handsList); i++ {
			currentCards := handsList[j].cards
			if handsList[j] == handsList[i] {
				continue
			}
			checkedCards := handsList[i].cards
			current := checkHand(currentCards)
			checked := checkHand(checkedCards)
			currentRank := handsList[j].rank
			checkedRank := handsList[i].rank
			if checked < current && currentRank < checkedRank && j < i {
				handsList[i].rank = currentRank
				handsList[j].rank = checkedRank
			} else if checked == current && currentRank < checkedRank && j < i {
				bigger := checkBiggerHand(currentCards, checkedCards)
				if bigger == currentCards {
					handsList[i].rank = currentRank
					handsList[j].rank = checkedRank
				}
			}
			sort.Slice(handsList, func(i, j int) bool {
				return handsList[i].rank < handsList[j].rank
			})
		}
	}

	for _, hand := range handsList {
		sum += (hand.rank * hand.bid)
	}

	return
}

// Check the value of cards
func checkHand(cards string) (val int) {
	count := make(map[rune]int)
	for _, c := range cards {
		_, ok := count[c]
		if ok {
			count[c] += 1
		} else {
			count[c] = 1
		}
	}
	var counts []int
	for _, val := range count {
		counts = append(counts, val)
	}
	slices.Sort(counts)
	if reflect.DeepEqual(counts, []int{1, 1, 1, 1, 1}) {
		val = 1
	} else if reflect.DeepEqual(counts, []int{1, 1, 1, 2}) {
		val = 2
	} else if reflect.DeepEqual(counts, []int{1, 2, 2}) {
		val = 3
	} else if reflect.DeepEqual(counts, []int{1, 1, 3}) {
		val = 4
	} else if reflect.DeepEqual(counts, []int{2, 3}) {
		val = 5
	} else if reflect.DeepEqual(counts, []int{1, 4}) {
		val = 6
	} else {
		val = 7
	}

	return
}

// Check which hand is bigger
func checkBiggerHand(current string, checked string) string {
	bigger := ""
	for i := 0; i < len(current); i++ {
		c1 := string(current[i])
		c2 := string(checked[i])

		if utils.IsDigit(c1[0]) && !utils.IsDigit(c2[0]) {
			return checked
		} else if !utils.IsDigit(c1[0]) && utils.IsDigit(c2[0]) {
			return current
		} else if utils.IsDigit(c1[0]) && utils.IsDigit(c2[0]) {
			num1, _ := strconv.Atoi(string(c1))
			num2, _ := strconv.Atoi(string(c2))
			if num1 > num2 {
				return current
			} else if num2 > num1 {
				return checked
			} else {
				continue
			}
		} else {
			if c1 == c2 {
				continue
			} else if c1 == "T" {
				return checked
			} else if c1 == "J" && c2 != "T" {
				return checked
			} else if c1 == "Q" && (c2 != "T" && c2 != "J") {
				return checked
			} else if c1 == "K" && c2 == "A" {
				return checked
			} else {
				return current
			}
		}

	}

	return bigger
}

// --------------------------------------------------------------------------------

// Part two of day 7
func PartTwo(lines []string) (sum int) {
	handsList := make([]Hand, len(lines))
	for i := 0; i < len(lines); i++ {
		vals := strings.Split(lines[i], " ")
		rank, _ := strconv.Atoi(vals[1])
		hand := Hand{i + 1, rank, vals[0]}
		handsList[i] = hand
	}

	for j := 0; j < len(handsList); j++ {
		for i := j + 1; i < len(handsList); i++ {
			currentCards := handsList[j].cards
			if handsList[j] == handsList[i] {
				continue
			}
			checkedCards := handsList[i].cards
			current := checkHandWithJoker(currentCards)
			checked := checkHandWithJoker(checkedCards)
			currentRank := handsList[j].rank
			checkedRank := handsList[i].rank
			if checked < current && currentRank < checkedRank && j < i {
				handsList[i].rank = currentRank
				handsList[j].rank = checkedRank
			} else if checked == current && currentRank < checkedRank && j < i {
				bigger := checkBiggerHandWithJoker(currentCards, checkedCards)
				if bigger == currentCards {
					handsList[i].rank = currentRank
					handsList[j].rank = checkedRank
				}
			}
			sort.Slice(handsList, func(i, j int) bool {
				return handsList[i].rank < handsList[j].rank
			})
		}
	}

	for _, hand := range handsList {
		sum += (hand.rank * hand.bid)
	}

	return
}

// Check which hand is bigger, now joker J is the lowest
func checkBiggerHandWithJoker(current string, checked string) string {
	bigger := ""
	for i := 0; i < len(current); i++ {
		c1 := string(current[i])
		c2 := string(checked[i])

		if utils.IsDigit(c1[0]) && !utils.IsDigit(c2[0]) {
			if c2 == "J" {
				return current
			} else {
				return checked
			}
		} else if !utils.IsDigit(c1[0]) && utils.IsDigit(c2[0]) {
			if c1 == "J" {
				return checked
			} else {
				return current
			}
		} else if utils.IsDigit(c1[0]) && utils.IsDigit(c2[0]) {
			num1, _ := strconv.Atoi(string(c1))
			num2, _ := strconv.Atoi(string(c2))
			if num1 > num2 {
				return current
			} else if num2 > num1 {
				return checked
			} else {
				continue
			}
		} else {
			if c1 == c2 {
				continue
			} else if c1 == "J" {
				return checked
			} else if c1 == "T" && c2 != "J" {
				return checked
			} else if c1 == "Q" && (c2 != "T" && c2 != "J") {
				return checked
			} else if c1 == "K" && c2 == "A" {
				return checked
			} else {
				return current
			}
		}

	}

	return bigger
}

// Check the value of cards, now joker J is the lowest
func checkHandWithJoker(cards string) (val int) {
	count := make(map[string]int)
	jokers := 0
	for i := 0; i < len(cards); i++ {
		c := string(cards[i])
		if c == "J" {
			jokers++
			continue
		}
		_, ok := count[c]
		if ok {
			count[c] += 1
		} else {
			count[c] = 1
		}
	}

	var counts []int
	for _, val := range count {
		counts = append(counts, val)
	}
	slices.Sort(counts)
	if jokers == 0 {
		if reflect.DeepEqual(counts, []int{1, 1, 1, 1, 1}) {
			val = 1
		} else if reflect.DeepEqual(counts, []int{1, 1, 1, 2}) {
			val = 2
		} else if reflect.DeepEqual(counts, []int{1, 2, 2}) {
			val = 3
		} else if reflect.DeepEqual(counts, []int{1, 1, 3}) {
			val = 4
		} else if reflect.DeepEqual(counts, []int{2, 3}) {
			val = 5
		} else if reflect.DeepEqual(counts, []int{1, 4}) {
			val = 6
		} else if reflect.DeepEqual(counts, []int{5}) {
			val = 7
		}
	} else if jokers == 5 {
		val = 7
	} else if jokers == 4 {
		val = 7
	} else if jokers == 3 {
		max := slices.Max(counts)
		if max == 1 {
			val = 6
		} else {
			val = 7
		}
	} else if jokers == 2 {
		max := slices.Max(counts)
		if max == 1 {
			val = 4
		} else if max == 2 {
			val = 6
		} else {
			val = 7
		}
	} else {
		max := slices.Max(counts)
		if max == 1 {
			val = 2
		} else if max == 2 && reflect.DeepEqual(counts, []int{1, 1, 2}) {
			val = 4
		} else if max == 2 && reflect.DeepEqual(counts, []int{2, 2}) {
			val = 5
		} else if max == 3 {
			val = 6
		} else {
			val = 7
		}
	}
	return
}
