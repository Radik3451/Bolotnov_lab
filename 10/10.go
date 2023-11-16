package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type Interval struct {
	L, R float64
}

func NegativeInterval(X Interval) Interval {
	res := Division(X, Interval{L: -1, R: -1})
	return res
}

func Multiplication(X, Y Interval) Interval {
	tmp := make([]float64, 4)
	tmp[0] = X.L * Y.L
	tmp[1] = X.L * Y.R
	tmp[2] = X.R * Y.L
	tmp[3] = X.R * Y.R

	sort.Float64s(tmp)

	res := Interval{L: tmp[0], R: tmp[len(tmp)-1]}

	return res
}

func Division(X, Y Interval) Interval {
	tmp := make([]float64, 4)
	tmp[0] = X.L / Y.L
	tmp[1] = X.L / Y.R
	tmp[2] = X.R / Y.L
	tmp[3] = X.R / Y.R

	sort.Float64s(tmp)

	res := Interval{L: tmp[0], R: tmp[len(tmp)-1]}

	return res
}

func Addition(X, Y Interval) Interval {
	tmp := make([]float64, 2)
	tmp[0] = X.L + Y.L
	tmp[1] = X.R + Y.R

	sort.Float64s(tmp)

	res := Interval{L: tmp[0], R: tmp[len(tmp)-1]}

	return res
}

func Substraction(X, Y Interval) Interval {
	tmp := make([]float64, 2)
	tmp[0] = X.L - Y.R
	tmp[1] = X.R - Y.L

	sort.Float64s(tmp)

	res := Interval{L: tmp[0], R: tmp[len(tmp)-1]}

	return res
}

func CreateMatrix(rad float64, V int, N int) ([][]Interval, []Interval) {
	B := make([]Interval, N)
	A := make([][]Interval, N)
	var tmp float64

	for i := 1; i < N; i++ {
		A[i] = make([]Interval, N)
		tmp = 10.0 * math.Cos(float64(i)+float64(V))
		B[i] = Interval{L: tmp - rad, R: tmp + rad}
		for j := 1; j < N; j++ {
			if i == j {
				tmp = 31.0 + math.Sin(float64(i))/float64(V)
				A[i][i] = Interval{L: tmp - rad, R: tmp + rad}
			} else {
				tmp = 0.01*float64(V) + math.Sin(float64(i)-float64(j))
				A[i][j] = Interval{L: tmp - rad, R: tmp + rad}
			}
		}

	}

	return A, B
}

func PrintArr(arr [][]Interval, name string, file *os.File) {
	n := len(arr)
	fmt.Println("\n\t\t", name)
	fmt.Fprintln(file, fmt.Sprintf("\n\t\t%s", name))
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			fmt.Printf("[%15.7f, %15.7f] ", arr[i][j].L, arr[i][j].R)
			fmt.Fprint(file, fmt.Sprintf("[%15.7f, %15.7f] ", arr[i][j].L, arr[i][j].R))
		}
		fmt.Println()
		fmt.Fprintln(file)
	}
}

func PrintTriangleArr(arr [][]Interval, name string, file *os.File) {
	n := len(arr)
	fmt.Println("\n\t\t", name)
	fmt.Fprintln(file, fmt.Sprintf("\n\t\t%s", name))
	for i := 1; i < n; i++ {
		for j := 1; j < n+1; j++ {
			fmt.Printf("[%15.7f, %15.7f] ", arr[i][j].L, arr[i][j].R)
			fmt.Fprint(file, fmt.Sprintf("[%15.7f, %15.7f] ", arr[i][j].L, arr[i][j].R))
		}
		fmt.Println()
		fmt.Fprintln(file)
	}
}

func PrintVector(arr []Interval, name string, file *os.File) {
	n := len(arr)
	fmt.Println("\n\t\t", name)
	fmt.Fprintln(file, fmt.Sprintf("\n\t\t%s", name))
	for i := 1; i < n; i++ {
		fmt.Printf("[%15.3f, %15.3f] ", arr[i].L, arr[i].R)
		fmt.Println()
		fmt.Fprintln(file, fmt.Sprintf("[%15.3f, %15.3f] ", arr[i].L, arr[i].R))
	}
}

