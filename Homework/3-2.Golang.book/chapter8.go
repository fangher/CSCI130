package main

func swapVal(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
