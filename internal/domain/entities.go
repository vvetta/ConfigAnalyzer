package domain

type Severity string

const (
	SeverityLow    Severity = "LOW"
	SeverityMedium Severity = "MEDIUM"
	SeverityHigh   Severity = "HIGH"
)

type Issue struct {
	Severity       Severity
	Rule           string
	Path           string
	Message        string
	Recommendation string
}
