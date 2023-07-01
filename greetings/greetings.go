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

func Hellos(names []string) (map[string]string, error) {
	// mapについて：https://go.dev/blog/maps
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randFormat() string {
	formats := []string{
		"こんにちは、%vさん。なんの用？",
		"%vさん、話しかけないでください。",
		"%vさん、大好きです。",
	}

	return formats[rand.Intn(len(formats))]
}