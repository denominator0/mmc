package mmc

func Rever[E any, S ~[]E](s S) S {
	if len(s) <= 1 {
		return s
	}
	return append(Rever(s[1:]), s[0])
}
