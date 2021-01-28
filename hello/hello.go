package main

import (
    "fmt"
    "log"
    "example.com/greetings"
)

func main() {
    // Logger setup
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Get a greeting message and print it.
    messages, err := greetings.Hellos(names)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(messages)
}