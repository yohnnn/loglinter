package main

import (
	"log/slog"

	"github.com/yohnnn/loglinter/analyzer"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	slog.Info("starting loglinter...") // This log is just for demonstration and will not be analyzed by the linter.
	singlechecker.Main(analyzer.New())
}
