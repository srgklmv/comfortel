package pointer

func ParsePointer[T any](val *T) T {
	if val == nil {
		var r T
		return r
	}

	return *val
}
