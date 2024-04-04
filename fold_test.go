package mmc

import "testing"

func TestFold(t *testing.T) {
	if sum := Fold([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(sum int64, itera Itera[int64]) int64 { return sum + itera.Cur })(0); sum != 45 {
		t.Fatalf("it should be 45, but it is %d", sum)
	}
	if sum := Fold([]int64{}, func(sum int64, itera Itera[int64]) int64 { return sum + itera.Cur })(10); sum != 10 {
		t.Fatalf("it should be 10, but it is %d", sum)
	}
}
