package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    if name == "" {
		return "", errors.New("empty name")
	}
    message := fmt.Sprintf(randFormat(), name)
    return message, nil
}

func randFormat() string {
	formats := []string{
		"こんにちは、%vさん。なんの用？",
		"%vさん、話しかけないでください。",
		"%vさん、大好きです。",
	}

	return formats[rand.Intn(len(formats))]
}