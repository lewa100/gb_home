package helper

import (
	"bytes"
	"fmt"
	"github.com/gb_home/hw5/openFile"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	OldName = ""
	NewName = ""
)

// saveToNewFile - save data to file NewName.
func saveToNewFile(data []byte){
	dataFile, err := os.Create(NewName)

	if err != nil {
		panic(err)
	}
	defer dataFile.Close()

	dataFile.Write(data)
	dataFile.Close()
	fmt.Println("Copy file to", dataFile.Name() )
}
// CopyFile - check file and copy data to buffer.
func CopyFile(){
	if openFile.FileIsNotExist(OldName){
		return
	}

	file, err := os.Open(OldName)
	if err != nil {
		return
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("ERROR with save")
	}
	saveToNewFile(data)
}

// RunCMD for start command.
func RunCMD(args []string) (out string, err error) {
	var stdout bytes.Buffer
	cmd := exec.Command(args[0],args[1:]...)
	cmd.Stdout = &stdout

	err = cmd.Run()
	if err != nil {
		return stdout.String(), err
	}

	return stdout.String(), nil
}