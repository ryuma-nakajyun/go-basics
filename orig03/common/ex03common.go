package common

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"runtime"
	"time"
)

const TimeFormatMilli = "2006-01-02 15:04:05.000"
const r0711world = "r0711world_utf8.tsv"

// slog の自動 timestamp を消すハンドラ
func newHandler() slog.Handler {
	return slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// slog が自動で付ける timestamp を削除
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			// 日付を先頭に追加
			if a.Key == slog.LevelKey {
				now := time.Now().Format(TimeFormatMilli)
				return slog.String("ts", now)
			}
			return a
		},
		AddSource: true,
	})
}

func readFile(name string) {
	logger := slog.New(newHandler())
	pc, _, _, _ := runtime.Caller(1) // ← ここだけ変更
	fn := runtime.FuncForPC(pc).Name()

	currentTime := time.Now()
	logger.Info(currentTime.Format(TimeFormatMilli)+" start", "func", fn)

	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(line))
	}

	currentTime = time.Now()
	logger.Info(currentTime.Format(TimeFormatMilli)+" end", "func", fn)
}
