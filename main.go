package main

import (
	"advent/challenges/day1"
	"fmt"
)

func main() {
	result, err := day1.Solution("./challenges/day1/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
