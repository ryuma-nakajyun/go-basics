package main

import (
	"fmt"
	"log/slog"
	"runtime"
	"time"
)

func main() {
	const TimeFormatMilli = "2006-01-02 15:04:05.000"
	pc, _, _, _ := runtime.Caller(0)
	currentTime := time.Now()
	slog.Info("start", "time", currentTime.Format(TimeFormatMilli), "func", runtime.FuncForPC(pc).Name())

	x := 10
	x = 20
	fmt.Println(x)
	x = 30

	currentTime = time.Now()
	slog.Info("end", "time", currentTime.Format(TimeFormatMilli), "func", runtime.FuncForPC(pc).Name())

}
