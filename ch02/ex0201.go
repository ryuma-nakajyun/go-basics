package main

import (
	"fmt"
	"math/cmplx"
	"runtime"
	"time"
)

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

	currentTime = time.Now()
	fmt.Println("currentTime", currentTime.Format(TimeFormatMilli), "end:", runtime.FuncForPC(pc).Name())
}
