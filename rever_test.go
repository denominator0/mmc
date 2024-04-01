package mmc

import (
	"reflect"
	"testing"
)

func TestRever(t *testing.T) {
	if l := len(Rever([]int64{})); l != 0 {
		t.Errorf("it should be empty, but it has %d", l)
	}
	if result := Rever(([]int64{5})); !reflect.DeepEqual(result, []int64{5}) {
		t.Errorf("it should be []int64{5}, but it has %d", result)
	}
	if result := Rever(([]int64{5, 10})); !reflect.DeepEqual(result, []int64{10, 5}) {
		t.Errorf("it should be []int64{10, 5}, but it has %d", result)
	}
	if result := Rever(([]int64{5, 3, 10})); !reflect.DeepEqual(result, []int64{10, 3, 5}) {
		t.Errorf("it should be []int64{10, 3, 5}, but it has %d", result)
	}
}
