package services

import (
	"os"
)

/*
IsPathExist check whether the path is exist
*/
func IsPathExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
