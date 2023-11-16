package main

import (
	"fmt"
	"math"
	"sort"
)

type Interval struct {
	L, R float64
}

func RoundToDecimal(num Interval, decimals int) Interval {
	scale := math.Pow(15, float64(decimals))
	return Interval{L: math.Round(num.L*scale) / scale, R: math.Round(num.R*scale) / scale}
}

// Возвращает наибольшее значениее из 2-ух интервалов
func Comparison(X, Y Interval) Interval {
	if X.R < Y.L {
		return RoundToDecimal(Y, 15)
	} else if X.L < Y.R {
		return RoundToDecimal(X, 15)
	} else {
		panic("Интервалы равны")
	}
}

// Операция умножения интервалов
func Multiplication(X, Y Interval) Interval {
	tmp := make([]float64, 4)
	tmp[0] = X.L * Y.L
	tmp[1] = X.L * Y.R
	tmp[2] = X.R * Y.L
	tmp[3] = X.R * Y.R

	sort.Float64s(tmp)

	res := Interval{L: tmp[0], R: tmp[len(tmp)-1]}

	return RoundToDecimal(res, 15)
}

// Операция деления интервалов
func Division(X, Y Interval) Interval {
	tmp := make([]float64, 4)
	tmp[0] = X.L / Y.L
	tmp[1] = X.L / Y.R
	tmp[2] = X.R / Y.L
	tmp[3] = X.R / Y.R

	sort.Float64s(tmp)

	res := Interval{L: tmp[0], R: tmp[len(tmp)-1]}

	return RoundToDecimal(res, 15)
}

// Операция сложения интервалов
func Addition(X, Y Interval) Interval {
	tmp := make([]float64, 2)
	tmp[0] = X.L + Y.L
	tmp[1] = X.R + Y.R

	sort.Float64s(tmp)

	res := Interval{L: tmp[0], R: tmp[len(tmp)-1]}

	return RoundToDecimal(res, 15)
}

// Операция вычитания интервалов
func Substraction(X, Y Interval) Interval {
	tmp := make([]float64, 2)
	tmp[0] = X.L - Y.R
	tmp[1] = X.R - Y.L

	sort.Float64s(tmp)

	res := Interval{L: tmp[0], R: tmp[len(tmp)-1]}

	return RoundToDecimal(res, 15)
}

func Getinstruction() string {
	str := "Инструкция для работы с библиотекой\ncomparison(X, Y Interval) Interval - Возвращает набольшее значениее из 2-ух интервалов\nmultiplication(X, Y Interval) Interval - Возвращает интервал, являющийся произведением 2-ух интервалов\ndivision(X, Y Interval) Interval - Возвращает интервал, являющийся результатом деления 2-ух интервалов\naddition(X, Y Interval) Interval - Возвращает интервал, являющийся суммой 2-ух интервалов\nsubstraction(X, Y Interval) Interval - Возвращает интервал, являющийся разницей 2-ух интервалов\n}"
	return str
}

func test() {
	var a, b Interval
	a = Interval{L: 0.999999999, R: 0.999999999}
	b = Interval{L: 0.999999999, R: 0.999999999}
	count := 1

	for i := 1; i < 1000000000; i++ {
		A := Addition(a, b)
		if i == 200000000*count {
			fmt.Printf("[%d, %15.15f, %15.15f, %15.15f, %15.15f]\n", i, A.L, A.R, (A.L+A.R)/2, (A.L+A.R)/2-A.L)
			count++
		}

	}
}

func main() {
	// fmt.Println("Если хотите прочитать инструкцию нажмите 1")
	// var choice int
	// fmt.Scan(&choice)

	// if choice == 1 {
	// 	str := Getinstruction()
	// 	fmt.Print(str)
	// } else {
	// 	var a, b Interval
	// 	fmt.Print("Введите границы первого интервала: ")
	// 	fmt.Scan(&a.L, &a.R)
	// 	fmt.Print("Введите границы второго интервала: ")
	// 	fmt.Scan(&b.L, &b.R)

	// 	fmt.Println()

	// 	result := Multiplication(a, b)
	// 	fmt.Printf("Умножение: [%.15f, %.15f]\n\n", result.L, result.R)

	// 	result = Division(a, b)
	// 	fmt.Printf("Деление: [%.15f, %.15f]\n\n", result.L, result.R)

	// 	result = Addition(a, b)
	// 	fmt.Printf("Сложение: [%.15f, %.15f]\n\n", result.L, result.R)

	// 	result = Substraction(a, b)
	// 	fmt.Printf("Вычитание: [%.15f, %.15f]\n\n", result.L, result.R)

	// 	result = Division(a, b)
	// 	fmt.Printf("Нестандартное деление: [%.15f, %.15f]\n\n", result.L, result.R)

	// 	result = Substraction(a, b)
	// 	fmt.Printf("Нестандартное вычитание: [%.15f, %.15f]\n\n", result.L, result.R)

	// }
	test()
}
