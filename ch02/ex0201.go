package main

import (
	"fmt"
	"math/cmplx"
	"runtime"
	"time"
)

// 定数宣言
const TimeFormatMilli = "2006-01-02 15:04:05.000"

func main() {
	pc, _, _, _ := runtime.Caller(0)
	currentTime := time.Now()
	fmt.Println("currentTime", currentTime.Format(TimeFormatMilli), "start:", runtime.FuncForPC(pc).Name())

	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))

	pc, _, _, _ = runtime.Caller(0)
	currentTime = time.Now()
	fmt.Println("currentTime", currentTime.Format(TimeFormatMilli), "end:", runtime.FuncForPC(pc).Name())

}
