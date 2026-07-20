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

	ａ := "こんにちは"
	a := "サヨウナラ"
	fmt.Println(ａ)
	fmt.Println(a)

	currentTime = time.Now()
	fmt.Println("currentTime", currentTime.Format(TimeFormatMilli), "end:", runtime.FuncForPC(pc).Name())
}
