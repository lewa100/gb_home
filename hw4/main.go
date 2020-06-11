package main

import (
	"fmt"
	"github.com/gb_home/hw4/calc"
	"github.com/gb_home/hw4/chess"
	"github.com/gb_home/hw4/os"
	"github.com/gb_home/hw4/sortBook"
	"sort"
)

var (
	macOS *os.Unix
	win10 *os.Windows
)


func initMacOS() {
	macOS = &os.Unix{
		Name:           "macOS",
		Version:        "14.15.2",
		RealStatus:     false,
		ProgForInstall: "brew",
	}
}

func initWin() {
	win10 = &os.Windows{
		Name:       "Windows 10",
		Version:    "2541.10",
		RealStatus: true,
	}
}

func activeCalculate(){
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "help"{
			calc.GetInfo()
			continue
		}

		if input == "exit" {
			break
		}

		if res, err := calc.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}

func main() {
	// macOS.
	initMacOS()
	macOS.StartOS()
	os.PrintStatus(macOS)

	// Windows 10.
	initWin()
	win10.StopOS()
	os.PrintStatus(win10)

	//sorting by phone
	book := sortBook.PhoneBook{
		{"Pasha", 89991112233},
		{"Gosha", 89121112233},
		{"Tosha", 89161112233},
		{"Ksusha", 89167772233},
		{"Tolya", 89166456233},
		{"Zoya", 89161143233},
	}
	sort.Sort(book)
	fmt.Println(book) // [{Gosha 89121112233} {Tosha 89161112233} {Zoya 89161143233} {Tolya 89166456233} {Ksusha 89167772233} {Pasha 89991112233}]

	//Calc
	activeCalculate()

	//Chess
	horse := &chess.Chess{}
	horse.GetPos()
	horse.PrintPos()
	fmt.Println("True positions: ", horse.SearchForOptions())
}
