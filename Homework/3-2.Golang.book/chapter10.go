package main

func SleepProcess(timeInMs time.Duration) {
	<-time.After(timeInMs * time.Second / 1000)
}

func minMaxFuncUsingPackage() {
	myIntArray := []float64{1, 2, -1, 20, 3, 4}
	avg := math.Average(myIntArray)
	fmt.Println(avg)
	fmt.Println(math.Min(myIntArray))
	fmt.Println(math.Max(myIntArray))
}
