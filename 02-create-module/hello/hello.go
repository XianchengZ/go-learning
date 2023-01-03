package main

import (
	"fmt"
	"log"

	"greetings"
)

func main() {
    // Get a greeting message and print it.
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

		message, err := greetings.Hello("Andrew")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
				// Fatal is equivalent to Print() followed by a call to os.Exit(1).
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}