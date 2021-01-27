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

    // Get a greeting message and print it.
    message, err := greetings.Hello("Ray")

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}