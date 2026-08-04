package main

import (
	"log/slog"
	"runtime"
	"time"
	"github.com/ryuma-nakajyun/go-basics/orig03/common"
)

// main
func main() {
	logger := slog.New(common.newHandler())
	pc, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(pc).Name()

	currentTime := time.Now()
	logger.Info(currentTime.Format(common.TimeFormatMilli)+" start", "func", fn)

	common.readFile(common.R0711world)

	currentTime = time.Now()
	logger.Info(currentTime.Format(common.TimeFormatMilli)+" end", "func", fn)
}
