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
		return sum, BatchE(len(slice), 1, func(_ int, i, j int) error {
			sum, err = fold(sum, Itera[E]{Cur: slice[i], Pos: i})
			return err
		})
	}
}
