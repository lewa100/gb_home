package phoneBook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const fileName = "./Person.json"

// saveToFile - save JSON i file.
func saveToFile(jsn []byte){
	jsonFile, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsn)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())
}

// LoadFromFile - load JSON from file and save in PhoneBook and lastId.
func LoadFromFile(){
	file, _ := ioutil.ReadFile(fileName)
	var data []Person

	err := json.Unmarshal(file, &data)

	if err != nil {
		panic(err)
	}

	for _, v := range data{
		PhBook.pList = append(PhBook.pList, v)
	}
	lastId = PhBook.pList[len(PhBook.pList)-1].Id + 1
}