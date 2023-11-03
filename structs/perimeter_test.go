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

	testCases := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36},
	}

	for _, tc := range testCases {
		got := tc.shape.Area()
		want := tc.want
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}
}
