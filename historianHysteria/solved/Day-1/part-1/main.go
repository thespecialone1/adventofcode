package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var leftList, rightList []int

func main() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		log.Fatal("Error while opening the file", err)
	}
	scanner := bufio.NewScanner(file)
	// fmt.Println(scanner.Text())
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		leftStr := parts[0]
		rightStr := parts[1]

		left, err1 := strconv.Atoi(leftStr)
		right, err2 := strconv.Atoi(rightStr)
		if err1 != nil || err2 != nil {
			log.Fatal("Invalid number in line", line)
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while Scanning: %v", err)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0

	for i, leftVal := range leftList {
		rightVal := rightList[i]
		diff := abs(leftVal - rightVal)
		fmt.Printf("Pair %d: %d vs %d -> Distace: %d\n", i+1, leftVal, rightVal, diff)
		totalDistance += diff
	}
	fmt.Println("Total Distance", totalDistance)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
