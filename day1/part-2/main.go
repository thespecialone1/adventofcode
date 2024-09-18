package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("C:/Lahrasab/go-projects/adventofcode/day1/exp.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	
	content := string(data)
	nummberMap := map[string]string{
		"one":"1", "two":"2", "three":"3", "four": "4", "five":"5", "six": "6", "seven": "7", 
		"eight": "8", "nine":"9",
	}
	lines := strings.Split(strings.TrimSpace(content), "\n")
	
	re := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)

	log.Print(re)

	var totalSum int
	for _, line:= range lines{
		numbers := re.FindAllString(line, -1)
			for i, num := range numbers {
				if digit, ok := nummberMap[num]; ok {
					numbers[i] = digit
				}
			}

			// If there are no numbers in the line, skip it
		if len(numbers) == 0 {
			continue;
		}

		var firstNum, lastNum string
		if len(numbers) == 1 {
			firstNum = numbers[0]
			lastNum = numbers[0]
		}else
		{
			firstNum = numbers[0]
			lastNum = numbers[len(numbers)-1]
		}

		combined, _ := strconv.Atoi(firstNum + lastNum)
		totalSum += combined
	}
	
	fmt.Print("The Total Sum is: ", totalSum)
}
