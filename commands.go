package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const (
	GeneralStatus = iota
	ProjectStatus
	evidenceStatus
	commitStatus
	FinishedStatus
)

type cliCommand struct {
	name        string
	description string
	callback    func(string) (string, error)
}

func (c *Config) getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"init": {
			name:        "init",
			description: "Initializes a PM",
			callback:    c.init,
		},
		"status": {
			name:        "status",
			description: "Displays the status of the PM",
			callback:    c.status,
		},
		"config": {
			name:        "config",
			description: "print the current configuration",
			callback:    c.config,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    c.help,
		},
		"edit": {
			name:        "edit",
			description: "Edit PM txt",
			callback:    c.edit,
		},
		"xml": {
			name:        "xml",
			description: "Edit xml requests",
			callback:    c.xml,
		},
	}
}

func (c Config) getProject(id int, keys []string) (Project, string, error) {
	if id < 0 || id >= len(keys) {
		return Project{}, "", fmt.Errorf("Error: Not a valid project")
	}
	return c.AllProjects[keys[id]], keys[id], nil
}

func (c Config) getVersion(id int, keys []string) (VersionFolders, string, error) {
	if id < 0 || id >= len(keys) {
		return VersionFolders{}, "", fmt.Errorf("Error: Not a valid version")
	}
	return c.AllVersionFolders[keys[id]], keys[id], nil
}

func getMapKeys[T any](collection map[string]T) []string {
	collectionKeys := make([]string, 0, len(collection))
	for k := range collection {
		collectionKeys = append(collectionKeys, k)
	}
	sort.Strings(collectionKeys)
	return collectionKeys
}

func (c *Config) init(param string) (string, error) {
	var projectName string
	var versionName string

	_, err := strconv.Atoi(param)
	width := 7
	if err != nil {
		return "", err
	}

	if len(param) < width {
		c.TicketNumber = strings.Repeat("0", width-len(param))
	}
	c.TicketNumber += param

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Select the project that the ticket will modify: \n")
	projectKeys := getMapKeys(c.AllProjects)
	for i, key := range projectKeys {
		fmt.Printf("%d) %s\n", i, key)
	}
	projectSelection, err := getIntInput(scanner)
	if err != nil {
		return "", err
	}
	if projectSelection < 0 || projectSelection >= len(c.AllProjects) {
		return "", fmt.Errorf("Error: Not a valid project")
	}
	c.CurrentProject, projectName, err = c.getProject(projectSelection, projectKeys)
	if err != nil {
		return "", err
	}

	fmt.Printf("Select which version you are working on: \n")
	versionKeys := getMapKeys(c.AllVersionFolders)
	for i, key := range versionKeys {
		fmt.Printf("%d) %s\n", i, key)
	}
	versionSelection, err := getIntInput(scanner)
	if err != nil {
		return "", err
	}
	c.CurrentVersion, versionName, err = c.getVersion(versionSelection, versionKeys)
	if err != nil {
		return "", err
	}

	if err = c.createFiles(projectName, versionName); err != nil {
		return "", err
	}

	UpdateJSONConfig(c.TicketNumber)

	return "Ticket directory initialized correctly", nil
}

func (c Config) status(param string) (string, error) {
	return "", nil
}

func (c Config) config(param string) (string, error) {
	var out strings.Builder
	fmt.Fprintf(&out, `
	version: %s
	project: %s
	ticket: %s
	`, c.VersionID, c.ProjectID, c.TicketNumber)
	return out.String(), nil
}

func (c Config) help(s string) (string, error) {
	var out strings.Builder
	out.WriteString("Usage: \n")
	for _, val := range c.getCommands() {
		fmt.Fprintf(&out, "%s: %s\n", val.name, val.description)
	}
	return out.String(), nil
}

func (c Config) edit(s string) (string, error) {
	//if c.Status < 1 {
	//return "", fmt.Errorf("To edit the PM you need to be in a PM directory")
	//}
	ticketDir := filepath.Join(c.Path, c.TicketNumber+".txt")
	return "", openEditor(ticketDir)
}

func (c Config) xml(s string) (string, error) {
	//if c.Status < 1 {
	//return "", fmt.Errorf("To edit the XML you need to be in a PM directory")
	//}
	ticketDir := filepath.Join(c.Path, c.TicketNumber+".txt")

	_, err := os.Stat(ticketDir)
	if errors.Is(err, os.ErrNotExist) {
		_, err = createFileInPath("./", ticketDir)
		if err != nil {
			return "", err
		}
	}
	return "", openEditor(ticketDir)
}
