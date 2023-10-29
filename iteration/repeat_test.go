package iteration

import (
	"fmt"
	"learn/util"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"
	util.AssertCorrectMessage(t, repeated, expected)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat("foo", 2)
	fmt.Println(repeated)
	// Output: foofoo
}
