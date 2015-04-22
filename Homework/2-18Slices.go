package main

import "fmt"

func main() {

	fmt.Println("Numbers from 1 - 100:")
  	i := 1
  	for {
  		if i > 100 {
  			break
  		}
  		if i%2 == 1 {
  			i++
  			continue
  		}
  		fmt.Print(i, " ")
  		i++
  	}
	fmt.Println()

	sliceNames := []string{"Her", "Him", "Who", "Them"}

	someYear := map[string]int{
		"Nick": 2015,
		"Dan":  2016,
		"Kim":  2012,
		"Mike": 2009,
	}

	for _, item := range sliceNames {
		fmt.Println(item, "Year is", someYear[item])
	}

}
