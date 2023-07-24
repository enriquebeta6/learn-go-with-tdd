package arrayAndSlice

import "testing"

func TestSum(t *testing.T) {
	t.Run("should return 10 with an array of 1, 2, 3, 4", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
