package main

import (
	"fmt"
	"runtime"
)

func main() {
	pc, _, _, _ := runtime.Caller(0);
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println("start:",  funcName)
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
	fmt.Println("end:", funcName)
}
