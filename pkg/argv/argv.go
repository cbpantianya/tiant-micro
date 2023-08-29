package argv

import (
	"errors"
)

// Decode file path from argv
func DecodeFilePath(argv []string)(path string, err error){
	// Find tag -f
	index := -1
	for i, v := range argv {
		if v == "-f" {
			index = i + 1
			break
		}
	}
	if index >= len(argv) || index == -1{
		err = errors.New("-f parameter is empty")
		return
	}
	path = argv[index]
	return
}