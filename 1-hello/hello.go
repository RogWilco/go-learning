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

	if language == "French" {
		return frenchHelloPrefix + name + suffix
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name + suffix
	}

	if language == "Swedish" {
		return swedishHelloPrefix + name + suffix
	}

	return englishHelloPrefix + name + suffix
}

func main() {
	fmt.Println(Hello("Nick", "English"))
}
