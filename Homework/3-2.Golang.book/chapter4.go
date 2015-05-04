package main

import "fmt"

func main() {
	var conver ftom = 36.5
	fmt.Println(ftToMeter(conver))

	fmt.Println(ftoC(conver))

}

func ftToMeter(feet ftom) ftom {
	return feet * .3048
}

func ftoC(fahra ftom) ftom {
	fahra = fahra - 32
	return fahra * .55555
}
