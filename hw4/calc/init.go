package calc

import "math"

// Создание массива функций для обработки данных
func init() {
	funcMap = make(map[string]Func)
	funcMap["sqrt"] = Func{
		Name: "sqrt",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Sqrt(args[0])
		},
	}
	funcMap["abs"] = Func{
		Name: "abs",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Abs(args[0])
		},
	}
	funcMap["log"] = Func{
		Name: "log",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Log(args[0]) / math.Log(args[1])
		},
	}
	funcMap["ln"] = Func{
		Name: "ln",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Log(args[0])
		},
	}
	funcMap["sin"] = Func{
		Name: "sin",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Sin(args[0])
		},
	}
	funcMap["cos"] = Func{
		Name: "cos",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Cos(args[0])
		},
	}
	funcMap["tan"] = Func{
		Name: "tan",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Tan(args[0])
		},
	}
	funcMap["arcsin"] = Func{
		Name: "arcsin",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Asin(args[0])
		},
	}
	funcMap["arccos"] = Func{
		Name: "arccos",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Acos(args[0])
		},
	}
	funcMap["arctan"] = Func{
		Name: "arctan",
		Args: 1,
		Func: func(args ...float64) float64 {
			return math.Atan(args[0])
		},
	}
	funcMap["max"] = Func{
		Name: "max",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Max(args[0], args[1])
		},
	}
	funcMap["min"] = Func{
		Name: "min",
		Args: 2,
		Func: func(args ...float64) float64 {
			return math.Min(args[0], args[1])
		},
	}
}
