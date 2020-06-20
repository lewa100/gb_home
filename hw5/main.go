package main

import (
	"flag"
	"fmt"
	"github.com/gb_home/hw5/csv"
	"github.com/gb_home/hw5/helper"
	of "github.com/gb_home/hw5/openFile"
	"os"
)

const (
	filename = "/openfile/filterRead.go"
	filecsv  = "/csv/csvFile.csv"
)

func flagInit() {
	tmpOldName := flag.String("in", "", "Old filename for copy")
	tmpNewName := flag.String("out", "", "New filename for copy")
	tmpGrep := flag.Bool("grep", false, "Grep for find text in file")
	flag.Parse()
	if *tmpOldName != "" {
		path := of.GetPath() + "/copy/"
		helper.OldName = path + *tmpOldName
		if *tmpNewName == "" {
			*tmpNewName = *tmpOldName + "_copy"
		}
		helper.NewName = path + *tmpNewName
		helper.CopyFile()
	}
	if *tmpGrep {
		var args []string
		args = append(args, "grep","-i")
		args = append(args, flag.Args()...)
		out , err := helper.RunCMD(args)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(out)
	}
}

func main() {
	if len(os.Args) > 1 {
		flagInit()
		os.Exit(1)
	}
	//Open file.
	fullpath := of.GetPath() + filename
	if of.FileIsNotExist(fullpath) {
		return
	}
	text := of.OpenFile(fullpath)
	fmt.Println(text)

	//CSV
	csvPath := of.GetPath() + filecsv
	domainList := &csv.DomainsList{}
	domainList.ReadCSV(csvPath)
	domainList.Print()

}
