package mmc

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	Chunk([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4, func(time int, iteras ...Itera[int64]) {
		if time == 1 {
			if l := len(iteras); l != 4 {
				t.Errorf("it should has length 4, but it is %d\n", l)
			}
			if !reflect.DeepEqual(iteras, []Itera[int64]{
				{Cur: 1, Pos: 0}, {Cur: 2, Pos: 1},
				{Cur: 3, Pos: 2}, {Cur: 4, Pos: 3},
			}) {
				t.FailNow()
			}
		}
		if time == 2 {
			if l := len(iteras); l != 4 {
				t.Errorf("it should has length 4, but it is %d\n", l)
			}
			if !reflect.DeepEqual(iteras, []Itera[int64]{
				{Cur: 5, Pos: 4}, {Cur: 6, Pos: 5},
				{Cur: 7, Pos: 6}, {Cur: 8, Pos: 7},
			}) {
				t.FailNow()
			}
		}
		if time == 3 {
			if l := len(iteras); l != 1 {
				t.Errorf("it should has length 1, but it is %d\n", l)
			}
			if !reflect.DeepEqual(iteras, []Itera[int64]{{Cur: 9, Pos: 8}}) {
				t.FailNow()
			}
		}
	})
	Chunk([]int64{1, 2, 3, 4}, 4, func(time int, iteras ...Itera[int64]) {
		if time != 1 {
			t.Errorf("it should be 1, but it is %d", time)
		}
		if !reflect.DeepEqual(iteras, []Itera[int64]{
			{Cur: 1, Pos: 0}, {Cur: 2, Pos: 1},
			{Cur: 3, Pos: 2}, {Cur: 4, Pos: 3},
		}) {
			t.FailNow()
		}
	})
	Chunk([]int64{1, 2, 3, 4}, 0, func(time int, iteras ...Itera[int64]) {
		t.FailNow()
	})

}
