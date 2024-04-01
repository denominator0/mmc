package mmc

func Chunk[E any, S ~[]E](s S, chunkLimit int, chunkF func(time int, iteras ...Itera[E])) {
	ChunkE(s, chunkLimit, func(time int, iteras ...Itera[E]) error { chunkF(time, iteras...); return nil })
}

func ChunkE[E any, S ~[]E](s S, chunkLimit int, chunkF func(time int, iteras ...Itera[E]) error) error {
	return BatchE(len(s), chunkLimit, func(time, i, j int) error {
		return chunkF(time, Map(s[i:j], func(itera Itera[E]) Itera[E] { itera.Pos += (time - 1) * chunkLimit; return itera })...)
	})
}

// Deprecated: use ChunkE instead
func OldChunkE[E any, S ~[]E](s S, chunkLimit int, chunkF func(time int, iteras ...Itera[E]) error) error {
	if len(s) == 0 || chunkLimit == 0 {
		return nil
	}
	if len(s) <= chunkLimit {
		return chunkF(1, Map(s, ID)...)
	}
	if err := chunkF(1, Map(s[:chunkLimit], ID)...); err != nil {
		return err
	}
	return ChunkE(s[chunkLimit:], chunkLimit, func(time int, iteras ...Itera[E]) error {
		for i := range iteras {
			iteras[i].Pos += chunkLimit
		}
		return chunkF(time+1, iteras...)
	})
}
