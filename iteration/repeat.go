package iteration

func Repeat(a string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += a
	}
	return repeated
}
