package mmc

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	if aft := Filter([]int64{5, 6, 7, 3}, func(itera Itera[int64]) bool { return itera.Cur < 6 }); !reflect.DeepEqual(aft, []int64{5, 3}) {
		t.Fatalf("it should be [5, 3], but it is %d", aft)
	}
	if aft := Filter([]int64{}, func(itera Itera[int64]) bool { return itera.Cur < 6 }); !reflect.DeepEqual(aft, []int64{}) {
		t.Fatalf("it should be [], but it is %d", aft)
	}
}
