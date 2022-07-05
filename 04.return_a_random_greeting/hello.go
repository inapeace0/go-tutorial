package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomFormat() string {
	rand.Seed(time.Now().UnixNano())
	formats := []string{
		"Hi, %v. Welcome!",
		"Greeat to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}

func main() {
	name := "Sun"
	message := fmt.Sprintf(randomFormat(), name)
	fmt.Println(message)
}