package main

import (
	"fmt"
	"runtime"
	"time"
)

// 定数宣言
//const TimeFormatMilli = "2006-01-02 15:04:05.000"

func main() {
	pc, _, _, _ := runtime.Caller(0)
	currentTime := time.Now()
	fmt.Println("currentTime", currentTime.Format(TimeFormatMilli), "start:", runtime.FuncForPC(pc).Name())

	data := []int{2, 4, 6, 8}
	fmt.Println(data)

	currentTime = time.Now()
	fmt.Println("currentTime", currentTime.Format(TimeFormatMilli), "end:", runtime.FuncForPC(pc).Name())
}
