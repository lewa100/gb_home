package os

import (
	"fmt"
	"reflect"
	"strings"
)

// oSystem - Operating system.
type OSystem interface {
	StartOS()
	GetStatus() bool
}

// PrintStatus - Print status os
func PrintStatus(os OSystem) {
	if !os.GetStatus() {
		fmt.Printf("%s is SHUTDOWN\n", strings.ToUpper(reflect.TypeOf(os).String()[1:]))
		return
	}
	fmt.Printf("%s is WORK\n", strings.ToUpper(reflect.TypeOf(os).String()[1:]))
}
