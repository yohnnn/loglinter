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
		confMap, ok := conf.(map[string]any)
		if !ok {
			return nil, errors.New("conf must be a map[string]any")
		}

		if enabled, ok := confMap["enable_sensitive"].(bool); ok {
			cfg.EnableSensitive = enabled
		}

		if rawPatterns, ok := confMap["sensitive_patterns"]; ok {
			switch patterns := rawPatterns.(type) {
			case []string:
				cfg.SensitivePatterns = append(cfg.SensitivePatterns, patterns...)
			case []any:
				for _, p := range patterns {
					s, ok := p.(string)
					if !ok {
						continue
					}

					cfg.SensitivePatterns = append(cfg.SensitivePatterns, s)
				}
			}
		}
	}

	return []*analysis.Analyzer{analyzer.NewWithConfig(cfg)}, nil
}
