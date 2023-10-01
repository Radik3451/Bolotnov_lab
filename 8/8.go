package main

import (
	"fmt"
	"intervalLib"
)

func main() {
	fmt.Println("Если хотите прочитать инструкцию нажмите 1")
	var choice int
	fmt.Scan(&choice)

	if choice == 1 {
		str := intervalLib.Getinstruction()
		fmt.Print(str)
	} else {
		var a, b intervalLib.Interval
		fmt.Print("Введите границы первого интервала: ")
		fmt.Scan(&a.L, &a.R)
		fmt.Print("Введите границы второго интервала: ")
		fmt.Scan(&b.L, &b.R)

		fmt.Println()

		result := intervalLib.Multiplication(a, b)
		fmt.Printf("Умножение: [%.15f, %.15f]\n\n", result.L, result.R)

		result = intervalLib.Division(a, b)
		fmt.Printf("Деление: [%.15f, %.15f]\n\n", result.L, result.R)

		result = intervalLib.Addition(a, b)
		fmt.Printf("Сложение: [%.15f, %.15f]\n\n", result.L, result.R)

		result = intervalLib.Substraction(a, b)
		fmt.Printf("Вычитание: [%.15f, %.15f]\n\n", result.L, result.R)

		result = intervalLib.NonstandardDivision(a, b)
		fmt.Printf("Нестандартное деление: [%.15f, %.15f]\n\n", result.L, result.R)

		result = intervalLib.NonstandardSubtraction(a, b)
		fmt.Printf("Нестандартное вычитание: [%.15f, %.15f]\n\n", result.L, result.R)
	}
}
