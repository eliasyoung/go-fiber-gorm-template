package pkg

func NilSliceFormater[T any](slice []T) []T {
	if len(slice) == 0 {
		return make([]T, 0)
	} else {
		return slice
	}
}
