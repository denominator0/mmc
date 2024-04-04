package mmc

func Map[E, AfterE any, S ~[]E](slice S, m func(itera Itera[E]) AfterE) []AfterE {
	result, _ := MapE(slice, func(itera Itera[E]) (AfterE, error) { return m(itera), nil })
	return result
}

func MapE[E, AfterE any, S ~[]E](slice S, m func(itera Itera[E]) (AfterE, error)) ([]AfterE, error) {
	return FoldE(slice, func(sum []AfterE, itera Itera[E]) ([]AfterE, error) {
		piece, err := m(itera)
		if err != nil {
			return nil, err
		}
		return append(sum, piece), nil
	})(make([]AfterE, 0, len(slice)))
}
