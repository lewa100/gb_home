package openFile

import (
	"fmt"
	"os"
)

// GetPath get pwd for program.
func GetPath() string{
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
}

// FileIsNotExist detect file in directory.
func FileIsNotExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File %s is not found\n", path)
			return true
		}
	}
	return false
}

// OpenFile output file in console.
func OpenFile(path string) string{
	file, err := os.Open(path)
	if err != nil {
		return "Error with OpenFile"
	}
	defer file.Close()

	// getting size of file
	stat, err := file.Stat()
	if err != nil {
		return "Error with Stat"
	}

	// reading file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return "Error with read file"
	}

	return string(bs)
}
