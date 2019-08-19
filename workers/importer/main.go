package main

import (
	"fmt"
	"sort"

	_ "github.com/go-sql-driver/mysql"
)

var configuration = getProgramConfiguration()

func main() {
	fmt.Printf("Hachi isn initializing!\n")

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
}
