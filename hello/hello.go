package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	//log.SetOutput(io.Discard)

	message, err := greetings.Hello("Daniil")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
