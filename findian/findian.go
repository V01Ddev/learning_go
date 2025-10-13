package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)
	input = strings.ToLower(input)
	
	if strings.HasPrefix(input, "i") && strings.Contains(input, "a") && strings.HasSuffix(input, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
