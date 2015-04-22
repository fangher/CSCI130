package main

import (
	"fmt"
)

type my_operator func(int, int) int

func my_operation(num1 int, num2 int, op my_operator) int {
	  result := op(num1, num2)
  	if result < 0 {
		  fmt.Println("Negative value")
	  } else {
		  fmt.Println("Positive value")
	  }
	  return result
}

func operator_creation() (my_operator, my_operator, my_operator, my_operator) {
  	return func(a int, b int) int {
  			return a + b
  		}, func(a int, b int) int {
  			return a - b
  		}, func(a int, b int) int {
  			return a * b
  		}, func(a int, b int) int {
  			return a / b
  		}
}

func main() {
    	add, sub, mult, div := operator_creation()
    	a, b := 10, 6
    	var result int
    	operation := "multiply"
    	if a < b {
    		operation = "subtract"
    	}
    	switch operation {
    	case "add":
    		result = my_operation(a, b, add)
    	case "sub":
    		result = my_operation(a, b, subtract)
    	case "mult":
    		result = my_operation(a, b, multiply)
    	case "div":
    		result = my_operation(a, b, divide)
    	}
    	fmt.Println("The result is:", result)
}
