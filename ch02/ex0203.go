package main

import (
	"fmt"
	"runtime"
	"time"
)

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	pc, _, _, _ := runtime.Caller(0)
	currentTime := time.Now()
	fmt.Println("currentTime", currentTime.Format(TimeFormatMilli), "start:", runtime.FuncForPC(pc).Name())

	const y = "hello"
	fmt.Println(x)
	fmt.Println(y)

	x = x + 1
	y = "bye"
	fmt.Println(x)
	fmt.Println(y)

	currentTime = time.Now()
	fmt.Println("currentTime", currentTime.Format(TimeFormatMilli), "end:", runtime.FuncForPC(pc).Name())
}
