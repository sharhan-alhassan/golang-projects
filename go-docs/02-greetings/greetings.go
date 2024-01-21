package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == ""{
		return "", errors.New("no name provided, please input a name")
	}

	// If a name was received, create a random greeting and return it
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}


// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map(dictionary) to associate names with messages
	messages := make(map[string]string)
	// Loop through the received slice of names, calling the Hello
	// function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrievd message with the name
		messages[name] = message
	}
	return messages, nil
}


// randomGreeting returns a random greeting message
func randomGreeting() string {
	// A slice of message formats
	greetings := []string {
		"Hi, %v. Welcome!",
		"Nice to see you, %v!",
		"Good day, %v!",
	}
	// return the randomly selected message
	return greetings[rand.Intn(len(greetings))]
}