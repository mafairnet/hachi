package main

type Town struct {
	IDTown      int   `json:"id_town"`
	Description string   `json:"description"`
	Township    Township `json:"township"`
}
