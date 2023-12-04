package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func RemoveWhiteLine(list []string) (array []string) {
	for _, num := range list {
		newVal := strings.ReplaceAll(num, " ", "")
		if newVal != "" {
			array = append(array, strings.ReplaceAll(num, " ", ""))
		}
	}
	return
}
