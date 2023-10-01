// package intervalLib

// import (
// 	"sort"
//  "math"
// )

// type Interval struct {
// 	L, R float64
// }

// func RoundToDecimal(num Interval, decimals int) Interval {
// 	scale := math.Pow(10, float64(decimals))
// 	return Interval{L: math.Round(num.L*scale) / scale, R: math.Round(num.R*scale) / scale}
// }

// // Возвращает наибольшее значениее из 2-ух интервалов
// func Comparison(X, Y Interval) Interval {
// 	if X.R < Y.L {
// 		return RoundToDecimal(Y, 15)
// 	} else if X.L < Y.R {
// 		return RoundToDecimal(X, 15)
// 	} else {
// 		panic("Интервалы равны")
// 	}
// }

// // Операция умножения интервалов
// func Multiplication(X, Y Interval) Interval {
// 	tmp := make([]float64, 4)
// 	tmp[0] = X.L * Y.L
// 	tmp[1] = X.L * Y.R
// 	tmp[2] = X.R * Y.L
// 	tmp[3] = X.R * Y.R

// 	sort.Float64s(tmp)

// 	res := Interval{L: tmp[0], R: tmp[len(tmp)-1]}

// 	return RoundToDecimal(res, 15)
// }

// // Операция деления интервалов
// func Division(X, Y Interval) Interval {
// 	tmp := make([]float64, 4)
// 	tmp[0] = X.L / Y.L
// 	tmp[1] = X.L / Y.R
// 	tmp[2] = X.R / Y.L
// 	tmp[3] = X.R / Y.R

// 	sort.Float64s(tmp)

// 	res := Interval{L: tmp[0], R: tmp[len(tmp)-1]}

// 	return RoundToDecimal(res, 15)
// }

// // Операция сложения интервалов
// func Addition(X, Y Interval) Interval {
// 	tmp := make([]float64, 2)
// 	tmp[0] = X.L + Y.L
// 	tmp[1] = X.R + Y.R

// 	sort.Float64s(tmp)

// 	res := Interval{L: tmp[0], R: tmp[len(tmp)-1]}

// 	return RoundToDecimal(res, 15)
// }

// // Операция вычитания интервалов
// func Substraction(X, Y Interval) Interval {
// 	tmp := make([]float64, 2)
// 	tmp[0] = X.L - Y.R
// 	tmp[1] = X.R - Y.L

// 	sort.Float64s(tmp)

// 	res := Interval{L: tmp[0], R: tmp[len(tmp)-1]}

// 	return RoundToDecimal(res, 15)
// }

// // Операция нестандартного вычитания интервалов
// func NonstandardSubtraction(X, Y Interval) Interval {
// tmp := make([]float64, 2)
// tmp[0] = X.L - Y.L
// tmp[1] = X.R - Y.R

// sort.Float64s(tmp)

// res := Interval{L: tmp[0], R: tmp[len(tmp)-1]}

// return res, 15
// }

// // Операция нестандартного деления интервалов
// func NonstandardDivision(X, Y Interval) Interval {
// 	if Comparison(Multiplication(X, Y), Interval{L: 0, R: 0}) == Multiplication(X, Y) {
// 		tmp := make([]float64, 2)
// 		tmp[0] = X.L / Y.L
// 		tmp[1] = X.R / Y.R

// 		sort.Float64s(tmp)
// 		res := Interval{L: tmp[0], R: tmp[1]}

// 		return res
// 	} else if Comparison(Multiplication(X, Y), Interval{L: 0, R: 0}) != Multiplication(X, Y) {
// 		tmp := make([]float64, 2)
// 		tmp[0] = X.L / Y.R
// 		tmp[1] = X.R / Y.L

// 		sort.Float64s(tmp)
// 		res := Interval{L: tmp[0], R: tmp[1]}

// 		return res
// 	} else if X.L <= 0 && X.R >= 0 && Y.L > 0 && Y.R > 0 {
// 		tmp := 1 / Y.L

// 		return Interval{L: tmp * X.L, R: tmp * X.R}
// 	} else if X.L <= 0 && X.R >= 0 && Y.L < 0 && Y.R < 0 {
// 		tmp := 1 / Y.R

// 		return Interval{L: tmp * X.L, R: tmp * X.R}
// 	} else {
// 		panic("Данные вычисления не предусмотренны библиотекой")
// 	}
// }

// func Getinstruction() string {
// 	str := "Инструкция для работы с библиотекой\ncomparison(X, Y Interval) Interval - Возвращает набольшее значениее из 2-ух интервалов\nmultiplication(X, Y Interval) Interval - Возвращает интервал, являющийся произведением 2-ух интервалов\ndivision(X, Y Interval) Interval - Возвращает интервал, являющийся результатом деления 2-ух интервалов\naddition(X, Y Interval) Interval - Возвращает интервал, являющийся суммой 2-ух интервалов\nsubstraction(X, Y Interval) Interval - Возвращает интервал, являющийся разницей 2-ух интервалов\nnonstandardSubtraction(X, Y Interval) Interval - Выполняет нестандартное делениt 2-ух интервалов\nnonstandardDivision(X, Y Interval) Interval - Выполняет нестандартное вычитание 2-ух интервалов\n}"
// 	return str
// }