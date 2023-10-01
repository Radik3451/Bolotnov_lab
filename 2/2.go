package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return 3 + 4*x - 4*math.Pow(x, 2) + 2*math.Pow(x, 3)
}

func df(x float64) float64 {
	return 4 - 8*x + 6*math.Pow(x, 2)
}

func signChange(x1, x2 float64) bool {
	return df(x1)*df(x2) < 0
}

func bisection(x1, x2, maxIterations, epsilon float64) float64 {
	if !signChange(x1, x2) {
		// fmt.Println("Не удовлетворяет условиям метода бисекции.")
		return math.NaN()
	}

	var mid float64
	for i := 0.0; i < maxIterations; i++ {
		mid = (x1 + x2) / 2.0
		if math.Abs(df(mid)) < epsilon {
			return mid
		}

		if signChange(x1, mid) {
			x2 = mid
		} else {
			x1 = mid
		}
	}

	return mid
}

func solution(x1, x2, maxIterations, epsilon float64) {
	h := 0.01
	var points []float64
	for i := x1; i < x2; i += h {
		border1 := i
		border2 := i + h
		if signChange(border1, border2) {
			points = append(points, bisection(border1, border2, maxIterations, epsilon))
		}
	}

	points = append(points, x1, x2)
	var solutionY []float64
	for _, point := range points {
		solutionY = append(solutionY, f(point))
	}

	minVal := math.Inf(1)
	maxVal := math.Inf(-1)
	for _, y := range solutionY {
		if y < minVal {
			minVal = y
		}
		if y > maxVal {
			maxVal = y
		}
	}

	fmt.Printf("Ответ: [%.16f, %.16f]\n", minVal, maxVal)
}

func main() {
	var x1, x2 float64
	fmt.Print("Введите диапазон [x1, x2]: ")
	fmt.Scan(&x1, &x2)

	maxIterations := 20000.0
	epsilon := 1e-6

	solution(x1, x2, maxIterations, epsilon)
}
