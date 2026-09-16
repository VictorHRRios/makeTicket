package main

import "os"

type File struct {
	file *os.File
	path string
}

type Project struct {
	Name          string    `json:"name"`
	SVNProjectURL string    `json:"svnProjectURL"`
	Programs      []Program `json:"programs,omitempty"`
	ProjectID     int
}

type Program struct {
	Name       string  `json:"name"`
	TypeOfFile string  `json:"typeOfFile"`
	Version    float64 `json:"version"`
	ProgramID  int
}

type VersionFolders struct {
	EvidenceURL string `json:"evidenceURL"`
	SvnURL      string `json:"svnURL"`
}

type ConfigFile struct {
	Projects     map[string]Project        `json:"projects"`
	Versions     map[string]VersionFolders `json:"versions"`
	Current      string                    `json:"current"`
	CloudURL     string                    `json:"cloudURL"`
	TemplatePath string                    `json:"templatePath"`
}

type Config struct {
	AllProjects       map[string]Project
	AllVersionFolders map[string]VersionFolders
	TicketNumber      string
	CurrentProject    Project
	CurrentVersion    VersionFolders
	Path              string
	ProjectID         string
	VersionID         string
	Status            int
	DateStr           string
	CloudURL          string
	templatePath      string
}
