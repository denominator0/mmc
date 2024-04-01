package mmc

func Filter[E any, S ~[]E](slice S, okF func(itera Itera[E]) bool) S {
	result, _ := FilterE(slice, func(itera Itera[E]) (bool, error) { return okF(itera), nil })
	return result
}

func FilterE[E any, S ~[]E](slice S, okF func(itera Itera[E]) (bool, error)) (S, error) {
	return FoldE(slice, func(sum S, itera Itera[E]) (S, error) {
		ok, err := okF(itera)
		if !ok || err != nil {
			return sum, err
		}
		return append(sum, itera.Cur), nil
	})(make(S, 0, len(slice)))
}
