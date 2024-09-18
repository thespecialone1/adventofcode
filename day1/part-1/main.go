package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//	func check(e error) {
//		if e != nil {
//			panic(e)
//		}
//	}

func main()  {
	data, err := os.ReadFile("doc.txt")
	if err != nil {
		fmt.Print(err)
	
}
		// Convert content to string and trim any extra spaces or newlines
	content := string(data)	
		// Split the content into lines
	lines := strings.Split(strings.TrimSpace(content), "\n")
		
		// Compile a regular expression to find sequences of digits
	 re := regexp.MustCompile("\\d")

	var totalSum int
		// Loop through each line
	for _, line := range lines {
		// Find all numbers in the line
		numbers := re.FindAllString(line, -1)
		// if len(numbers)>0 {
		// 	fmt.Println("Numbers in Line are: ", numbers)
		// }

		// If there are no numbers in the line, skip it
		if len(numbers) == 0 {
			continue
		}
		var firstNum, lastNum string

		if 			// If there's only one number, use it twice
			len(numbers) == 1 {
			firstNum = numbers[0]
			lastNum = numbers[0]
		}else		// Get the first number from the left and the first from the right
		{		
			firstNum = numbers[0]
			lastNum = numbers[len(numbers)- 1]
		}
			combined, _ := strconv.Atoi(firstNum + lastNum)

			totalSum += combined
	}
	fmt.Print("The total sum is: ", totalSum)
}


