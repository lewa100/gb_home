package openFile

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestFileIsNotExist(t *testing.T) {
	path := GetPath() + "/filterRead.go"
	assert.False(t, FileIsNotExist(path),"File is exist")
}

func TestOpenFile(t *testing.T) {
	path := GetPath() + "/filterRead.go"
	buf := OpenFile(path)
	txt := `Where does it come from?
Contrary to popular belief,
Lorem Ipsum is not simply random text.
It has roots in a piece of classical
Latin literature from 45 BC, making it
over 2000 years old. Richard McClintock,
a Latin professor at Hampden-Sydney College
in Virginia, looked up one of the more
obscure Latin words, consectetur, from
a Lorem Ipsum passage, and going through`
	assert.Contains(t,buf, txt, "File is read")
}