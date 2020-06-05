package main

import (
	"fmt"
	c "github.com/gb_home/hw3/car"
	ph "github.com/gb_home/hw3/phoneBook"
	q "github.com/gb_home/hw3/queue"
)

func main() {
	//LADA
	fmt.Println(*c.Lada)
	c.Lada.IsRunning = true
	c.Lada.WindowIsOpen = false
	fmt.Println(*c.Lada)

	//MAZDA
	fmt.Println(*c.Mazda)
	c.Mazda.IsRunning = true
	fmt.Println(*c.Mazda)

	//KAMAZ
	fmt.Println(*c.Kamaz)
	c.Kamaz.NumberWheels = 8
	c.Kamaz.IsCrane = true
	fmt.Println(*c.Kamaz)

	//Queue
	q.QR.AddInQueue("Bob", "79991112233")
	q.QR.AddInQueue("Toi", "79891165663")
	q.QR.AddInQueue("Kok", "79222355533")
	q.QR.AddInQueue("Zob", "70000000033")
		fmt.Println(q.QR)
	for i:= 0; i < 6; i++{
		tmp := q.QR.GetFromQueue()
		if tmp.Id != 0 {
			fmt.Println(tmp)
		}
	}

	//PhoneBook
	ph.LoadFromFile()

	ph.PhBook.AddPerson("Tex", "79991666233")
	ph.PhBook.AddPerson("Ios", "79895555663")
	ph.PhBook.AddPerson("Kros", "79222355533")
	ph.PhBook.AddPerson("Zdros", "72900000033")

	fmt.Println(ph.PhBook)
}