package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(greeting string)  {
	fmt.Printf("Good morning, %v \n", greeting)
}

func sayBye(bye string)  {
	fmt.Printf("Goodbye %v \n", bye)
}

func cycleNames(n []string, f func(string))  {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func pointerExample(x *string) string {
	return *x
}

func main()  {

	fn, ln := getInitials("sharhan alhassan")
	fmt.Println("Initials is: ", fn, ln)

	greeting := "Hello, sharhan is good"
	// ages := []{5, 3, 67, 23, 4, 6, 20, 99, 234, 90}

	fmt.Println(strings.Contains(greeting, "Hello"))
	sayGreeting("Sharhan")
	sayBye("Mom")
	cycleNames([]string{"cloud", "azure", "pipeline"}, sayGreeting)
	cycleNames([]string{"cloud", "azure", "pipeline"}, sayBye)

	// maps
	menu := map[string]float64 {
		"soup": 4.33,
		"lace": 55.22,
		"talia": 98.33,
	}

	for k, v := range menu {
		fmt.Println("Key: ",k, "value: ", v)
	}

	name := "Sharhan"
	pointer := &name
	fmt.Printf("The memory of %v is: %v \n", name, pointer)
	fmt.Printf("The original value: %v \n", *pointer)

	fmt.Println(pointerExample(pointer))
}