package main

import (
	"fmt"
	// "log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("exp.txt")
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
	
	var totalSum int
	
	for _, line:= range lines{
		numbers := re.FindAllString(line, -1)
			for i, num := range numbers {
				if digit, ok := nummberMap[num]; ok {
					numbers[i] = digit
				}
			}
			
	var firstNum string
			// If there are no numbers in the line, skip it
		if len(numbers) == 0 {
			continue;
		}

		// If there's only one number, use it for both first and last
		if 	
			len(numbers) == 1 {
				fmt.Println("Extracted:", []string{numbers[0], numbers[0]})
				continue
			}
				
		// Reverse the line to capture last meaningful number from the right side
		
		reversedLine := reverseString(line)
		reversedNumbers := re.FindAllString(reversedLine, -1)

		// Convert words to digits (right-to-left)
		for i, num:= range reversedNumbers{
			if digit, ok := nummberMap[num]; ok{
				reversedNumbers[i] = digit
			}
		}
		
		// Extract the last number from right-to-left
		lastNum := reversedNumbers[0] // After reversing, first element is last num

		if 			// If there's only one number, use it twice
			len(numbers) == 1 {
			firstNum = numbers[0]
			lastNum = numbers[0]
		}

		// Reverse the last number string back to normal order
		lastNum = reverseString(lastNum)

		// Extract the first number from left-to-right
		firstNum = numbers[0]
			// Output the extracted first and last numbers
			fmt.Println("Extracted:", []string{firstNum, lastNum})
			
		combined, _ := strconv.Atoi(firstNum + lastNum)
		totalSum += combined
	}
	fmt.Print("The Total Sum is: ", totalSum)
	}


	

func reverseString(str string) string {
	runes:= []rune(str)
		for i,j := 0, len(runes)-1; i<j; i,j=i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
	return string(runes)
}
