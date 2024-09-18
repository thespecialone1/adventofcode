package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

//	func check(e error) {
//		if e != nil {
//			panic(e)
//		}
//	}
func main() {
	dat, err := os.ReadFile("doc.txt")
	if err != nil {
		fmt.Print(err)
		return
	}

	content := string(dat)
	lines := strings.Split(strings.TrimSpace(content), "\n")

	re := regexp.MustCompile("\\d+")

	for  _, line:= range lines{
		numbers:= re.FindAllString(line, -1)
		 if len(numbers) > 0 {
			return
		 }
	}
	slices = 
}
