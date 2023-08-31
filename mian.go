package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func searchInDirectory(directory, pattern string) error {
	return filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			return searchInFile(path, pattern)
		})
}
func searchInFile(filepath, pattern string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// increased the buffer size to 1mb
	const maxTokenSize = 1024 * 1024
	buffer := make([]byte, maxTokenSize)
	scanner.Buffer(buffer, maxTokenSize)

	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, pattern) {
			fmt.Printf("%s:%d: %s\n", filepath, lineNumber, line)
		}
		lineNumber++
	}

	return scanner.Err()
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: search-tool <directory> <pattern>")
		os.Exit(1)
	}

	directory := os.Args[1]
	pattern := os.Args[2]

	err := searchInDirectory(directory, pattern)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
