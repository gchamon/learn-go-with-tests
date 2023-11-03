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
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.shape.Area()
			want := tc.want
			if got != want {
				t.Errorf("%#v, got %g, want %g", tc.shape, got, want)
			}
		})
	}
}
