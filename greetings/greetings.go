package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("There is no name to greet")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	rand.Seed(time.Now().UnixNano())

	formats := []string{
		"Hi, %v Welcome!",
		"Great to see you %s",
		"Hil, %v! Well met!!!",
	}

	return formats[rand.Intn(len(formats))]
}
