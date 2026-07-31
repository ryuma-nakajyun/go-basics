package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

const TimeFormatMilli = "2006-01-02 15:04:05.000"
const r0711world = "r0711world_utf8.tsv"
const Hemisphere = 0.0

// slog の自動 timestamp を消すハンドラ
func newHandler() slog.Handler {
	return slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// slog が自動で付ける timestamp を削除
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
		AddSource: true, // file:line を出す
	})
}

// readFile
func readFile(name string) {
	logger := slog.New(newHandler())
	pc, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(pc).Name()

	currentTime := time.Now()
	logger.Info(fmt.Sprintf("%s start", currentTime.Format(TimeFormatMilli)), "func", fn)

	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)

	// 1. ヘッダー行を読む
	headerLine, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	headers := strings.Split(string(headerLine), "\t")

	// 2. 列名 → インデックスの map を作る
	colIndex := make(map[string]int)
	for i, h := range headers {
		colIndex[h] = i
	}

	// 3. データ行を読む
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fields := strings.Split(string(line), "\t")
		lotStr := fields[colIndex["capital_longitude"]]
		lot, err := strconv.ParseFloat(lotStr, 64)
		if err != nil {
			log.Fatalf("invalid lotitude: %v", err)
		}

		// 南半球にあるものだけ抽出
		if lot > Hemisphere {
			continue
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0) // カラムの幅を整える
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%.8f\n",
			fields[colIndex["iso_country_code"]],
			fields[colIndex["country_name_ja_common"]],
			fields[colIndex["capital_name_ja"]],
			fields[colIndex["capital_latitude"]],
			lot,
		)
		w.Flush()
	}

	currentTime = time.Now()
	logger.Info(fmt.Sprintf("%s end", currentTime.Format(TimeFormatMilli)), "func", fn)
}

// main
func main() {
	logger := slog.New(newHandler())
	pc, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(pc).Name()

	currentTime := time.Now()
	logger.Info(fmt.Sprintf("%s start", currentTime.Format(TimeFormatMilli)), "func", fn)

	readFile(r0711world)

	currentTime = time.Now()
	logger.Info(fmt.Sprintf("%s end", currentTime.Format(TimeFormatMilli)), "func", fn)
}
