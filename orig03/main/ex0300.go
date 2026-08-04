package main

import (
	"log/slog"
	"runtime"
	"time"
)

// main
func main() {
	logger := slog.New(newHandler())
	pc, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(pc).Name()

	currentTime := time.Now()
	logger.Info(currentTime.Format(TimeFormatMilli)+" start", "func", fn)

	readFile(r0711world)

	currentTime = time.Now()
	logger.Info(currentTime.Format(TimeFormatMilli)+" end", "func", fn)
}
