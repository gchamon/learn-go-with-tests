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

	checkArea := func(t testing.TB, shape Shape, want float64) {
		got := shape.Area()

		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rec := Rectangle{12, 6}
		want := 72.0
		checkArea(t, rec, want)
	})
	t.Run("circles", func(t *testing.T) {
		cir := Circle{10}
		want := 314.1592653589793
		checkArea(t, cir, want)
	})

}
