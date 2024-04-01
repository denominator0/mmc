package mmc

func Batch(length int, limit int, batchF func(time, i, j int)) {
	BatchE(length, limit, func(time, i, j int) error { batchF(time, i, j); return nil })
}

func BatchE(length int, limit int, batchF func(time, i, j int) error) error {
	if length == 0 || limit == 0 {
		return nil
	}
	if length <= limit {
		return batchF(1, 0, length)
	}
	if err := batchF(1, 0, limit); err != nil {
		return err
	}
	return BatchE(length-limit, limit, func(time, i, j int) error {
		return batchF(time+1, i+limit, j+limit)
	})
}
