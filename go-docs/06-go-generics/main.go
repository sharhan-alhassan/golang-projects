package main

import "fmt"

//sumInts adds together the values of m
func sumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

//sumFloats adds together the values of m
func sumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func main()  {
	// Initialize a map for the integer values
	ints := map[string]int64 {
		"first": 34,
		"second": 98,
	}

	// Initialize a map for the float values
	floats := map[string]float64 {
		"first": 45.2,
		"second": 34.23,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
	sumInts(ints),
	sumFloats(floats))	
}