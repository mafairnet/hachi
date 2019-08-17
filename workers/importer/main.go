package main

import (
	"fmt"
)

var configuration = getProgramConfiguration()

func main() {
	fmt.Printf("Hello!\n")

	states := fileImportStates()
	_ = states
	//fmt.Printf("States: %v\n", states)

	townShips := fileImportTownships()
	_ = townShips
	//fmt.Printf("TownShips: %v\n", townShips)

	towns := fileImportTowns()
	fmt.Printf("Towns: %v\n", towns)
}
