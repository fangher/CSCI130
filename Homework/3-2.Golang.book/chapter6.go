package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	sort.Ints(x)
	fmt.Println(x[0])

	fmt.Println(isEven(10))
	fmt.Println(isEven(7))

	fmt.Println(max(4, 5, 23, 4, 5, 64, 32, 1, 3))

	fmt.Println(giveNextOdd(11))
	fmt.Println(fibo(10, 0, 1))

	var a, b int = 5, 2
	swap(&a, &b)
	fmt.Println(a, b)
}

func isEven(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}

func max(m ...int) int {
	var max int = 0
	for _, num := range m {
		if num > max {
			max = num
		}
	}
	return max

}

func giveNextOdd(x int) int {
	if isEven(x) {
		x += 1
		return x
	}
	x += 2
	return x
}

func fibo(n, a, b int) int {
	if n == 0 {
		return a
	}
	temp := b
	b = a + b
	a = temp
	return fibo(n-1, a, b)
}


func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
