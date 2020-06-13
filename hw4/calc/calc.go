package calc

import (
	// Пакет для обработки исходных данных
	"fmt"
	"go/parser"
)

type Func struct {
	Name string
	Args int
	Func func(args ...float64) float64
}

const info = `
<<<<<REFERENCE>>>>>
The calculator is able to calculate both ordinary 
operations (addition, subtraction, multiplication, division) 
and to pull out operations with constants:
pi		- Pi  = 3.141592653589793
e		- E   = 2.718281828459045
phi		- Phi = 1.618033988749895
Operations with functions:
sqrt(x)		- Square root x
abs(x)		- modulus of x
log(x y) 	- logarithm x of y
ln(x) 		- decimal logarithm x
sin(x)		- the sine of x
cos(x) 		- the cosine of x
tan(x) 		- tangent of x
arcsin(x)	- arcsine of x
arccos(x)	- the acrosine of x
arctan(x)	- arctangent of x
max (x1 ... xn) - the maximum number among the entered values
min (x1 ... xn) - the minimum number among the entered values

exit 		- Exit from Calculate
`

// GetInfo - get REFERENCE calculate
func GetInfo(){
	fmt.Println(info)
}

var funcMap map[string]Func

// Calculate рассчитает данные выражения, переданного в строке
func Calculate(expr string) (float64, error) {
	// Поиск управляющих операторов в строке и построение дерева
	// Вернет корневой элемент в переменную root все математические действия, построенные
	// в виде дерева
	// Node будет содержать в себе действие для вычисления
	// exprNode() даст доступ к ветвям дерева, исходящим из корня/ноды
	//
	// type Expr interface {
	//  Node
	//  exprNode()
	// }
	root, err := parser.ParseExpr(expr)


	if err != nil {
		return -1, err
	} else {
		return eval(root)
	}
}




