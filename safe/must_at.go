package safe

var PanicOutOfRange = "index out of range"

func MustAt[T any](xs []T, i int) T {
	if i < 0 || i >= len(xs) {
		panic(PanicOutOfRange)
	}
	return xs[i]
}
