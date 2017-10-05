package slices

// StringInSlice returns if needle is in the slice haystack.
func StringInSlice(needle string, haystack []string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

// IntInSlice returns if needle is in the slice haystack.
func IntInSlice(needle int, haystack []int) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

// ReverseStringSlice reverses a slice of strings.
func ReverseStringSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
