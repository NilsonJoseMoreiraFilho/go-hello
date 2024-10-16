package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	message, err := greetings.Hello("Nilson")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	names := []string{"Nilson", "Vinicius", "Nayara"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for name, message := range messages {
		fmt.Printf("%s: %s\n", name, message)
	}

}
