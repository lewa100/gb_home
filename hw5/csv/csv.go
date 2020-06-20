package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/gb_home/hw5/openFile"
	. "os"
)

// Structs for CSV.
type domain struct {
	Domain string
	Type string
	Organization string
}

type DomainsList struct {
	Domains []domain
}

//ReadCSV read csv file.
func (ds *DomainsList)ReadCSV (path string){
	if openFile.FileIsNotExist(path){
		return
	}
	file, err := Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		d := &domain{
			Domain:       record[0],
			Type:         record[1],
			Organization: record[2],
		}

		ds.Domains = append(ds.Domains, *d)
	}
}

// Print - output DomainsList to console.
func (ds *DomainsList)Print(){
	for i, d := range ds.Domains{
		fmt.Println("<<<<< Line:",i + 1,">>>>>")
		fmt.Println("Domain:",d.Domain)
		fmt.Println("Type:",d.Type)
		fmt.Println("Organization:",d.Organization, "\n")
	}
}