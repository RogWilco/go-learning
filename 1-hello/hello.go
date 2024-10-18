package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const suffix = "!"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name + suffix
	}

	return englishHelloPrefix + name + suffix
}

func main() {
	fmt.Println(Hello("Nick", "English"))
}
