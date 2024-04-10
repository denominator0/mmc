package mmc

type Option[T any] func(*T)

func OPT[T any](opts ...Option[T]) T {
	return Fold(opts, func(sum T, itera Itera[Option[T]]) T {
		itera.Cur(&sum)
		return sum
	})(Var[T]())
}
