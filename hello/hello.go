package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Daniil")
	fmt.Println(message)
}
