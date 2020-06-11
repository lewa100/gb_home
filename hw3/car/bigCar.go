package car

var Kamaz *BigCar

type BigCar struct {
	Brand        string
	Year         int
	BodyVolume   int
	IsRunning    bool
	WindowIsOpen bool
	BootVolume   int
	NumberWheels int
	IsCrane      bool
}

func init(){
	Kamaz = &BigCar{
		Brand:        "KAMAZ",
		Year:         2013,
		BodyVolume:   65115,
		IsRunning:    true,
		WindowIsOpen: false,
		BootVolume:   30,
		NumberWheels: 6,
		IsCrane: false,
	}
}