package mmc

import (
	"testing"
)

func TestBatch(t *testing.T) {
	// batch that the last time is not enough
	Batch(10, 3, func(time, i, j int) {
		if time == 1 && (i != 0 || j != 3) {
			t.Fatalf("it should be time:1, i:0, j:3, but it is time:%d, i:%d, j:%d", time, i, j)
		}
		if time == 2 && (i != 3 || j != 6) {
			t.Fatalf("it should be time:2, i:3, j:6, but it is time:%d, i:%d, j:%d", time, i, j)
		}
		if time == 3 && (i != 6 || j != 9) {
			t.Fatalf("it should be time:3, i:6, j:9, but it is time:%d, i:%d, j:%d", time, i, j)
		}
		if time == 4 && (i != 9 || j != 10) {
			t.Fatalf("it should be time:4, i:9, j:10, but it is time:%d, i:%d, j:%d", time, i, j)
		}
	})

	// batch all
	Batch(3, 3, func(time, i, j int) {
		if time == 1 && (i != 0 || j != 3) {
			t.Fatalf("it should be time:1, i:0, j:3, but it is time:%d, i:%d, j:%d", time, i, j)
		}
		if time != 1 {
			t.FailNow()
		}
	})

	// batch nothing if length is 0
	Batch(0, 3, func(time, i, j int) {
		t.FailNow()
	})

	// batch nothing if limit is 0
	Batch(3, 0, func(time, i, j int) {
		t.FailNow()
	})
}
