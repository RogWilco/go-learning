package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Nick", "")
		want := "Hello, Nick!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Nicolás", "French")
		want := "Bonjour, Nicolás!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Nikito", "Spanish")
		want := "Hola, Nikito!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Swedish", func(t *testing.T) {
		got := Hello("Nils", "Swedish")
		want := "Hej, Nils!"

		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
