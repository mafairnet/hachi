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

	states := fileImportStates()
	//_ = states
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
	//fmt.Printf("Towns: %v\n", towns)
}
