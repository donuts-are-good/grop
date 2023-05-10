package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./grop \"keyword\" /path/to/directory")
		fmt.Println("or")
		fmt.Println("some-program | ./grop \"keyword\"")
		return
	}

	keyword := os.Args[1]

	if len(os.Args) == 2 {
		processInput(keyword, os.Stdin, "")
		return
	}

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

	processInput(keyword, file, path)
}

func processInput(keyword string, input io.Reader, path string) {
	chunkyChunk := 64 * 1024 // chunk size: 64KB
	reader := bufio.NewReader(input)
	buffer := make([]byte, chunkyChunk)
	lineNumber := 1

	for {
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading input:", err)
			return
		}

		if n == 0 {
			break
		}

		lineNumber = searchInChunk(keyword, buffer[:n], lineNumber, path)
	}

}

func searchInChunk(keyword string, chunk []byte, startLine int, path string) int {
	lineNumber := startLine
	scanner := bufio.NewScanner(bytes.NewReader(chunk))

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			start := strings.Index(line, keyword) - 10
			if start < 0 {
				start = 0
			}
			end := start + len(keyword) + 20
			if end > len(line) {
				end = len(line)
			}
			excerpt := strings.TrimSpace(line[start:end])
			if path != "" {
				fmt.Printf("%s: ", path)
			}
			fmt.Printf("line %d: \"%s\"\n", lineNumber, excerpt)
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading chunk:", err)
	}

	return lineNumber
}
