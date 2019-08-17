package main

//Township The object that handles the township data
type Township struct {
	IDTownship    int    `json:"id_township"`
	Description   string `json:"description"`
	TwonShipState State  `json:"state"`
}
