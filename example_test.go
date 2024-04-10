package mmc

import (
	"fmt"
)

func ExampleBatch() {
	input := []string{"H", "e", "l", "l", "o", "W", "o", "r", "l", "d", "!"}
	Batch(len(input), 5, func(time, i, j int) {
		fmt.Println(time, input[i:j])
	})
	// Output:
	// 1 [H e l l o]
	// 2 [W o r l d]
	// 3 [!]
}

func ExampleChunk() {
	Chunk([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4, func(time int, iteras ...Itera[int64]) {
		fmt.Println(time, iteras)
	})
	// Output:
	// 1 [{1 0} {2 1} {3 2} {4 3}]
	// 2 [{5 4} {6 5} {7 6} {8 7}]
	// 3 [{9 8}]
}

func ExampleFilter() {
	fmt.Println(Filter([]int64{5, 6, 7, 3}, func(itera Itera[int64]) bool { return itera.Cur < 6 }))
	// Output:
	// [5 3]
}

func ExampleFold() {
	fmt.Println(Fold([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(sum int64, itera Itera[int64]) int64 { return sum + itera.Cur })(0))
	// Output:
	// 45
}

func ExampleMap() {
	fmt.Println(Map([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(itera Itera[int64]) int64 {
		return itera.Cur + 1
	}))
	// Output:
	// [2 3 4 5 6 7 8 9 10]
}

func ExampleOPT() {
	type opt struct {
		first  bool
		second int
	}
	func(foo string, options ...Option[opt]) {
		fmt.Println(foo, OPT(options...))
	}("bar", func(o *opt) { o.first = true }, func(o *opt) { o.second = 2 })
	// Output:
	// bar {true 2}
}
