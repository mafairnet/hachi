package main

//Town The object that handles the town data
type Town struct {
	IDTown       int      `json:"id_town"`
	Description  string   `json:"description"`
	TwonTownship Township `json:"township"`
}

//TownDb The object that handles the town data
type TownDb struct {
	IDTown      int    `json:"id_town"`
	Description string `json:"description"`
	IDTownship  int    `json:"id_township"`
}
