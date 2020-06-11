package calc

import (
	"errors"
	"go/ast"
	"go/token"
	"math"
	"strconv"
	"strings"
)

// Разбор полученных данных
func eval(expr ast.Expr) (float64, error) {
	switch expr.(type) {
	case *ast.BasicLit:
		return basic(expr.(*ast.BasicLit))
	// Бинарные выражения
	case *ast.BinaryExpr:
		return binary(expr.(*ast.BinaryExpr))
	// Вложенные вычисления
	case *ast.ParenExpr:
		return eval(expr.(*ast.ParenExpr).X)
	// Обработка посредством математических функций
	case *ast.CallExpr:
		return call(expr.(*ast.CallExpr))
	// Случай для обработки констант
	case *ast.Ident:
		return ident(expr.(*ast.Ident))
	default:
		return -1, errors.New("Не удалось распознать оператор")
	}
}


// Разбор чисел
func basic(lit *ast.BasicLit) (float64, error) {
	switch lit.Kind {
	case token.INT:
		i, err := strconv.ParseInt(lit.Value, 10, 64)

		if err != nil {
			return -1, err
		} else {
			return float64(i), nil
		}
	case token.FLOAT:
		i, err := strconv.ParseFloat(lit.Value, 64)

		if err != nil {
			return -1, err
		} else {
			return i, nil
		}
	default:
		return -1, errors.New("Неизвестный аргумент")
	}
}

func binary(expr *ast.BinaryExpr) (ret float64, err error) {
	x, err1 := eval(expr.X)
	y, err2 := eval(expr.Y)
	ret = -1

	if (err1 == nil) && (err2 == nil) {

		switch expr.Op {
		case token.ADD:
			ret = x + y
		case token.SUB:
			ret = x - y
		case token.MUL:
			ret = x * y
		case token.QUO:
			ret = x / y
		case token.REM:
			ret = float64(int64(x) % int64(y))
		case token.AND:
			ret = float64(int64(x) & int64(y))
		case token.OR:
			ret = float64(int64(x) | int64(y))
		case token.XOR:
			ret = math.Pow(x, y)
		default:
			err = errors.New("Неизвестный бинарный оператор")
		}
	} else {
		if err1 != nil {
			err = err1
		} else {
			err = err2
		}
	}

	return
}
// Обработка констант
func ident(id *ast.Ident) (float64, error) {
	switch n := strings.ToLower(id.Name); n {
	case "pi":
		return math.Pi, nil
	case "e":
		return math.E, nil
	case "phi":
		return math.Phi, nil
	default:
		return -1, errors.New("Неизвестная константа " + n)
	}
}


// Обработка функциональных операторов с помощью созданного массива функций для обработки
func call(c *ast.CallExpr) (float64, error) {
	switch t := c.Fun.(type) {
	case *ast.Ident:
	default:
		_ = t
		return -1, errors.New("Неизвестный тип функции")
	}

	ident := c.Fun.(*ast.Ident)

	args := make([]float64, len(c.Args))
	for i, expr := range c.Args {
		var err error
		args[i], err = eval(expr)
		if err != nil {
			return -1, err
		}
	}

	name := strings.ToLower(ident.Name)

	if val, ok := funcMap[name]; ok {
		if len(args) == val.Args {
			return val.Func(args...), nil
		} else {
			return -1, errors.New("Слишком много аргументов для " + name)
		}
	} else {
		return -1, errors.New("Неизвестная функция " + name)
	}
}
