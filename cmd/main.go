package main

import (
	cli "ConfigAnalyzer/internal/transport/cli"
	transporthttp "ConfigAnalyzer/internal/transport/http"
	docsHandler "ConfigAnalyzer/internal/transport/http/docs"
	analyzerHandler "ConfigAnalyzer/internal/transport/http/handlers"
	"ConfigAnalyzer/internal/usecase"
	"errors"
	"flag"
	"net/http"
	"os"
)

func main() {
	var (
		httpMode    bool
		port        string
		silentShort bool
		silentLong  bool
		stdinMode   bool
	)

	flag.BoolVar(&httpMode, "http", false, "Run in HTTP server mode")
	flag.StringVar(&port, "port", "8080", "Port for HTTP server (used with --http flag)")
	flag.BoolVar(&silentShort, "s", false, "Do not exit with error when issues are found")
	flag.BoolVar(&silentLong, "silent", false, "Do not exit with error when issues are found")
	flag.BoolVar(&stdinMode, "stdin", false, "Read config from stdin instead of file")
	flag.Parse()

	analyzer := usecase.NewConfigAnalyzer()

	if httpMode {
		runHTTPServer(analyzer, port)
		return
	}

	if err := cli.RunCLI(analyzer, stdinMode, silentLong || silentShort); err != nil {
		os.Exit(2)
	}
}

func runHTTPServer(analyzer usecase.Analyzer, port string) {
	analyzerhandler := analyzerHandler.NewConfigAnalyzerHandler(analyzer)
	docshandler := docsHandler.NewDocsHandler()
	router := transporthttp.NewRouter(analyzerhandler, docshandler)

	server := &http.Server{
		Addr:    "localhost:" + port,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		//TODO Добавить ошибку!
		os.Exit(1)
	}
}
