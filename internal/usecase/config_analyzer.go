package usecase

type Analyzer interface {
	CheckJSON()
	CheckYAML()
}

type ConfigAnalyzer struct{}

func NewConfigAnalyzer() *ConfigAnalyzer {
	return &ConfigAnalyzer{}
}

func (c *ConfigAnalyzer) CheckJSON() {}
func (c *ConfigAnalyzer) CheckYAML() {}
