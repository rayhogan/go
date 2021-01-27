package greetings

import (
	 "fmt"
	 "errors"
	 "math/rand"
	 "time"
)

// Hello returns a greeting for the name person.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}
	// Returns a greeting that embeds the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init(){
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}