package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	const TimeFormatMilli = "2006-01-02 15:04:05.000"
	pc, _, _, _ := runtime.Caller(0)
	currentTime := time.Now()
	fmt.Println("currentTime", currentTime.Format(TimeFormatMilli), "start:", runtime.FuncForPC(pc).Name())

	_0 := 0_0
	_1 := 20
	π := 3
	a := "hello"
	数値１ := 33

	fmt.Println(_0)
	fmt.Println(_1)
	fmt.Println(π)
	fmt.Println(a)
	fmt.Println(数値１)

	currentTime = time.Now()
	fmt.Println("currentTime", currentTime.Format(TimeFormatMilli), "end:", runtime.FuncForPC(pc).Name())
}
