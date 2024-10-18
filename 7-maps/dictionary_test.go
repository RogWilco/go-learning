package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("unknown")

		if got == nil {
			t.Fatal("expected an error")
		}

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{}

	dict.Add("test", "this is just another test")

	got, err := dict.Search("test")
	want := "this is just another test"

	if err != nil {
		t.Fatal("unexpected error")
	}

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
