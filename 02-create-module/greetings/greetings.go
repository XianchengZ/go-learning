package greetings

// declared greetings package to collect related functions
// implement a hello function to return the greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // Return a greeting that embeds the name in a message.
		// if no name was given, return an error with a message

		if name == "" {
			return "", errors.New("empty name")
		}

    message := fmt.Sprintf("Hi, %v. Welcome!", name)
		// := operator is a shortcut for declaring and initializing a variable in one line.
		
		// alternatively:
		// var message string
		// message = fmt.Sprintf("Hi, %v. Welcome!", name)
    
		return message, nil
}