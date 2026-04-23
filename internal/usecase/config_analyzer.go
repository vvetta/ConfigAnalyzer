package usecase

import "ConfigAnalyzer/internal/domain"

type Analyzer interface {
	Analyze(req AnalyzeRequest) ([]domain.Issue, error)
}

type Rule interface {
	Check(entry Entry) []domain.Issue
}

type Entry struct {
	Path      []string
	Key       string
	Value     any
	Parent    map[string]any
	PathText  string
	KeyLower  string
	ValueText string
}

type ConfigAnalyzer struct {
	rules []Rule
}

func NewConfigAnalyzer() *ConfigAnalyzer {
	return &ConfigAnalyzer{
		rules: []Rule{
			DebugLoggingRule{},
			PlaintextPasswordRule{},
			OpenBindRule{},
			TLSDisabledRule{},
			WeakAlgorithmRule{},
		},
	}
}

type AnalyzeRequest struct {
	Data   []byte
	Format string
}

func (c *ConfigAnalyzer) Analyze(req AnalyzeRequest) ([]domain.Issue, error) {
	return nil, nil
}
