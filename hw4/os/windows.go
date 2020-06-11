package os

// Windows - OS struct
type Windows struct {
	Name       string
	Version    string
	RealStatus bool
}

// StartOS for start OS.
func (u *Windows) StartOS(){
	u.RealStatus = true
}

// GetStatus for detect real status OS.
func (u *Windows) GetStatus() bool{
	return u.RealStatus
}

// StopOS for shutdown OS.
func (u *Windows) StopOS() {
	u.RealStatus = false
}