package utils

func Min[T int | uint32 | int64 | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}
