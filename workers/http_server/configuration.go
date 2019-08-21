package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//Configuration The object that stores the configuration variables
type Configuration struct {
	DbServer      string
	DbUsername    string
	DbPassword    string
	DbPort        string
	DbSchema      string
}

func getCurrentProgramLocation() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	return dir
}

func getProgramConfiguration() Configuration {
	//key := []byte("pr1c3tr4v3l53cr3")

	filename := ""

	programLocation := getCurrentProgramLocation()

	var configuration = Configuration{}

	//Define variables
	if _, err := os.Stat(programLocation + "/config/config.development.json"); err == nil {
		//log.Println("Dev Configuration found, using this values")
		filename = programLocation + "/config/config.development.json"
	} else if _, err := os.Stat(programLocation + "/config/config.production.json"); err == nil {
		//log.Println("Prod Configuration found, using this values")
		filename = programLocation + "/config/config.production.json"
	}

	if filename == "" {
		//log.Fatal("No Configuration File Found, stop running task!")
		os.Exit(3)
	}

	//Open File
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	//Decode File
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	//configuration.GrackleDbPassword = Decrypt(key, configuration.GrackleDbPassword)
	//configuration.PbxDbPassword = Decrypt(key, configuration.PbxDbPassword)
	return configuration
}
