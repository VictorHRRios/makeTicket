package main

import (
	"errors"
	"log"
	"os"
	"strings"
)

func (c Config) readTemplatePM() (string, error) {
	tokens := map[string]string{
		"ticket":          c.TicketNumber,
		"projectURL":      c.CurrentVersion.SvnURL,
		"dependenciesURL": c.CurrentProject.SVNProjectURL,
		"evidenceURL":     c.CurrentVersion.EvidenceURL,
		"cloudURL":        c.CloudURL,
		"date":            c.DateStr,
	}

	data, err := os.ReadFile(c.templatePath)
	resultingTemplate := []byte{}
	if err != nil {
		return "", err
	}
	var startIdx int
	var endIdx int
	isOpen := []rune{}
	for idx, value := range data {
		switch value {
		case '{':
			isOpen = append(isOpen, '{')
			startIdx = idx
		case '}':
			if len(isOpen) == 0 {
				return "", errors.New("inconsistent values in template")
			}
			if isOpen[len(isOpen)-1] == '{' {
				isOpen = isOpen[:len(isOpen)-1]
				endIdx = idx
				token := strings.TrimSpace(string(data[startIdx+1 : endIdx-1]))
				value, ok := tokens[token]
				if !ok {
					log.Printf("WARNING: { %s } does not exist in the definitions", token)
					continue
				}
				resultingTemplate = append(resultingTemplate, []byte(value)...)
			}
		default:
			if len(isOpen) == 0 {
				resultingTemplate = append(resultingTemplate, value)
			}
		}
	}
	if len(isOpen) != 0 {
		return "", errors.New("inconsistent values in template")
	}
	return string(resultingTemplate), nil
}

func (c Config) fillTemplatePM(file *os.File) error {
	resultingTemplate, err := c.readTemplatePM()
	if err != nil {
		return err
	}
	_, err = file.WriteString(resultingTemplate)
	if err != nil {
		return err
	}
	return nil
}
