package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func fileImportStates() []string {
	csvFile, err := os.Open("pnn_Publico_15_08_2019.csv")

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	//reader := csv.NewReader(bufio.NewReader(csvFile))
	reader := csv.NewReader(csvFile)

	//var number []Number
	var states []string
	for {
		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		//fmt.Printf("Line: %v", line)
		//Get All States, Remove Duplicates and Store in DB
		if line[3] != " ESTADO" {
			states = append(states, line[3])
		}
	}
	states = removeDuplicatesSliceString(states)
	return states
}

func fileImportTownships() []string {
	var townShips []string

	csvFile, err := os.Open("pnn_Publico_15_08_2019.csv")

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	reader := csv.NewReader(csvFile)

	for {
		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		//fmt.Printf("Line: %v\n", line2)
		//Get All States, Remove Duplicates and Store in DB
		if line[2] != " MUNICIPIO" {
			townShips = append(townShips, line[3]+"-"+line[2])
		}
	}
	townShips = removeDuplicatesSliceString(townShips)
	return townShips
}

func fileImportTowns() []string {
	var towns []string

	csvFile, err := os.Open("pnn_Publico_15_08_2019.csv")

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	reader := csv.NewReader(csvFile)

	for {
		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		//fmt.Printf("Line: %v\n", line2)
		//Get All States, Remove Duplicates and Store in DB
		if line[1] != " POBLACION" {
			towns = append(towns, line[1])
		}
	}
	towns = removeDuplicatesSliceString(towns)
	return towns
}
