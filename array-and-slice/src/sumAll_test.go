package arrayAndSlice

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2}
	slice2 := []int{0, 9}

	got := SumAll(slice1, slice2)
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v but got %v", want, got)
	}
}
