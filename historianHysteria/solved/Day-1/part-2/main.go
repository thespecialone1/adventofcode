package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var rightList, leftlist []int

func main() {
	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		leftStr := line[0]
		rightStr := line[1]

		convertLeftIntoInt, err1 := strconv.Atoi(leftStr)
		convertRightIntoIntoInt, err2 := strconv.Atoi(rightStr)
		if err1 != nil || err2 != nil {
			log.Fatal(err1, err2)
		}
		leftlist = append(leftlist, convertLeftIntoInt)
		rightList = append(rightList, convertRightIntoIntoInt)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	rightCount := make(map[int]int)
	for _, num := range rightList {
		rightCount[num]++
	}

	similarityScore := 0
	for _, leftVal := range leftlist {
		countInRight := rightCount[leftVal]
		score := leftVal * countInRight
		fmt.Printf("Left value %d apears %d times in right -> +%d points \n", leftVal, countInRight, score)
		similarityScore += score
	}
	fmt.Printf("\nTotal Similarity Score: %d\n", similarityScore)
}
