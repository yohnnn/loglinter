package main

import (
	"log/slog"

	"github.com/yohnnn/loglinter/analyzer"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	slog.Info("starting loglinter...") // Test log message.
	singlechecker.Main(analyzer.New())
}
