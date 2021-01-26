package greetings

import (
	 "fmt"
	 "errors"
)

// Hello returns a greeting for the name person.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}
	// Returns a greeting that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}