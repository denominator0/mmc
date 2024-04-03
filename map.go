package mmc

func Map[BEF, AFT any, S ~[]BEF](slice S, m func(itera Itera[BEF]) AFT) []AFT {
	result, _ := MapE(slice, func(itera Itera[BEF]) (AFT, error) { return m(itera), nil })
	return result
}

func MapE[BEF, AFT any, S ~[]BEF](slice S, m func(itera Itera[BEF]) (AFT, error)) ([]AFT, error) {
	if len(slice) == 0 {
		return []AFT{}, nil
	}
	head, err := m(Itera[BEF]{Cur: slice[0]})
	if err != nil {
		return nil, err
	}
	tail, err := MapE(slice[1:], func(itera Itera[BEF]) (AFT, error) {
		itera.Pos += 1
		return m(itera)
	})
	if err != nil {
		return nil, err
	}
	result := append(make([]AFT, 0, 1+len(tail)), head)
	return append(result, tail...), nil
}
