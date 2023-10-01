package main

import (
	"fmt"
	"math"
	"sort"
)

const V = 3
const rad = 0.001

type Interval struct {
	l, r float64
}

func (i Interval) Add(other Interval) Interval {
	return Interval{i.l + other.l, i.r + other.r}
}

func makeArr(N int) []Interval {
	arr := make([]Interval, N)
	for i := 1; i < N; i++ {
		arr[i].l = math.Sin(float64(V+i)) - rad
		arr[i].r = math.Sin(float64(V+i)) + rad
	}
	return arr
}

func displayIntervals(arr []Interval) {
	fmt.Println("\nИнтервалы:")
	for _, interval := range arr {
		fmt.Printf("[%0.5f, %0.5f]\n", interval.l, interval.r)
	}
	fmt.Println()
}

func calculateSum(arr []Interval) Interval {
	var sum Interval
	for _, interval := range arr {
		sum = sum.Add(interval)
	}
	return sum
}

func main() {
	var N int
	fmt.Print("Введите N: ")
	fmt.Scan(&N)

	arr := makeArr(N)

	// Вычисляем сумму до сортировки
	sumBeforeSorting := calculateSum(arr)

	fmt.Println("Исходный массив")
	displayIntervals(arr)

	// Сортируем интервалы по полю l (по левому краю каждого интервала)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].l < arr[j].l
	})

	fmt.Println("Сортированные интервалы по полю l (по левому краю каждого интервала):")
	displayIntervals(arr)

	// Вычисляем сумму после сортировки
	sumAfterSorting := calculateSum(arr)

	fmt.Println("Сумма интервалов до сортировки: \t", sumBeforeSorting)
	fmt.Println("Сумма интервалов после сортировки: \t", sumAfterSorting)

	// Нет необходимости освобождать память в Go, сборка мусора заботится об этом
}
