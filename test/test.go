package main

import (
	"fmt"
	"math"
)

func MakeArr(M, V int) ([]float64, []float64, []float64, []float64) {
	A := make([]float64, M)
	B := make([]float64, M)
	C := make([]float64, M)
	D := make([]float64, M)

	// Заполняем матрицу
	for i := 1; i < M; i++ {
		// Центральная диагональ
		B[i] = 10*float64(V) + float64(i)/float64(V)

		// Верхняя диагональ
		if i < M-1 {
			C[i] = 0.4 * math.Cos(float64(i)) / float64(V)
		} else {
			C[i] = 0
		}

		// Нижняя диагональ
		if i > 1 {
			A[i] = 0.3 * math.Sin(float64(i)) / float64(V)
		} else {
			A[i] = 0
		}

		D[i] = 1.3 + float64(i)/float64(V)
	}

	return C, B, A, D
}

func Solution(C, B, A, D []float64) []float64 {
	n := len(C)
	x := make([]float64, n)
	alpha := make([]float64, n)
	beta := make([]float64, n)
	y := B[1]

	alpha[1] = -C[1] / y
	beta[1] = D[1] / y

	for i := 2; i < n-1; i++ {
		y = B[i] + A[i]*alpha[i-1]
		alpha[i] = -C[i] / y
		beta[i] = (D[i] - A[i]*beta[i-1]) / y
	}
	y = B[n-1] + A[n-1]*alpha[n-2]
	beta[n-1] = (D[n-1] - A[n-1]*beta[n-2]) / y

	x[n-1] = beta[n-1]
	for i := n - 2; i > 0; i-- {
		x[i] = alpha[i]*x[i+1] + beta[i]
	}
	return x
}

func PrintVector(arr []float64, name string) {
	n := len(arr)
	fmt.Println("\n\t\t", name)
	for i := 1; i < n; i++ {
		fmt.Printf("[%.6e]  ", arr[i])
	}
}

func dot(v []float64, w []float64) float64 {
	sum := 0.0
	for i := 0; i < len(v); i++ {
		sum += v[i] * w[i]
	}
	return sum
}

func Residual(C, B, A, D, x []float64) ([]float64, float64) {
	n := len(B)
	// Вычисляем вектор невязки
	res := make([]float64, n)

	for i := 1; i < n; i++ {
		res[i] = B[i] * x[i]

		if i < n-1 {
			res[i] += C[i] * x[i+1]
		}

		if i > 1 {
			res[i] += A[i] * x[i-1]
		}
	}

	NevVector := make([]float64, n)
	for i := 1; i < n; i++ {
		NevVector[i] = D[i] - res[i]
	}

	// Вычисляем норму вектора невязки
	norm := 0.0
	for i := 1; i < len(NevVector); i++ {
		norm += math.Pow(NevVector[i], 2)
	}
	norm = math.Sqrt(norm)

	return res, norm
}

func main() {
	// Задаем параметры
	M := 8
	V := 3

	C, B, A, D := MakeArr(M+1, V)

	// Выводим матрицу
	PrintVector(C, "Вектор C")
	PrintVector(B, "Вектор B")
	PrintVector(A, "Вектор A")
	PrintVector(D, "Вектор D")
	x := Solution(C, B, A, D)
	PrintVector(x, "Вектор X")
	res, norm := Residual(C, B, A, D, x)
	PrintVector(res, "Вектор невязки")
	fmt.Println("\nВектор норма вектора нормали")
	fmt.Print(norm, "\n")
}
