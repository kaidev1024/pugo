package puslice

func Filter[T any](input []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range input {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Contains[T comparable](slice []T, val T) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
