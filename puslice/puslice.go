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

func Contains[T comparable](slice []T, val T) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}
