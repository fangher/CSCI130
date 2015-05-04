package main

func halfCheck(myHalfBool float64) bool {
	myHalfBool = myHalfBool / 2
	boolValue := false
	if math.Mod(myHalfBool, 2) == 0 {
		boolValue = true
	}
	return boolValue
}

func gList(myArgList ...int) int {
	greatest := 0
	for _, v := range myArgList {
		if greatest < v {
			greatest = v
		}
	}
	return greatest
}

func MOG() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func Recfib(fibInt int) int {
	if fibInt == 0 || fibInt == 1 {
		return fibInt
	} else {
		return (Recfib(fibInt-1) + Recfib(fibInt-2))
	}
}


