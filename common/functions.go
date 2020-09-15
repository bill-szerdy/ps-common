package common

// SliceContainsString traverses and array and returning if the value is in the array
func SliceContainsString(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}
