package main

import "fmt"
import "slices"
import "strconv"

func main() {


	var input string
	var sli [] int
	sli = make([]int, 3)

	i := 0
	for {
		fmt.Scan(&input)
		if input != "X" {
			if i < 3 {
				sli[0], _ = strconv.Atoi(input)
				slices.Sort(sli)
				fmt.Println(sli)
				i++
			} else {
				x, _ := strconv.Atoi(input)
				sli = append(sli, x)
				slices.Sort(sli)
				fmt.Println(sli)
			}
		} else {
			break
		}
	}
}
