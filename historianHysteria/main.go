package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open("puzzleInput.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(data)
	safeCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		var levels []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			levels = append(levels, num)
		}
		if isSafe(levels) {
			safeCount++
		}
	}
	fmt.Println(safeCount)
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	diff := levels[1] - levels[0]
	if diff == 0 || abs(diff) > 3 {
		return false
	}
	increasing := diff > 0
	for i := 1; i < len(levels); i++ {
		d := levels[i] - levels[i-1]
		if d == 0 || abs(d) > 3 {
			return false
		}

		if increasing && d <= 0 {
			return false
		}
		if !increasing && d >= 0 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
