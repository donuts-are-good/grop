package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./grop \"keyword\" /path/to/directory")
		return
	}

	keyword := os.Args[1]
	path := os.Args[2]

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			searchInFile(keyword, path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking the path:", err)
	}
}

func searchInFile(keyword, path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			fmt.Printf("%s -> line %d: \"%s\"\n", path, lineNumber, line)
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
