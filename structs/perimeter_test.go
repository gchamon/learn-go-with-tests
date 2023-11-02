package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10, 10}
	got := Perimeter(rec)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rec := Rectangle{12, 6}
	got := Area(rec)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
