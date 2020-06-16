package testing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumArgs(t *testing.T) {
	type structForTest struct {
		args   []int
		result int
	}

	dataForTest := []structForTest{
		{
			args:   []int{3, 6, 7, 3, 10, -100},
			result: -71}, {
			args:   []int{311, 6, 3447, 33, 6, -432},
			result: 3371,
		},
	}
	t.Run("Success", func(t *testing.T) {
		for _, v := range dataForTest {
			f := SumArgs(v.args...)
			assert.Equal(t, v.result, f, "Its OK")
		}
	})

}
