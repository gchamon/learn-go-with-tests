package util

import "testing"

func AssertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertCorrectMessageIntGivenArray(t testing.TB, got, want int, given []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d, given %v", got, want, given)
	}
}
