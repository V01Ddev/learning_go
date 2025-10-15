package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct {
	Fname string
	Lname string
}

func main() {
	var p_slice []person
	var file_name string
	fmt.Print("Please enter the data file: ")
	fmt.Scan(&file_name)

	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// loading data into p_slice
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		split_string := strings.Fields(line)
		fname := split_string[0]
		lname := split_string[1]
		var tmp_person = person{Fname: fname, Lname: lname}
		p_slice = append(p_slice, tmp_person)
	}

	// incrementing through p_slice and printing first and last name
	for _, per := range p_slice {
		fmt.Printf("%s %s\n", per.Fname, per.Lname)
	}
}
