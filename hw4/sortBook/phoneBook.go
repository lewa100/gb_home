package sortBook

type Person struct {
	Name  string
	Phone int
}

//
type PhoneBook []Person

func (a PhoneBook) Len() int {
	return len(a)
}
func (a PhoneBook) Less(i, j int) bool {
	return a[i].Phone < a[j].Phone
}
func (a PhoneBook) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
