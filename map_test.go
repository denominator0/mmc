package mmc

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	if s := Map([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(itera Itera[int64]) string {
		return strconv.FormatInt(itera.Cur, 10)
	}); !reflect.DeepEqual(s, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}) {
		t.Fatalf("it should be [\"1\", \"2\", \"3\", \"4\", \"5\", \"6\", \"7\", \"8\", \"9\"], but it is %s", s)
	}
	if s := Map([]int64{}, func(itera Itera[int64]) string {
		return strconv.FormatInt(itera.Cur, 10)
	}); len(s) != 0 {
		t.Fatalf("it should be empty, but it has %d", len(s))
	}
}
