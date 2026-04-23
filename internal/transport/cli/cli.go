package cli

import (
	"ConfigAnalyzer/internal/domain"
	"ConfigAnalyzer/internal/usecase"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func RunCLI(analyzer usecase.Analyzer, stdinMode, silent bool, args []string) error {

	var (
		data         []byte
		configFormat string
		err          error
	)

	if stdinMode {
		data, err := os.ReadFile("/dev/stdin")
		if err != nil {
			return fmt.Errorf("read stdin: %w", err)
		}
		configFormat = detectFormat("", data)
	} else {
		if len(args) == 0 {
			return nil //TODO Ошибка
		}
		path := args[0]
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read config: %w", err)
		}
		configFormat = detectFormat(path, data)
	}

	issues, err := analyzer.Analyze(usecase.AnalyzeRequest{Data: data, Format: configFormat})
	if err != nil {
		return err
	}

	printIssues(issues)

	return nil
}

func printIssues(issues []domain.Issue) {
	if len(issues) == 0 {
		fmt.Println("Проблем не найдено.")
		return
	}

	for _, issue := range issues {
		fmt.Printf("%s: %s\n", issue.Severity, issue.Message)
		if issue.Path != "" {
			fmt.Printf("  path: %s\n", issue.Path)
		}
		fmt.Printf("  recommendation: %s\n", issue.Recommendation)
	}
}

func detectFormat(path string, data []byte) string {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".json":
		return "json"
	case ".yaml":
		return "yaml"
	default:
		//TODO Можно попытаться достать из файла.
		return ""
	}
}
