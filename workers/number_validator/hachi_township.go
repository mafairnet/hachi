package main

type Township struct {
	IDTownship  int `json:"id_township"`
	Description string `json:"description"`
	State       State  `json:"state"`
}
