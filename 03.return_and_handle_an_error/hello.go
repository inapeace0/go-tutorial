package main

import (
	"fmt"
	"errors"
	"log"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name") 
	}

	message := fmt.Sprintf("Hi, %v, Welcome!", name)
	return message, nil
}

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := Hello("asdfasdf")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(message)
}