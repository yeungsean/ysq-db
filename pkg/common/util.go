package common

// VarArgGetFirst ...
func VarArgGetFirst[T any](args ...T) T {
	var zero T
	if len(args) > 0 {
		zero = args[0]
	}
	return zero
}
