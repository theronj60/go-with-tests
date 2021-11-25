package iteration

import (
	"testing"
)

func Repeat(character string, iterations int) string {
	var repeated string
	for i := 0; i < iterations; i++ {
		repeated += character
	}
	return repeated
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 7)
	expected := "aaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}

// func ExampleRepeat() {
// 	for i := 0; i < 5; i++ {
// 		Repeat("a", 5)
// 	}
//     fmt.Println()
//     // Output: aaaaa
// }

