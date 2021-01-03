package validator

import "os"

func ValidateFilesPath(paths []string) (bool, string) {
	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return false, path
		}
	}
	return true, ""
}

func ValidateArgumentsSize(num int) bool {
	return num == 2
}
