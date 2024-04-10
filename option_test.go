package mmc

import "testing"

func TestOption(t *testing.T) {
	type opt struct {
		first  bool
		second int
	}
	func(foo string, options ...Option[opt]) {
		if foo != "bar" {
			t.Fatalf("it should be 'bar', but it is %s", foo)
		}
		opt := OPT(options...)
		if !opt.first {
			t.Fatalf("it should be true, but it is %t", opt.first)
		}
		if opt.second != 2 {
			t.Fatalf("it should be true, but it is %d", opt.second)
		}
	}("bar", func(o *opt) { o.first = true }, func(o *opt) { o.second = 2 })
}
