package filesystem

import "os"

// FileExists is a simple check for the existence of a file
func FileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// FolderExists is a check if a folder exists
func FolderExists(foldername string) bool {
	info, err := os.Stat(foldername)

	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}
