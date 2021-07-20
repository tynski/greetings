package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greetings for the named person.

func Hello(name string) string {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("epmty name")
	}

	// If a name was reveived, return a value that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
