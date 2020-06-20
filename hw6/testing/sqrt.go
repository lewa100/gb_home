package testing

import (
	"github.com/sirupsen/logrus"
	"math"
)

//QECalculator - Quadratic Equation Calculator.
func QECalculator(a, b, c float64) (answer []float64, d float64) {
	d = math.Pow(b, 2) - 4*a*c
	switch true {
	case d < 0:
		{
			logrus.Info("Have not root, d = ",d)
		}
	case d == 0:
		{
			logrus.Info("Have x1, d = ",d)
			answer = append(answer, -b/(2*c))
		}
	case d > 0:
		{
			logrus.Info("Have x1, x2, d = ",d)
			answer = append(answer, (-b +math.Sqrt(d))/(2*a))
			answer = append(answer, (-b -math.Sqrt(d))/(2*a))
		}

	}
	return
}
