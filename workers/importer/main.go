package main

import (
	"fmt"
	"sort"

	_ "github.com/go-sql-driver/mysql"
)

var configuration = getProgramConfiguration()
var filename string

func main() {
	fmt.Printf("Hachi isn initializing!\n")

	filename = "pnn_Publico_18_08_2019"
	/*
		files, err := Unzip("downloads/"+filename+".zip", "uncompressed")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Unzipped:\n" + strings.Join(files, "\n"))

		//https://sns.ift.org.mx:8081/sns-frontend/planes-numeracion/descarga-publica.xhtml
		fileUrl := "https://golangcode.com/images/avatar.jpg"

		if err := DownloadFile("avatar.jpg", fileUrl); err != nil {
			panic(err)
		}*/

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
