package reader

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filePath string) []byte {

	dat, err := os.ReadFile(filePath)
	check(err)
	return dat
}
