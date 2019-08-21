package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func fileImportStates() []string {
	csvFile, err := os.Open("uncompressed/" + filename)

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

	csvFile, err := os.Open("uncompressed/" + filename)

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

	csvFile, err := os.Open("uncompressed/" + filename)

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
			towns = append(towns, line[1]+"-"+line[2]+"-"+line[3])
		}
	}
	towns = removeDuplicatesSliceString(towns)
	return towns
}

func fileImportProviders() []string {
	var providers []string

	csvFile, err := os.Open("uncompressed/" + filename)

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
		if line[14] != " RAZON_SOCIAL" {
			providers = append(providers, line[14])
		}
	}
	providers = removeDuplicatesSliceString(providers)
	return providers
}

func fileImportNumbers() []string {
	var numbers []string

	csvFile, err := os.Open("uncompressed/" + filename)

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
		//6-7-8_9
		if line[14] != " RAZON_SOCIAL" {
			numbers = append(numbers, line[7]+"_"+line[8]+"_"+line[9]+"_"+line[10]+"_"+line[12]+"_"+line[1]+"_"+line[2]+"_"+line[3]+"_"+line[14])
		}
	}
	numbers = removeDuplicatesSliceString(numbers)
	return numbers
}
