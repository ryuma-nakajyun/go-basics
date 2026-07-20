package main

import (
	"fmt"
	"log/slog"
	"runtime"
	"time"
)

func main() {
	pc, _, _, _ := runtime.Caller(0)
	currentTime := time.Now()
	slog.Info("start", "time", currentTime.Format(TimeFormatMilli), "func", runtime.FuncForPC(pc).Name())

	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)
	fmt.Println(z, d)

	currentTime = time.Now()
	slog.Info("end", "time", currentTime.Format(TimeFormatMilli), "func", runtime.FuncForPC(pc).Name())
}
