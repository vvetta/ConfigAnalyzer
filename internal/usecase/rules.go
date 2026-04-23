package usecase

import "ConfigAnalyzer/internal/domain"

type DebugLoggingRule struct{}
type PlaintextPasswordRule struct{}
type OpenBindRule struct{}
type TLSDisabledRule struct{}
type WeakAlgorithmRule struct{}

func (DebugLoggingRule) Check(entry Entry) []domain.Issue
func (PlaintextPasswordRule) Check(entry Entry) []domain.Issue
func (OpenBindRule) Check(entry Entry) []domain.Issue
func (TLSDisabledRule) Check(entry Entry) []domain.Issue
func (WeakAlgorithmRule) Check(entry Entry) []domain.Issue
