package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("retrieve existing word", func(t *testing.T) {
		assertDefinition(t, dictionary, "test", "this is just a test")
	})

	t.Run("retrieve unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")
	assertDefinition(t, dictionary, "test", "this is just a test")
}

func assertDefinition(t testing.TB, dict Dictionary, entry, definition string) {
	if got, err := dict.Search(entry); err != nil {
		t.Fatal("should find word:", err)
	} else {
		assertStrings(t, got, definition)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}
