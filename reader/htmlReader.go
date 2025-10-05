package reader

import (
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return dat, nil
}
