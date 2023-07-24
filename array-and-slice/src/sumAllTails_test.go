package arrayAndSlice

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v but got %v", want, got)
		}
	}

	t.Run("sum some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("sum an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
