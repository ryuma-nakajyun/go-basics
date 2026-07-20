package main

import (
	"fmt"
	"runtime"
	"time"
)

// 定数宣言
const TimeFormatMilli = "2006-01-02 15:04:05.000"

func main() {
	pc, _, _, _ := runtime.Caller(0)
	currentTime := time.Now()
	fmt.Println("currentTime", currentTime.Format(TimeFormatMilli), "start:", runtime.FuncForPC(pc).Name())

	{
		var data []int
		fmt.Println(data == nil)
	}

	{
		x := []int{} //
		fmt.Println(x == nil)
	}

	currentTime = time.Now()
	fmt.Println("currentTime", currentTime.Format(TimeFormatMilli), "end:", runtime.FuncForPC(pc).Name())
}
