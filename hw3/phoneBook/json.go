package phoneBook

import (
	"encoding/json"
)

// saveInJSON - converter struct to JSON.
func saveInJSON(){
	jsonData, err := json.Marshal(PhBook.pList)

	if err != nil {
		panic(err)
	}

	saveToFile(jsonData)
}