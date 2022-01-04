package coalesce

func Coalesce[T any](a T, b T) T {
	if interface{}(a) == nil {
		return b
	}

	return a
}
