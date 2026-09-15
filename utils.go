package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func createFileInPath(path, name string) (File, error) {
	filePath := filepath.Join(path, name)
	file, err := os.Create(filePath)
	if err != nil {
		return File{}, err
	}
	return File{file: file, path: filePath}, err
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func getIntInput(scannerObj *bufio.Scanner) (int, error) {
	scannerObj.Scan()
	textBytes := scannerObj.Text()
	cleanText := cleanInput(string(textBytes))[0]
	return strconv.Atoi(cleanText)
}
