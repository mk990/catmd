package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// Function to convert a text file to a Markdown section
func textToMarkdown(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	lang := ""
	ext := filepath.Ext(filename)
	switch ext {
	case ".txt":
		lang = "plane"
	case ".go":
		lang = "go"
	case ".py":
		lang = "python"
	case ".sh":
		lang = "bash"
	case ".php":
		lang = "php"
	case ".rs":
		lang = "rust"
	case ".html", ".htm":
		lang = "html"
	default:
		lang = ""
	}
	base := filepath.Base(filename)
	// title := strings.TrimSuffix(base, filepath.Ext(base))
	if string(content) != "" {
		return fmt.Sprintf("## %s\n\n```%s\n%s\n```\n-----\n\n", base, lang, string(content)), nil
	}
	return "", nil
}

func main() {
	if len(os.Args) > 1 {
		// If command-line arguments are provided, process those
		processFiles(os.Args[1:])
	} else {
		// If no command-line arguments are provided, try to read filenames from stdin
		scanner := bufio.NewScanner(os.Stdin)
		var filenames []string
		for scanner.Scan() {
			filenames = append(filenames, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from stdin:", err)
			return
		}

		if len(filenames) > 0 {
			processFiles(filenames)
		} else {
			fmt.Println("No filenames provided. Usage: go run export_notes.go file1.txt file2.txt ... OR ls -1 file* | go run export_notes.go")
		}
	}
}

func processFiles(files []string) {
	markdownContent := ""
	for _, file := range files {
		markdownSection, err := textToMarkdown(file)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", file, err)
		} else {
			markdownContent += markdownSection
		}
	}

	// Print the combined Markdown content to the console
	fmt.Print(markdownContent)
}
