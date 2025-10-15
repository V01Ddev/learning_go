package main

import (
	"bufio"
	"strings"
	"os"
	"encoding/json"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	m := make(map[string]string)

	fmt.Print("Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name", err)
		return
	}
	name = strings.TrimSpace(name)

	fmt.Print("Address: ")
	address, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading address", err)
		return
	}
	address = strings.TrimSpace(address)

	m["name"] = name
	m["address"] = address

	resB, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error while marshaling:", err)
		return
	}

	fmt.Println(string(resB))
}
