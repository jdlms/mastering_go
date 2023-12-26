package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var Global int = 1234
var AnotherGlobal int = -5678

func Basics() {

	// vars

	var j int
	i := Global + AnotherGlobal
	fmt.Println("Initial j value", j)
	j = Global

	// math.Abs() requires a float64 param
	k := math.Abs(float64(AnotherGlobal))
	fmt.Printf("Global=%d, i=%d, j=%d, k=%.2f.\n", Global, i, j, k)

	// switch

	if len(os.Args) != 2 {
		fmt.Println("Please give a command line argument...")
		return
	}
	argument := os.Args[1]

	// expression after switch
	switch argument {
	case "0":
		fmt.Println("Zero!")
	case "1":
		fmt.Println("One!")
	case "2", "3", "4":
		fmt.Println("2 or 3 or 4")
		fallthrough
	default:
		fmt.Println("Value:", argument)
	}

	value, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Cannot convert to int:", argument)
		return
	}

	// no expression after switch

	switch {
	case value == 0:
		fmt.Println("Zero!")
	case value > 0:
		fmt.Println("Positive int!")
	case value < 0:
		fmt.Println("Negative int!")
	default:
		fmt.Println("This should not happen:", value)

	}

	// loops

	// traditional for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()

	// idiomatic and verbose for loop in Go
	idx := 0
	for ok := true; ok; ok = (idx != 10) {
		fmt.Print(idx*idx, "")
		idx++
	}
	fmt.Println()

	// for loop used as a while loop
	l := 0
	for {
		if l == 10 {
			break
		}
		fmt.Print(l*l, " ")
		l++
	}
	fmt.Println()

	// range
	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index", i, "value: ", v)
	}
}
