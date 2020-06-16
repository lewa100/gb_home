package testing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestQECalculator - test Quadratic Equation Calculator.
func TestQECalculator(t *testing.T) {
	type structForTest struct {
		a,b,c,d float64
		result []float64
	}

	t.Run("Have x1, x2", func(t *testing.T) {
		dataForTest := []structForTest{
			{a:4, b:5, c:1, d: 9,result: []float64{-0.25 , -1}},
			{a:5, b:100, c:43, d:9140, result: []float64{-0.43966527782630466, -19.560334722173696}},
		}
		for _, v := range dataForTest {
			f,d := QECalculator(v.a,v.b,v.c)
			assert.Equal(t, v.d, d, "Discriminant is OK")
			assert.Equal(t, v.result, f, "Its OK")
		}
	})

	t.Run("Have x1", func(t *testing.T) {
		dataForTest := []structForTest{
			{a:1, b:-6, c:9, d: 0,result: []float64{0.3333333333333333}},
			{a:1, b:12, c:36, d:0, result: []float64{-0.16666666666666666}},
		}
		for _, v := range dataForTest {
			f,d := QECalculator(v.a,v.b,v.c)
			assert.Equal(t, v.d, d, "Discriminant is OK")
			assert.Equal(t, v.result, f, "Its OK")
		}
	})

	t.Run("Have not root", func(t *testing.T) {
		dataForTest := []structForTest{
			{a:5, b:3, c:7, d: -131},
			{a:1, b:-2, c:10, d:-36},
		}
		for _, v := range dataForTest {
			f,d := QECalculator(v.a,v.b,v.c)
			assert.Equal(t, d,v.d, "Discriminant is OK")
			assert.Nil(t,  f, "Its OK")
		}
	})
}