func PrintVectorX(arr []Interval, name string, file *os.File) {
	n := len(arr)
	fmt.Println("\n\t\t", name)
	fmt.Fprintln(file, fmt.Sprintf("\n\t\t%s", name))
	for i := 1; i < n; i++ {
		fmt.Printf("[%15.10f, %15.10f] ", arr[i].L, arr[i].R)
		fmt.Println()
		fmt.Fprintln(file, fmt.Sprintf("[%15.10f, %15.10f] ", arr[i].L, arr[i].R))
	}
}

func Gaus(A [][]Interval, B []Interval) ([]Interval, [][]Interval) {
	n := len(B)

	triangle := make([][]Interval, n)
	for i := 1; i < n; i++ {
		triangle[i] = make([]Interval, n+1)
		for j := 1; j < n; j++ {
			triangle[i][j] = A[i][j]
		}
		triangle[i][n] = B[i]
	}

	for i := 1; i < n; i++ {
		// Проверка и приведение элемента triangle[i][i] к 1 (если он не ноль).
		if triangle[i][i].L != 0 && triangle[i][i].R != 0 {
			pivot := triangle[i][i]
			for j := i; j <= n; j++ {
				triangle[i][j] = Division(triangle[i][j], pivot)
			}
		}

		// Вычитание строки i из остальных строк.
		for j := i + 1; j < n; j++ {
			factor := triangle[j][i]
			for k := i; k <= n; k++ {
				triangle[j][k] = Substraction(triangle[j][k], Multiplication(factor, triangle[i][k]))
			}
		}
	}

	// Обратный ход
	solution := make([]Interval, n)
	for i := n - 1; i >= 1; i-- {
		solution[i] = triangle[i][n]
		for j := i + 1; j < n; j++ {
			solution[i] = Substraction(solution[i], Multiplication(triangle[i][j], solution[j]))
		}
	}

	return solution, triangle
}

func MatrixVectorMultiply(matrix [][]Interval, vector []Interval) []Interval {
	numRows := len(matrix)
	numCols := len(matrix[1])
	result := make([]Interval, numRows)

	for i := 1; i < numRows; i++ {
		result[i] = Interval{}
		for j := 1; j < numCols; j++ {
			term := Multiplication(matrix[i][j], vector[j])
			result[i] = Addition(result[i], term)
		}
	}

	return result
}

func checkSolution(A [][]Interval, B []Interval, X []Interval) ([]Interval, Interval) {
	n := len(A)

	res := make([]Interval, n)
	res = MatrixVectorMultiply(A, X)

	nev := make([]Interval, n)
	for i := 1; i < n; i++ {
		nev[i] = Substraction(B[i], res[i])
	}

	norm := Interval{}
	for i := 1; i < n; i++ {
		norm = Addition(norm, Interval{L: math.Pow(nev[i].L, 2), R: math.Pow(nev[i].R, 2)})
	}

	var normRes Interval
	normRes = Interval{L: math.Sqrt(norm.L), R: math.Sqrt(norm.R)}

	return nev, normRes
}

func main() {
	V := 3
	rad := 0.01
	N := 5

	file, err := os.Create("data.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	A, B := CreateMatrix(rad, V, N+1)
	PrintArr(A, "Вектор А", file)
	PrintVector(B, "Вектор B", file)
	X, triangle := Gaus(A, B)
	PrintVectorX(X, "Вектор X", file)
	PrintTriangleArr(triangle, "Матрица в треугольном виде", file)
	res, norm := checkSolution(A, B, X)
	PrintVector(res, "Вектор невязки", file)
	fmt.Print("\n\tНорма вектора невязки\n")
	fmt.Printf("[%15.6e, %15.6e] \n", norm.L, norm.R)
	fmt.Fprintln(file, fmt.Sprintf("\n\tНорма вектора невязки"))
	fmt.Fprintln(file, fmt.Sprintf("[%15.6e, %15.6e] \n", norm.L, norm.R))

}
