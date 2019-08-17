package main

//Number The object that handles the number data
type Number struct {
	IDNumber          int        `json:"id_number"`
	Prefix            int        `json:"prefix"`
	Series            int        `json:"series"`
	InitialNumeration int        `json:"initial_numeration"`
	FinalNumeration   int        `json:"final_numeration"`
	NumberProvider    Provider   `json:"provider"`
	NumberNumberType  NumberType `json:"number_type"`
	NumberTown        Town       `json:"town"`
}
