package convert

func Map[T, U any](src []T, fn func(T) U) []U {
	if len(src) == 0 {
		return nil
	}

	dst := make([]U, 0, len(src))
	for _, item := range src {
		dst = append(dst, fn(item))
	}
	return dst
}
