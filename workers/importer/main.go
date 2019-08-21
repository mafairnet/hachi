package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var configuration = getProgramConfiguration()
var filename string

func main() {
	fmt.Printf("Hachi isn initializing!\n")

	filename = "pnn_Publico_latest"

	fileUrl := configuration.IftCatalogURL + filename + ".zip?raw=true"

	fmt.Printf("Downloading: %v\n", fileUrl)

	if err := DownloadFile("downloads/"+filename+".zip", fileUrl); err != nil {
		panic(err)
	}

	fmt.Println("Downloaded!")

	fmt.Printf("Unziping file...\n")

	files, err := Unzip("downloads/"+filename+".zip", "uncompressed")
	if err != nil {
		log.Fatal(err)
	}

	filenameTemp := strings.Join(files, "")

	newFilename := strings.Split(filenameTemp, "\\")

	filename = newFilename[1]

	modifyFileToProcess("uncompressed/" + filename)

	cleanDB()

	insertNumberTypesIntoDb()
	numberTypes := getDbNumberTypes()
	fmt.Printf("NumberTypes: %v\n", numberTypes)

	providers := fileImportProviders()
	_ = providers
	sort.Strings(providers)
	fmt.Printf("Providers: %v\n", providers)
	insertProvidersIntoDb(providers)
	providersDb := getDbProviders()
	fmt.Printf("ProvidersDB: %v\n", providersDb)

	states := fileImportStates()
	_ = states
	sort.Strings(states)
	fmt.Printf("States: %v\n", states)
	insertStatesIntoDb(states)
	statesDb := getDbStates()
	fmt.Printf("StatesDB: %v\n", statesDb)

	townships := fileImportTownships()
	_ = townships
	sort.Strings(townships)
	fmt.Printf("TownShips: %v\n", townships)
	insertTownshipsIntoDb(townships, statesDb)
	townshipsDb := getDbTownships()
	fmt.Printf("TownShipsDB: %v\n", townshipsDb)

	towns := fileImportTowns()
	_ = towns
	sort.Strings(towns)
	fmt.Printf("Towns: %v\n", towns)
	insertTownsIntoDb(towns, townshipsDb, statesDb)
	townsDb := getDbTowns()
	fmt.Printf("TownsDB: %v\n", townsDb)

	numbers := fileImportNumbers()
	_ = numbers
	sort.Strings(numbers)
	fmt.Printf("Numbers: %v\n", numbers)
	insertNumbersIntoDb(numbers, townsDb, townshipsDb, statesDb, providersDb)

}
