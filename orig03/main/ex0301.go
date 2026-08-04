package main

import (
	"fmt"
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
	logger.Info(fmt.Sprintf("%s start", currentTime.Format(common.TimeFormatMilli)), "func", fn)

	common.ReadFile(common.R0711world)

	currentTime = time.Now()
	logger.Info(fmt.Sprintf("%s end", currentTime.Format(common.TimeFormatMilli)), "func", fn)
}
