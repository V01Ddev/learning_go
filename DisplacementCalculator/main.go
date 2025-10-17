package main

import (
	"fmt"
	"math"
	"strconv"
)

func GenDisplaceFn(a float64, v0 float64, s0 float64) func (float64) float64 {
	fn := func (t float64) float64 {
		return (0.5 * a * math.Pow(t, 2)) + (v0 * t) + s0
	}
	return fn
}

func main() {
	var a, v0, s0 float64 

	fmt.Print("Acceleration: ")
	fmt.Scan(&a)

	fmt.Print("Initial velocity: ")
	fmt.Scan(&v0)

	fmt.Print("Initial displacement: ")
	fmt.Scan(&s0)

	final_displacement := GenDisplaceFn(a, v0, s0)
	
	fmt.Println("type X to exit")
	var in string
	for {
		fmt.Print("Time: ")
		fmt.Scan(&in)
		if in == "X" {
			break
		}
		var t int
		t, _ = strconv.Atoi(in)
		fmt.Println("Calculated final displacement:", final_displacement(float64(t)))
	}
}
