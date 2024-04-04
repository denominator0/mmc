package mmc

func Map[BEF, AFT any, S ~[]BEF](slice S, m func(itera Itera[BEF]) AFT) []AFT {
	result, _ := MapE(slice, func(itera Itera[BEF]) (AFT, error) { return m(itera), nil })
	return result
}

func MapE[BEF, AFT any, S ~[]BEF](slice S, m func(itera Itera[BEF]) (AFT, error)) ([]AFT, error) {
	return FoldE(slice, func(sum []AFT, itera Itera[BEF]) ([]AFT, error) {
		piece, err := m(itera)
		if err != nil {
			return nil, err
		}
		return append(sum, piece), nil
	})(make([]AFT, 0, len(slice)))
}
