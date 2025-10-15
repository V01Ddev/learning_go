package main

import "fmt"

func BubbleSort(data *[3]int) {
	for i := range len(data) {
		for j := range i {
			fmt.Println(j, i)
			if j < i {
				Swap(data, j)
			}
		}
	}
}

func Swap(data *[3]int, indx int) {
	tmp := data[indx]
	data[indx] = data[indx+1]
	data[indx+1] = tmp
}

func main() {
	var data [3]int
	for i := range data {
		fmt.Printf("Input %d: ", i+1)
		fmt.Scan(&data[i])
	}
	BubbleSort(&data)
	fmt.Println(data)
}
