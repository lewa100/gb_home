package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRunCMD(t *testing.T) {
	cmd := "grep -i command helper.go"
	args := strings.Split(cmd," ")
	out , _ := RunCMD(args)
	fmt.Println(out)
	txt := `// RunCMD for start command.`
	assert.Contains(t, out, txt,"Grep is work")
}
