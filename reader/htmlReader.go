package reader

import (
	"errors"
	"os"
)

func isExistingFile(filename string) bool {
	_, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func ReadFile(filePath string) []byte {
	dat, err := os.ReadFile(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return make([]byte, 0)
	}
	return dat
}
