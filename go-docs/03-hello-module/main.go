
package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main()  {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
	log.SetPrefix("greetings: ")
    log.SetFlags(0)

	// A slice of names
	names := []string{"Sharhan", "kiwam", "Dada"}

	// Request a greeting message
	// message, err := greetings.Hello("Sharhan")
	messages, err := greetings.Hellos(names)

	// If an error occurred, print it and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error occurred, print the returned message
	// fmt.Println(message)
	fmt.Println(messages)
}