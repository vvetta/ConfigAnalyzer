package main

import (
	transporthttp "ConfigAnalyzer/internal/transport/http"
	docsHandler "ConfigAnalyzer/internal/transport/http/docs"
	analyzerHandler "ConfigAnalyzer/internal/transport/http/handlers"
	"ConfigAnalyzer/internal/usecase"
	"flag"
	"net/http"
	"os"
)

func main() {
	httpMode := flag.Bool("http", false, "Run in HTTP server mode")
	port := flag.String("port", "8080", "Port for HTTP server (used with --http flag)")

	flag.Parse()

	analyzer := usecase.NewConfigAnalyzer()

	if *httpMode {

		analyzerhandler := analyzerHandler.NewConfigAnalyzerHandler(analyzer)
		docshandler := docsHandler.NewDocsHandler()
		router := transporthttp.NewRouter(analyzerhandler, docshandler)

		server := &http.Server{
			Addr:    "localhost:" + *port,
			Handler: router,
		}

		if err := server.ListenAndServe(); err != nil {
			os.Exit(1)
		}

	} else {

	}
}
