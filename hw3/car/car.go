package car

var Lada, Mazda *car

type car struct {
	Brand        string
	Year         int
	BodyVolume   int
	IsRunning    bool
	WindowIsOpen bool
	BootVolume   int
}

func init(){
	Lada = &car{
		Brand:        "LADA",
		Year:         2000,
		BodyVolume:   400,
		IsRunning:    false,
		WindowIsOpen: true,
		BootVolume:   70,
	}

	Mazda = &car{
		Brand:        "MAZDA",
		Year:         2011,
		BodyVolume:   500,
		IsRunning:    false,
		WindowIsOpen: false,
		BootVolume:   35,
	}
}

