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
