package structs

type Rectangle struct {
	height float64
	width  float64
}

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.height + rec.width)
}

func Area(rec Rectangle) float64 {
	return rec.height * rec.width
}
