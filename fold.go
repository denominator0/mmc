package mmc

func Fold[E, SUM any, S ~[]E](slice S, fold func(sum SUM, itera Itera[E]) SUM) func(sum SUM) SUM {
	return func(sum SUM) SUM {
		sum, _ = FoldE(slice, func(sum SUM, itera Itera[E]) (SUM, error) {
			return fold(sum, itera), nil
		})(sum)
		return sum
	}
}

func FoldE[E, SUM any, S ~[]E](slice S, fold func(sum SUM, itera Itera[E]) (SUM, error)) func(sum SUM) (SUM, error) {
	return func(sum SUM) (SUM, error) {
		var err error
		return sum, ChunkE(slice, 1, func(_ int, iteras ...Itera[E]) error {
			sum, err = fold(sum, iteras[0])
			return err
		})
	}
}
