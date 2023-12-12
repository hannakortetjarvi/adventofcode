package main

import (
	"2023/utils"
	"fmt"
	"log"
	"strings"
)

func main() {
	path := "inputs/input10.txt"
	lines, err := utils.ReadLines(path)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("Result of part one: %v\n", PartOne(lines))
	//fmt.Printf("Result of part two: %v\n", PartTwo(lines))
}

// Part one of day 10
func PartOne(lines []string) (sum int) {
	var loop [][]string
	x := 0
	y := 0
	for i := 0; i < len(lines); i++ {
		var line []string
		for _, c := range lines[i] {
			line = append(line, string(c))
		}
		loop = append(loop, line)
		if strings.Contains(lines[i], "S") {
			x = strings.Index(lines[i], "S")
			y = i
		}
	}

	steps := -1
	prev_x := x
	prev_y := y
	for {
		steps++
		if string(loop[y][x]) == "0" {
			break
		}

		if string(loop[y][x]) == "S" {
			loop[y][x] = fmt.Sprint(steps)
			if y == 0 && x == 0 {
				x++
			} else if y == len(loop[y])-1 && x == len(loop[y][x])-1 {
				x--
			} else if y == 0 && x == len(loop[y][x])-1 {
				y++
			} else if y == len(loop[y])-1 && x == 0 {
				x++
			} else if y == 0 {
				if loop[y][x+1] == "-" || loop[y][x+1] == "7" {
					x++
				} else {
					x--
				}
			} else if y == len(loop[y])-1 {
				if loop[y][x+1] == "-" || loop[y][x+1] == "J" {
					x++
				} else {
					x--
				}
			} else if x == 0 {
				if loop[y+1][x] == "|" || loop[y+1][x] == "L" {
					y++
				} else {
					y--
				}
			} else if x == len(loop[y][x])-1 {
				if loop[y+1][x] == "|" || loop[y+1][x] == "F" {
					y--
				} else {
					y++
				}
			} else {
				if loop[y][x+1] == "-" || loop[y][x+1] == "J" || loop[y][x+1] == "7" {
					x++
				} else if loop[y][x-1] == "-" || loop[y][x-1] == "F" || loop[y][x+1] == "L" {
					x--
				} else {
					y++
				}
			}
			continue
		}

		if string(loop[y][x]) == "J" {
			loop[y][x] = fmt.Sprint(steps)
			if prev_x == x-1 {
				prev_x = x
				y--
			} else {
				prev_y = y
				x--
			}
			continue
		}

		if string(loop[y][x]) == "F" {
			loop[y][x] = fmt.Sprint(steps)
			if prev_x == x+1 {
				prev_x = x
				y++
			} else {
				prev_y = y
				x++
			}
			continue
		}

		if string(loop[y][x]) == "7" {
			loop[y][x] = fmt.Sprint(steps)
			if prev_x == x-1 {
				prev_x = x
				y++
			} else {
				prev_y = y
				x--
			}
			continue
		}

		if string(loop[y][x]) == "L" {
			loop[y][x] = fmt.Sprint(steps)
			if prev_x == x+1 {
				prev_x = x
				y--
			} else {
				prev_y = y
				x++
			}
			continue
		}

		if string(loop[y][x]) == "-" {
			loop[y][x] = fmt.Sprint(steps)
			if prev_x == x-1 {
				prev_x = x
				x++
			} else {
				prev_x = x
				x--
			}
			continue
		}

		if string(loop[y][x]) == "|" {
			loop[y][x] = fmt.Sprint(steps)
			if prev_y == y-1 {
				prev_y = y
				y++
			} else {
				prev_y = y
				y--
			}
			continue
		}
	}

	sum = steps / 2

	return
}

func PrintLoop(loop [][]string) {
	for _, s := range loop {
		fmt.Println(s)
	}
}
