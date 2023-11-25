
package main

import (
	"fmt"
	"github.com/sharhan-alhassan/golang-projects/go-docs/02-greetings"
)

func main()  {
	// Get a greeting message and print it
	message := greetings.Hello("Sharhan")
	fmt.Println(message)
}