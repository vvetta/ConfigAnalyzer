package handlers

import (
	"ConfigAnalyzer/internal/usecase"
	"net/http"
)

type Handler interface {
	CheckJSON(w http.ResponseWriter, r *http.Request)
	CheckYAML(w http.ResponseWriter, r *http.Request)
}

type ConfigAnalyzerHandler struct {
	usecase usecase.Analyzer
}

func NewConfigAnalyzerHandler(usecase usecase.Analyzer) *ConfigAnalyzerHandler {
	return &ConfigAnalyzerHandler{usecase: usecase}
}

func (h *ConfigAnalyzerHandler) CheckJSON(w http.ResponseWriter, r *http.Request) {}
func (h *ConfigAnalyzerHandler) CheckYAML(w http.ResponseWriter, r *http.Request) {}
