package main

import (
	"errors"

	"github.com/yohnnn/loglinter/analyzer"
	"golang.org/x/tools/go/analysis"
)

// golangci-linter plugin (.so lib that exports New function)
func New(conf any) ([]*analysis.Analyzer, error) {
	cfg := analyzer.DefaultConfig()

	if conf != nil {
		confMap, ok := conf.(map[string]interface{})
		if !ok {
			return nil, errors.New("conf must be a map[string]interface{}")
		}

		if enabled, ok := confMap["enable_sensitive"].(bool); ok {
			cfg.EnableSensitive = enabled
		}
	}

	return []*analysis.Analyzer{analyzer.NewWithConfig(cfg)}, nil
}
