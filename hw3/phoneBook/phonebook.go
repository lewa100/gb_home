package phoneBook

var (
	PhBook phoneBook
	lastId = 1
)

type Person struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type phoneBook struct {
	pList []Person
}

// AddPerson add Person in PhoneBook and save in JSON.
func (ph *phoneBook) AddPerson( name, phone string){
	tmp := Person{
		Id: lastId,
		Name:  name,
		Phone: phone,
	}
	lastId++
	ph.pList = append(ph.pList, tmp)
	saveInJSON()
}