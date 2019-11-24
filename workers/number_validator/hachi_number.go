package main

type Number struct {
	IDNumber          int        `json:"id_number"`
	Prefix            int        `json:"prefix"`
	Series            int        `json:"series"`
	InitialNumeration int        `json:"initial_numeration"`
	FinalNumeration   int        `json:"final_numeration"`
	Provider          Provider   `json:"provider"`
	NumberType        NumberType `json:"number_type"`
	Town              Town       `json:"town"`
}
