package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func (c Config) createFiles(project, version string) error {
	fmt.Printf("Creating ticket %s schema in directory: %s\n", c.TicketNumber, c.DateStr)

	ticketDir := "./" + c.DateStr + "/" + c.TicketNumber
	if err := os.MkdirAll(ticketDir, os.ModePerm); err != nil {
		return err
	}
	ticketPMEmpty, err := createFileInPath(ticketDir, fmt.Sprintf("%s.txt", c.TicketNumber))
	if err != nil {
		return err
	}
	defer ticketPMEmpty.file.Close()
	log.Printf("Created file %s\n", ticketPMEmpty.path)

	if err := c.fillTemplatePM(ticketPMEmpty.file); err != nil {
		return err
	}

	metadata, err := createFileInPath(ticketDir, ".config")
	if err != nil {
		return err
	}
	defer metadata.file.Close()
	_, err = fmt.Fprintf(metadata.file, "ticket,project,version\n%s,%s,%s", c.TicketNumber, project, version)
	if err != nil {
		return err
	}
	return nil
}

func openEditor(filename string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nvim"
	}
	cmd := exec.Command(editor, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
