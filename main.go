package main

import (
	"advent/challenges/day2"
	"fmt"
)

func main() {
	result, err := day2.Solution("./challenges/day2/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
