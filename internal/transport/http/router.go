package transporthttp

import (
	docsHandler "ConfigAnalyzer/internal/transport/http/docs"
	analyzerHandler "ConfigAnalyzer/internal/transport/http/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(
	analyzer analyzerHandler.Handler,
	docs docsHandler.Handler,
) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/swagger/", docs.ServeSpec).Methods(http.MethodGet)
	router.HandleFunc("/swagger/", docs.ServeUI).Methods(http.MethodGet)
	router.HandleFunc("/swagger", docs.RedirectToUI).Methods(http.MethodGet)

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/validate/json", analyzer.CheckJSON).Methods(http.MethodPost)
	api.HandleFunc("/validate/yaml", analyzer.CheckYAML).Methods(http.MethodPost)
	//TODO добавить Health ручку

	return router
}
