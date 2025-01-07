package utils

func Iterator[V any](numbers []V) func() (V, bool) {
	i := 0

	return func() (V, bool) {
		if i >= len(numbers) {
			var zero V 
			return zero, false
		}

		val := numbers[i]
		i++
		return val, true
	}
}
