package main

import "fmt"

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "
const swedishHelloPrefix = "Hej, "
const suffix = "!"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	case "Swedish":
		prefix = swedishHelloPrefix
	}

	return prefix + name + suffix
}

func main() {
	fmt.Println(Hello("Nick", "English"))
}
