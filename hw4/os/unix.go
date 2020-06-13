package os

// Unix - OS struct
type Unix struct {
	Name       string
	Version    string
	RealStatus bool
	ProgForInstall string
}

// StartOS for start OS.
func (u *Unix) StartOS() {
	u.RealStatus = true
}

// GetStatus for detect real status OS.
func (u *Unix) GetStatus() bool {
	return u.RealStatus
}

// StopOS for shutdown OS.
func (u *Unix) StopOS() {
	u.RealStatus = false
}
