package utils

// 値がスライス内に存在するかどうか判定
func IsSliceContains[T comparable](slice []T, value T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
