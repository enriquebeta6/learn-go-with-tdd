package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("should repeat 5 characters", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertRepeatAsExpected(t, repeated, expected)
	})

	t.Run("should repeat 10 characters", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		assertRepeatAsExpected(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func assertRepeatAsExpected(t testing.TB, repeated, expected string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
