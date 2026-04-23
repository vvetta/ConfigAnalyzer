package usecase

import "ConfigAnalyzer/internal/domain"

type Analyzer interface {
	Analyze(req AnalyzeRequest) ([]domain.Issue, error)
}

type ConfigAnalyzer struct{}

func NewConfigAnalyzer() *ConfigAnalyzer {
	return &ConfigAnalyzer{}
}

type AnalyzeRequest struct {
	Data   []byte
	Format string
}

func (c *ConfigAnalyzer) Analyze(req AnalyzeRequest) ([]domain.Issue, error) {
	return nil, nil
}
