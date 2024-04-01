package mmc

import (
	"fmt"
	"testing"
)

func TestBatch(t *testing.T) {
	Batch(10, 3, func(time, i, j int) { fmt.Println(time, i, j) })
}
