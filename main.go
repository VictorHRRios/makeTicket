package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

func configPath() string {
	if dir, err := os.UserConfigDir(); err == nil {
		return filepath.Join(dir, "makeTicket", "config.json")
	}
	return "config.json"
}

func getJSONConfig() (ConfigFile, error) {
	config := ConfigFile{}
	cfgPath := configPath()
	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return ConfigFile{}, err
	}
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("Error al parsear config.json: %s", err)
	}
	return config, nil
}

func UpdateJSONConfig(ticketNumber string) error {
	config, err := getJSONConfig()
	if err != nil {
		return err
	}
	config.Current = ticketNumber
	jsonData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath(), jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func getConfig() (Config, error) {
	today := time.Now()
	year, week := today.ISOWeek()
	dateStr := fmt.Sprintf("%d_%d", year, week)
	log.Printf("ticketMaker %s", dateStr)

	config, err := getJSONConfig()
	if err != nil {
		return Config{}, err
	}

	ticketConfig := Config{
		AllProjects:       config.Projects,
		AllVersionFolders: config.Versions,
		Status:            GeneralStatus,
		DateStr:           dateStr,
		CloudURL:          config.CloudURL,
	}
	var configFilePath string
	configFilePath = ".config"

	configPath := path.Join(".", dateStr, config.Current)

	_, err = os.Stat(configFilePath)
	if errors.Is(err, os.ErrNotExist) {
		log.Println("Reading local config from config path")
		configFilePath = path.Join(configPath, ".config")
	}

	metadata, err := os.Open(configFilePath)
	if err != nil {
		log.Printf("LOG: not a project: %s", err)
		return ticketConfig, nil
	}

	defer metadata.Close()

	reader := csv.NewReader(metadata)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("LOG: not a project: %s", err)
		return ticketConfig, nil
	}

	if len(records) < 2 || len(records[1]) < 3 {
		log.Printf("LOG: not a project: invalid metadata format")
		return ticketConfig, nil
	}

	ticketConfig.TicketNumber = records[1][0]
	ticketConfig.ProjectID = records[1][1]
	ticketConfig.VersionID = records[1][2]
	ticketConfig.Path = configPath

	return ticketConfig, err
}

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}
	log.Printf("dateStr: %s", cfg.DateStr)
	var params string
	if len(os.Args) < 2 {
		log.Printf("Please use a command")
		message, _ := cfg.getCommands()["help"].callback("")
		log.Print(message)
		return
	}
	command := os.Args[1]
	if len(command) == 0 {
		log.Printf("Please use a command")
		message, _ := cfg.getCommands()["help"].callback("")
		log.Print(message)
		return
	}
	if len(os.Args) > 2 {
		params = os.Args[2]
	}
	if commandFunc, ok := cfg.getCommands()[command]; ok {
		message, err := commandFunc.callback(params)
		if err != nil {
			log.Printf("Error: %s", err)
			return
		}
		log.Printf("Message: %s", message)
	} else {
		log.Println("Please use a command")
		message, _ := cfg.getCommands()["help"].callback("")
		log.Print(message)
	}
}
