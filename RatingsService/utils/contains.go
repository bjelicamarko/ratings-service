package utils

func Contains(array []string, target string) bool {
	for _, element := range array {
		if element == target {
			return true
		}
	}
	return false
}
