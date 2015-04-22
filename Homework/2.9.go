package main

import "fmt"

type Pair struct {
	id   string
	data int
}

func PrintPair(p Pair) {
	fmt.Println(p.id, p.data)
}

func fibRec(n int) int {
	if n < 2 {
		return 1
	}
	return (fibRec(n-1) + fibRec(n-2))
}

func helloyou(x string) (a string, b string) {
	return "Hello " + x, "you!"
}

const (
	Name string = "Fang"
)

func main() {
	p1 := Pair{"pair 1", 2}
	PrintPair(p1)

	fmt.Println("4th Fib", fibRec(4))

	print_A, print_B := helloyou(Name)
	fmt.Println(print_A, ",", print_B)

	print_A, _ = helloWorldForYou("Steve")
	fmt.Println(print_A, ",", "Nothing")
}
