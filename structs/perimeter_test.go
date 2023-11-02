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
	t.Run("rectangles", func(t *testing.T) {
		rec := Rectangle{12, 6}
		got := rec.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {
		cir := Circle{10}
		got := cir.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})
}
