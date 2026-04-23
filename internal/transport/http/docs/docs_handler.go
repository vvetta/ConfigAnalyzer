package swaggerdocs

import "net/http"

type Handler interface {
	ServeSpec(w http.ResponseWriter, r *http.Request)
	ServeUI(w http.ResponseWriter, r *http.Request)
	RedirectToUI(w http.ResponseWriter, r *http.Request)
}

type DocsHandler struct {
	spec []byte
}

func NewDocsHandler() *DocsHandler {
	return &DocsHandler{}
}

func (h *DocsHandler) ServeSpec(w http.ResponseWriter, r *http.Request)
func (h *DocsHandler) ServeUI(w http.ResponseWriter, r *http.Request)
func (h *DocsHandler) RedirectToUI(w http.ResponseWriter, r *http.Request)
