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

	return greetingPrefix(language) + name + suffix
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	case "Swedish":
		prefix = swedishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Nick", "English"))
}
